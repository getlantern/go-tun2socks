package core

import (
	pool "github.com/libp2p/go-buffer-pool"
)

const BufSize = 2 * 1024

func NewBytes(size int) []byte {
	return pool.Get(size)
}

func FreeBytes(b []byte) {
	pool.Put(b)
}
