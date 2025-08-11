package aesx

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestAES_GCM(t *testing.T) {
	key, _ := hex.DecodeString("00112233445566778899aabbccddeeff00112233445566778899aabbccddeeff")
	nonce := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	pt := []byte("hello gcm")
	used, ct, err := EncryptGCM(key, nonce, nil, pt)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(used, nonce) {
		t.Fatal("nonce mismatch")
	}
	got, err := DecryptGCM(key, nonce, nil, ct)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, pt) {
		t.Fatalf("roundtrip mismatch: %q", got)
	}
}

func TestAES_CBC(t *testing.T) {
	key, _ := hex.DecodeString("00112233445566778899aabbccddeeff00112233445566778899aabbccddeeff")
	iv := []byte("1234567890abcdef")
	pt := []byte("hello cbc") // will be padded
	usedIV, ct, err := EncryptCBC(key[:16], iv, pt)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(usedIV, iv) {
		t.Fatal("iv mismatch")
	}
	got, err := DecryptCBC(key[:16], iv, ct)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, pt) {
		t.Fatalf("roundtrip mismatch: %q", got)
	}
}
