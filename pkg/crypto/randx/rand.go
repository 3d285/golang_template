package randx

import (
	"crypto/rand"
	"io"
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, b)
	return b, err
}

func Nonce12() ([]byte, error) { return Bytes(12) }
