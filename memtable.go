package ldbclone

import (
	"errors"
)

var ErrKeyNotFound = errors.New("key not found")

type Memtable struct {
	keys  []string
	store map[string][]byte
}

// Return value from map if key is present, else error
func (m *Memtable) Get(key []byte) ([]byte, error) {

	if val, ok := m.store[string(key)]; ok {
		return val, nil
	}
	err := ErrKeyNotFound
	return nil, err
}

// Check if key is present in map
func (m *Memtable) Has(key []byte) (bool, error) {
	if _, ok := m.store[string(key)]; ok {
		return true, nil
	}
	return false, nil
}

// Add key/value pair to map
func (m *Memtable) Put(key, value []byte) error {
	m.store[string(key)] = value
	m.keys = append(m.keys, string(key))

	return nil
}

// Delete key/value pair from map
func (m *Memtable) Delete(key []byte) error {
	keyStr := string(key)
	if _, ok := m.store[keyStr]; ok {
		delete(m.store, keyStr)

		for i, storedKey := range m.keys {
			if storedKey == keyStr {
				m.keys[i] = ""
			}
		}
		return nil
	}
	return ErrKeyNotFound
}

//  Returns an Iterator (see below) for scanning through all key-value pairs in the given
//  range, ordered by key ascending.
func (m *Memtable) RangeScan(start, limit []byte) (Iterator, error) {
	return nil, nil
}

func NewMemtable() *Memtable {
	m := &Memtable{}
	m.store = make(map[string][]byte)

	return m
}

// Used to iterate through values in a specified range in a memtable
type MemtableIterator struct {
	Idx        int64
	RangeStart int64
	RangeENd   int64
	Memtable   *Memtable
	Finished   bool
}

// Moves the iterator to the next key/value pair. Returns false if the iterator is exhausted.
func (i *MemtableIterator) Next() bool {
	if i.Finished {
		return false
	} else if i.Idx == int64(len(i.Memtable.keys)-1) {
		i.Finished = true
		return false
	}

	i.Idx += 1
	return true
}

// Error returns any accumulated error. Exhausting all the key/value pairs is not considered to
// be an error.
func (i *MemtableIterator) Error() error {
	return nil
}

// Key returns the key of the current key/value pair, or nil if done.
func (i *MemtableIterator) Key() []byte {
	if i.Finished {
		return nil
	}
	return []byte(i.Memtable.keys[i.Idx])
}

// Value returns the value of the current key/value pair, or nil if done.
func (i *MemtableIterator) Value() []byte {
	if i.Finished {
		return nil
	}
	key := i.Memtable.keys[i.Idx]
	return i.Memtable.store[key]
}

func NewMemtableIterator(m *Memtable) *MemtableIterator {
	i := &MemtableIterator{
		Idx:      0,
		Memtable: m,
	}

	return i
}
