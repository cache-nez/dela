package mem

import (
	"io"

	"go.dedis.ch/dela/core/store"
	"go.dedis.ch/dela/crypto"
	"golang.org/x/xerrors"
)

type item struct {
	value   []byte
	deleted bool
}

// Trie is an in-memory implementation of a store. It saves the updates in an
// internal store and only keep the updates of the current trie. When reading,
// it'll look up by following the parent trie if the key is not found.
//
// TODO: merkle trie + share
type Trie struct {
	parent      *Trie
	root        []byte
	store       map[string]item
	hashFactory crypto.HashFactory
}

// NewTrie creates a new in-memory trie.
func NewTrie() *Trie {
	return &Trie{
		parent:      nil,
		store:       make(map[string]item),
		hashFactory: crypto.NewSha256Factory(),
	}
}

// Get implements store.Reader. It returns the value associated with the key if
// it exists by first checking the current store then recursively checking the
// parent up to the root.
func (t *Trie) Get(key []byte) ([]byte, error) {
	item, found := t.store[string(key)]
	if found {
		if item.deleted {
			// Item is explicitly deleted so we skip recursive look up.
			return nil, nil
		}

		return item.value, nil
	}

	if t.parent == nil {
		return nil, nil
	}

	val, err := t.parent.Get(key)
	if err != nil {
		// The error is wrapped to avoid big recursive errors.
		return nil, err
	}

	return val, nil
}

// Set implements store.Writer. It writes the value to the store.
func (t *Trie) Set(key, value []byte) error {
	t.store[string(key)] = item{
		value: value,
	}

	return nil
}

// Delete implements store.Writer. It removes a key from the store and flag it
// as deleted.
func (t *Trie) Delete(key []byte) error {
	// Even if the item does not exist, it must be flagged so that a parent
	// won't return the value.
	t.store[string(key)] = item{
		deleted: true,
	}

	return nil
}

// GetRoot implements store.Trie.
func (t *Trie) GetRoot() []byte {
	return t.root
}

// GetShare implements store.Store.
func (t *Trie) GetShare(key []byte) (store.Share, error) {
	value, err := t.Get(key)
	if err != nil {
		return nil, err
	}

	return newShare(key, value, t.root), nil
}

// Fingerprint implements serde.Fingerprinter. It deterministically writes a
// binary representation of the trie.
func (t *Trie) Fingerprint(w io.Writer) error {
	if t.parent != nil {
		_, err := w.Write(t.parent.root)
		if err != nil {
			return xerrors.Errorf("couldn't write parent root: %v", err)
		}
	}

	for key, value := range t.store {
		if value.deleted {
			continue
		}

		_, err := w.Write([]byte(key))
		if err != nil {
			return xerrors.Errorf("couldn't write key: %v", err)
		}

		_, err = w.Write(value.value)
		if err != nil {
			return xerrors.Errorf("couldn't write value: %v", err)
		}
	}

	return nil
}

// Stage implements store.Trie. It executes the callback over a child of the
// current trie and return the trie with the root calculated.
func (t *Trie) Stage(fn func(store.ReadWriteTrie) error) (store.Store, error) {
	trie := t.makeChild()

	err := fn(trie)
	if err != nil {
		return nil, xerrors.Errorf("callback failed: %v", err)
	}

	h := t.hashFactory.New()
	err = trie.Fingerprint(h)
	if err != nil {
		return nil, xerrors.Errorf("couldn't compute root: %v", err)
	}

	trie.root = h.Sum(nil)

	return trie, nil
}

func (t *Trie) makeChild() *Trie {
	clone := NewTrie()
	clone.parent = t
	clone.hashFactory = t.hashFactory

	return clone
}
