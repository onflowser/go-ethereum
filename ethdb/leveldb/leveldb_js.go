//go:build js
// +build js

package leveldb

import (
	"github.com/onflow/go-ethereum/ethdb"
)

type Database struct{}

func New(file string, cache int, handles int, namespace string, readonly bool) (*Database, error) {
	panic("Unimplemented")
}

func (db *Database) Compact(start []byte, limit []byte) error {
	panic("Unimplemented")
}

func (db *Database) NewSnapshot() (ethdb.Snapshot, error) {
	panic("Unimplemented")
}

func (b *Database) Delete(key []byte) error {
	panic("Unimplemented")
}

func (d *Database) Get(key []byte) ([]byte, error) {
	panic("Unimplemented")
}

func (d *Database) Has(key []byte) (bool, error) {
	panic("Unimplemented")
}

func (d *Database) Release() {
	panic("Unimplemented")
}

// database until a final write is called.
func (d *Database) NewBatch() ethdb.Batch {
	panic("Unimplemented")
}

func (d *Database) NewBatchWithSize(size int) ethdb.Batch {
	panic("Unimplemented")
}

func (d *Database) NewIterator(prefix []byte, start []byte) ethdb.Iterator {
	panic("Unimplemented")
}

func (d *Database) Put(key, value []byte) error {
	panic("Unimplemented")
}

func (d *Database) ValueSize() int {
	panic("Unimplemented")
}

func (d *Database) Write() error {
	panic("Unimplemented")
}

func (d *Database) Reset() {
	panic("Unimplemented")
}

func (d *Database) Replay(w ethdb.KeyValueWriter) error {
	panic("Unimplemented")
}

func (d *Database) Stat(property string) (string, error) {
	panic("Unimplemented")
}

func (db *Database) Close() error {
	panic("Unimplemented")
}
