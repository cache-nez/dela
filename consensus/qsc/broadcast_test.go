package qsc

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"go.dedis.ch/dela/consensus/qsc/types"
	"go.dedis.ch/dela/internal/testing/fake"
	"go.dedis.ch/dela/mino"
	"go.dedis.ch/dela/mino/minoch"
	"golang.org/x/xerrors"
)

func TestTLCR_Basic(t *testing.T) {
	n := 5
	k := 5
	bcs := makeTLCR(t, n)

	wg := sync.WaitGroup{}
	wg.Add(n)
	for _, bc := range bcs {
		go func(bc *bTLCR) {
			msg := types.NewMessage(bc.node, types.NewProposal(fake.Message{}))

			for i := 0; i < k; i++ {
				bc.execute(context.Background(), msg)
			}
			wg.Done()
		}(bc)
	}

	wg.Wait()

	require.Equal(t, uint64(k), bcs[0].timeStep)
}

func TestHandlerTLCR_Process(t *testing.T) {
	ch := make(chan types.MessageSet, 1)
	h := hTLCR{
		ch: ch,
		store: &storage{
			previous: types.NewMessageSet(0, 1),
		},
	}

	resp, err := h.Process(mino.Request{Message: types.NewMessageSet(0, 0)})
	require.NoError(t, err)
	require.Nil(t, resp)
	require.NotNil(t, <-ch)

	resp, err = h.Process(mino.Request{Message: types.NewRequestMessageSet(0, nil)})
	require.NoError(t, err)
	require.Nil(t, resp)

	_, err = h.Process(mino.Request{Message: fake.Message{}})
	require.EqualError(t, err, "invalid message type 'fake.Message'")
}

func TestTLCR_Execute(t *testing.T) {
	buffer := new(bytes.Buffer)
	ch := make(chan types.MessageSet, 1)
	rpc := fake.NewRPC()
	bc := &bTLCR{
		logger:  zerolog.New(buffer),
		rpc:     rpc,
		ch:      ch,
		players: fakeSinglePlayer{},
		store:   &storage{},
	}

	ch <- types.NewMessageSet(0, 0, types.NewMessage(0, fake.Message{}))

	view, err := bc.execute(context.Background())
	require.NoError(t, err)
	require.NotNil(t, view)
	require.Equal(t, uint64(1), bc.timeStep)
	require.Equal(t, bc.store.previous.GetMessages(), view.GetReceived())
	require.Len(t, view.broadcasted, 0)

	rpc.SendResponseWithError(nil, xerrors.New("oops"))
	rpc.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	_, err = bc.execute(ctx)
	require.EqualError(t, err, "context deadline exceeded")
	require.Contains(t, buffer.String(), "oops")

	bc.rpc = fake.NewBadRPC()
	_, err = bc.execute(ctx)
	require.EqualError(t, err, "call aborted: fake error")
}

func TestTLCR_Merge(t *testing.T) {
	m1 := types.NewMessageSet(0, 0, types.NewMessage(1, nil), types.NewMessage(2, nil))
	m2 := types.NewMessageSet(0, 0, types.NewMessage(2, nil), types.NewMessage(3, nil))

	m3 := m1.Merge(m2)
	require.Len(t, m3.GetMessages(), 3)
}

func TestTLCR_CatchUp(t *testing.T) {
	ch := make(chan types.MessageSet, 1)
	rpc := fake.NewRPC()
	bc := &bTLCR{
		ch:       ch,
		timeStep: 0,
		rpc:      rpc,
		players:  fakeSinglePlayer{},
	}

	ctx := context.Background()

	m1 := types.NewMessageSet(0, 0, types.NewMessage(1, nil))
	m2 := types.NewMessageSet(2, 0)
	err := bc.catchUp(ctx, m1, m2)
	require.NoError(t, err)
	require.Equal(t, m2, <-ch)

	m2 = types.NewMessageSet(2, 1)
	rpc.SendResponse(nil, types.NewMessageSet(0, 0, types.NewMessage(2, fake.Message{})))
	err = bc.catchUp(ctx, m1, m2)
	require.NoError(t, err)
	require.Equal(t, m2, <-ch)

	rpc.SendResponse(nil, fake.Message{})
	err = bc.catchUp(ctx, m1, m2)
	require.EqualError(t, xerrors.Unwrap(err),
		"got message type 'fake.Message' but expected 'types.MessageSet'")
	require.Equal(t, m2, <-ch)

	rpc.SendResponseWithError(nil, xerrors.New("oops"))
	err = bc.catchUp(ctx, m1, m2)
	require.EqualError(t, err,
		"couldn't fetch previous message set: couldn't reach the node: oops")
	require.Equal(t, m2, <-ch)

	rpc.Done()
	err = bc.catchUp(ctx, m1, m2)
	require.EqualError(t, err,
		"couldn't fetch previous message set: couldn't get a reply")
	require.Equal(t, m2, <-ch)

	bc.rpc = fake.NewBadRPC()
	err = bc.catchUp(ctx, m1, m2)
	require.EqualError(t, err,
		"couldn't fetch previous message set: call aborted: fake error")
}

func TestTLCB_Basic(t *testing.T) {
	n := 3
	k := 5

	bcs := makeTLCB(t, n)
	wg := sync.WaitGroup{}
	wg.Add(n)
	for _, bc := range bcs {
		go func(bc *bTLCB) {
			defer wg.Done()
			var view View
			var err error
			for i := 0; i < k; i++ {
				view, err = bc.execute(context.Background(), fake.Message{})
				require.NoError(t, err)
				require.Len(t, view.broadcasted, n)
				require.Len(t, view.received, n)
			}
		}(bc)
	}

	wg.Wait()
}

func TestTLCB_Execute(t *testing.T) {
	bc := &bTLCB{
		b1: fakeTLCR{},
		b2: fakeTLCR{},
	}

	view, err := bc.execute(context.Background(), fake.Message{})
	require.NoError(t, err)
	require.NotNil(t, view)

	bc.b1 = fakeTLCR{err: xerrors.New("oops")}
	_, err = bc.execute(context.Background(), fake.Message{})
	require.EqualError(t, err, "couldn't broadcast: oops")

	bc.b1 = fakeTLCR{}
	bc.b2 = fakeTLCR{err: xerrors.New("oops")}
	_, err = bc.execute(context.Background(), fake.Message{})
	require.EqualError(t, err, "couldn't broadcast: oops")
}

func makeTLCR(t *testing.T, n int) []*bTLCR {
	manager := minoch.NewManager()
	players := &fakePlayers{}
	bcs := make([]*bTLCR, n)
	for i := range bcs {
		m, err := minoch.NewMinoch(manager, fmt.Sprintf("node%d", i))
		require.NoError(t, err)

		players.addrs = append(players.addrs, m.GetAddress())

		bc, err := newTLCR("tlcr", int64(i), m, players, fake.MessageFactory{})
		require.NoError(t, err)

		bcs[i] = bc
	}

	return bcs
}

func makeTLCB(t *testing.T, n int) []*bTLCB {
	manager := minoch.NewManager()
	players := &fakePlayers{}
	bcs := make([]*bTLCB, n)
	for i := range bcs {
		m, err := minoch.NewMinoch(manager, fmt.Sprintf("node%d", i))
		require.NoError(t, err)

		players.addrs = append(players.addrs, m.GetAddress())

		bc, err := newTLCB(int64(i), m, players, fake.MessageFactory{})
		require.NoError(t, err)

		bcs[i] = bc
	}

	return bcs
}

// -----------------
// Utility functions

type fakeIterator struct {
	mino.AddressIterator
	index int
	addrs []mino.Address
}

func (i *fakeIterator) HasNext() bool {
	return i.index+1 < len(i.addrs)
}

func (i *fakeIterator) GetNext() mino.Address {
	if i.HasNext() {
		i.index++
		return i.addrs[i.index]
	}
	return nil
}

type fakePlayers struct {
	addrs []mino.Address
}

func (p *fakePlayers) Take(filters ...mino.FilterUpdater) mino.Players {
	ff := mino.ApplyFilters(filters)
	addrs := make([]mino.Address, len(ff.Indices))
	for i, k := range ff.Indices {
		addrs[i] = p.addrs[k]
	}
	return &fakePlayers{addrs: addrs}
}

func (p *fakePlayers) AddressIterator() mino.AddressIterator {
	return &fakeIterator{addrs: p.addrs, index: -1}
}

func (p *fakePlayers) Len() int {
	return len(p.addrs)
}

type fakeSinglePlayer struct {
	mino.Players
}

func (p fakeSinglePlayer) Len() int {
	return 1
}

func (p fakeSinglePlayer) Take(...mino.FilterUpdater) mino.Players {
	return fakeSinglePlayer{}
}

type fakeTLCR struct {
	err error
}

func (b fakeTLCR) execute(context.Context, ...types.Message) (View, error) {
	return View{}, b.err
}
