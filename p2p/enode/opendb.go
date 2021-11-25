// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

//go:build !js
// +build !js

package enode

import (
	"bytes"
	"encoding/binary"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

// newPersistentNodeDB creates/opens a leveldb backed persistent node database,
// also flushing its contents in case of a version mismatch.
func newPersistentDB(path string) (*DB, error) {
	opts := &opt.Options{OpenFilesCacheCapacity: 5}
	db, err := leveldb.OpenFile(path, opts)
	if _, iscorrupted := err.(*errors.ErrCorrupted); iscorrupted {
		db, err = leveldb.RecoverFile(path, nil)
	}
	if err != nil {
		return nil, err
	}
	// The nodes contained in the cache correspond to a certain protocol version.
	// Flush all nodes if the version doesn't match.
	currentVer := make([]byte, binary.MaxVarintLen64)
	currentVer = currentVer[:binary.PutVarint(currentVer, int64(dbVersion))]

	blob, err := db.Get([]byte(dbVersionKey), nil)
	switch err {
	case leveldb.ErrNotFound:
		// Version not found (i.e. empty cache), insert it
		if err := db.Put([]byte(dbVersionKey), currentVer, nil); err != nil {
			db.Close()
			return nil, err
		}

	case nil:
		// Version present, flush if different
		if !bytes.Equal(blob, currentVer) {
			db.Close()
			if err = os.RemoveAll(path); err != nil {
				return nil, err
			}
			return newPersistentDB(path)
		}
	}
	return &DB{lvl: db, quit: make(chan struct{})}, nil
}
