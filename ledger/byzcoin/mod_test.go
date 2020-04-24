package byzcoin

import (
	"context"
	fmt "fmt"
	"testing"
	"time"

	proto "github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"go.dedis.ch/fabric/blockchain"
	"go.dedis.ch/fabric/crypto"
	"go.dedis.ch/fabric/crypto/bls"
	internal "go.dedis.ch/fabric/internal/testing"
	"go.dedis.ch/fabric/internal/testing/fake"
	"go.dedis.ch/fabric/ledger"
	"go.dedis.ch/fabric/ledger/arc/darc"
	"go.dedis.ch/fabric/ledger/arc/darc/contract"
	"go.dedis.ch/fabric/ledger/consumer"
	"go.dedis.ch/fabric/ledger/consumer/smartcontract"
	"go.dedis.ch/fabric/mino"
	"go.dedis.ch/fabric/mino/minoch"
	"golang.org/x/xerrors"
)

func TestMessages(t *testing.T) {
	messages := []proto.Message{
		&BlockPayload{},
		&Roster{},
		&GenesisPayload{},
	}

	for _, m := range messages {
		internal.CoverProtoMessage(t, m)
	}
}

// This test checks the basic behaviour of a Byzcoin ledger. The module should
// do the following steps without errors:
// 1. Run n nodes and start to listen for requests
// 2. Setup the ledger on the leader (as we use a leader-based view change)
// 3. Send transactions and accept them.
func TestLedger_Basic(t *testing.T) {
	ledgers, actors, ca := makeLedger(t, 20)
	defer func() {
		for _, actor := range actors {
			require.NoError(t, actor.Close())
		}
	}()

	require.NoError(t, actors[0].Setup(ca))

	for _, actor := range actors {
		err := <-actor.HasStarted()
		require.Nil(t, err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	txs := ledgers[2].Watch(ctx)

	txFactory := smartcontract.NewTransactionFactory(bls.NewSigner())

	// Try to create a DARC.
	tx, err := txFactory.New(contract.NewGenesisAction())
	require.NoError(t, err)

	err = actors[1].AddTransaction(tx)
	require.NoError(t, err)

	select {
	case res := <-txs:
		require.NotNil(t, res)
		require.Equal(t, tx.GetID(), res.GetTransactionID())
	case <-time.After(1 * time.Second):
		t.Fatal("timeout 1")
	}

	instance, err := ledgers[2].GetInstance(tx.GetID())
	require.NoError(t, err)
	require.Equal(t, tx.GetID(), instance.GetKey())
	require.Equal(t, tx.GetID(), instance.GetArcID())
	require.IsType(t, (*darc.AccessControlProto)(nil), instance.GetValue())

	// Then update it.
	tx, err = txFactory.New(contract.NewUpdateAction(tx.GetID()))
	require.NoError(t, err)

	err = actors[0].AddTransaction(tx)
	require.NoError(t, err)

	select {
	case res := <-txs:
		require.NotNil(t, res)
		require.Equal(t, tx.GetID(), res.GetTransactionID())
	case <-time.After(1 * time.Second):
		t.Fatal("timeout 2")
	}
}

func TestLedger_GetInstance(t *testing.T) {
	ledger := &Ledger{
		bc: fakeBlockchain{},
		proc: &txProcessor{
			inventory: fakeInventory{
				page: &fakePage{},
			},
		},
		consumer: fakeConsumer{},
	}

	instance, err := ledger.GetInstance([]byte{0xab})
	require.NoError(t, err)
	require.NotNil(t, instance)

	ledger.bc = fakeBlockchain{err: xerrors.New("oops")}
	_, err = ledger.GetInstance(nil)
	require.EqualError(t, err, "couldn't read latest block: oops")

	ledger.bc = fakeBlockchain{}
	ledger.proc.inventory = fakeInventory{err: xerrors.New("oops")}
	_, err = ledger.GetInstance(nil)
	require.EqualError(t, err, "couldn't read the page: oops")

	ledger.proc.inventory = fakeInventory{page: &fakePage{err: xerrors.New("oops")}}
	_, err = ledger.GetInstance(nil)
	require.EqualError(t, err, "couldn't read the instance: oops")

	ledger.proc.inventory = fakeInventory{page: &fakePage{}}
	ledger.consumer = fakeConsumer{errFactory: xerrors.New("oops")}
	_, err = ledger.GetInstance(nil)
	require.EqualError(t, err, "couldn't decode instance: oops")
}

// -----------------------------------------------------------------------------
// Utility functions

func makeLedger(t *testing.T, n int) ([]ledger.Ledger, []ledger.Actor, crypto.CollectiveAuthority) {
	manager := minoch.NewManager()

	minos := make([]mino.Mino, n)
	for i := 0; i < n; i++ {
		m, err := minoch.NewMinoch(manager, fmt.Sprintf("node%d", i))
		require.NoError(t, err)

		minos[i] = m
	}

	ca := fake.NewAuthorityFromMino(bls.NewSigner, minos...)
	ledgers := make([]ledger.Ledger, n)
	actors := make([]ledger.Actor, n)
	for i, m := range minos {
		ledger := NewLedger(m, ca.GetSigner(i), makeConsumer())
		ledgers[i] = ledger

		actor, err := ledger.Listen()
		require.NoError(t, err)

		actors[i] = actor
	}

	return ledgers, actors, ca
}

func makeConsumer() consumer.Consumer {
	c := smartcontract.NewConsumer()
	contract.RegisterContract(c)

	return c
}

type fakeBlock struct {
	blockchain.Block
}

func (b fakeBlock) GetIndex() uint64 {
	return 0
}

type fakeBlockchain struct {
	blockchain.Blockchain
	err error
}

func (bc fakeBlockchain) GetBlock() (blockchain.Block, error) {
	return fakeBlock{}, bc.err
}
