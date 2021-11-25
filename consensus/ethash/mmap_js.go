// Copyright 2017 The go-ethereum Authors
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

// Package ethash implements the ethash proof-of-work consensus engine.

//go:build js
// +build js

package ethash

import (
	"errors"
	"os"
	"reflect"
	"sync"
	"unsafe"
)

type MMap []byte

// cache wraps an ethash cache with some metadata to allow easier concurrent use.
type cache struct {
	epoch uint64    // Epoch for which this cache is relevant
	dump  *os.File  // File descriptor of the memory mapped cache
	mmap  MMap      // Memory map itself to unmap before releasing
	cache []uint32  // The actual cache data content (may be memory mapped)
	once  sync.Once // Ensures the cache is generated only once
}

// dataset wraps an ethash dataset with some metadata to allow easier concurrent use.
type dataset struct {
	epoch   uint64    // Epoch for which this cache is relevant
	dump    *os.File  // File descriptor of the memory mapped cache
	mmap    MMap      // Memory map itself to unmap before releasing
	dataset []uint32  // The actual cache data content
	once    sync.Once // Ensures the cache is generated only once
	done    uint32    // Atomic flag to determine generation status
}

var errNotSupported = errors.New("mmap not supported in js")

// memoryMap tries to memory map a file of uint32s for read only access.
func memoryMap(path string, lock bool) (*os.File, MMap, []uint32, error) {
	return nil, nil, nil, errNotSupported
}

// memoryMapFile tries to memory map an already opened file descriptor.
func memoryMapFile(file *os.File, write bool) (MMap, []uint32, error) {
	return nil, nil, errNotSupported
}

// memoryMapAndGenerate tries to memory map a temporary file of uint32s for write
// access, fill it with the data from a generator and then move it into the final
// path requested.
func memoryMapAndGenerate(path string, size uint64, lock bool, generator func(buffer []uint32)) (*os.File, MMap, []uint32, error) {
	return nil, nil, nil, errNotSupported
}

func (m *MMap) header() *reflect.SliceHeader {
	return (*reflect.SliceHeader)(unsafe.Pointer(m))
}

func (m *MMap) addrLen() (uintptr, uintptr) {
	header := m.header()
	return header.Data, uintptr(header.Len)
}

// Lock keeps the mapped region in physical memory, ensuring that it will not be
// swapped out.
func (m MMap) Lock() error {
	return errNotSupported
}

// Unlock reverses the effect of Lock, allowing the mapped region to potentially
// be swapped out.
// If m is already unlocked, aan error will result.
func (m MMap) Unlock() error {
	return errNotSupported
}

// Flush synchronizes the mapping's contents to the file's contents on disk.
func (m MMap) Flush() error {
	return errNotSupported
}

// Unmap deletes the memory mapped region, flushes any remaining changes, and sets
// m to nil.
// Trying to read or write any remaining references to m after Unmap is called will
// result in undefined behavior.
// Unmap should only be called on the slice value that was originally returned from
// a call to Map. Calling Unmap on a derived slice may cause errors.
func (m *MMap) Unmap() error {
	return errNotSupported
}
