package json

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.dedis.ch/dela/crypto/bls"
	"go.dedis.ch/dela/internal/testing/fake"
	"go.dedis.ch/dela/serde"
	"go.dedis.ch/kyber/v3"
	"golang.org/x/xerrors"
)

func TestPubkeyFormat_Encode(t *testing.T) {
	signer := bls.NewSigner()
	format := pubkeyFormat{}
	ctx := fake.NewContextWithFormat(serde.FormatJSON)

	data, err := format.Encode(ctx, signer.GetPublicKey())
	require.NoError(t, err)
	require.Contains(t, string(data), fmt.Sprintf(`{"Name":"%s","Data":`, bls.Algorithm))

	_, err = format.Encode(ctx, fake.Message{})
	require.EqualError(t, err, "unsupported message of type 'fake.Message'")

	_, err = format.Encode(ctx, bls.NewPublicKeyFromPoint(badPoint{}))
	require.EqualError(t, err, "couldn't marshal point: oops")

	_, err = format.Encode(fake.NewBadContext(), signer.GetPublicKey())
	require.EqualError(t, err, "couldn't marshal: fake error")
}

func TestPubkeyFormat_Decode(t *testing.T) {
	signer := bls.NewSigner()
	format := pubkeyFormat{}
	ctx := fake.NewContextWithFormat(serde.FormatJSON)

	data, err := signer.GetPublicKey().Serialize(ctx)
	require.NoError(t, err)

	pubkey, err := format.Decode(ctx, data)
	require.NoError(t, err)
	require.True(t, signer.GetPublicKey().Equal(pubkey.(bls.PublicKey)))

	_, err = format.Decode(ctx, []byte(`{"Data":[]}`))
	require.EqualError(t, err,
		"couldn't unmarshal point: bn256.G2: not enough data")

	_, err = format.Decode(fake.NewBadContext(), []byte(`{}`))
	require.EqualError(t, err, "couldn't deserialize data: fake error")
}

func TestSigFormat_Encode(t *testing.T) {
	sig := bls.NewSignature([]byte("deadbeef"))
	format := sigFormat{}
	ctx := fake.NewContextWithFormat(serde.FormatJSON)

	data, err := format.Encode(ctx, sig)
	require.NoError(t, err)
	require.Contains(t, string(data), fmt.Sprintf(`{"Name":"%s","Data":`, bls.Algorithm))

	_, err = format.Encode(ctx, fake.Message{})
	require.EqualError(t, err, "unsupported message of type 'fake.Message'")

	_, err = format.Encode(fake.NewBadContext(), sig)
	require.EqualError(t, err, "couldn't marshal: fake error")
}

func TestSigFormat_Decode(t *testing.T) {
	format := sigFormat{}
	ctx := serde.NewContext(fake.ContextEngine{})

	sig, err := format.Decode(ctx, []byte(`{"Data":"QQ=="}`))
	require.NoError(t, err)
	require.Equal(t, bls.NewSignature([]byte("A")), sig)

	_, err = format.Decode(fake.NewBadContext(), []byte(`{"Data":"QQ=="}`))
	require.EqualError(t, err, "couldn't deserialize data: fake error")
}

// -----------------------------------------------------------------------------
// Utility functions

type badPoint struct {
	kyber.Point
}

func (p badPoint) MarshalBinary() ([]byte, error) {
	return nil, xerrors.New("oops")
}