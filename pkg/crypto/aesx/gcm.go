package aesx

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"golang_template/pkg/crypto/randx"
)

func EncryptGCM(key, nonce, aad, pt []byte) (usedNonce, ct []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, nil, errors.New("aes: invalid key length")
	}
	if nonce == nil {
		nonce, err = randx.Nonce12()
		if err != nil {
			return nil, nil, err
		}
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	ct = gcm.Seal(nil, nonce, pt, aad)
	return nonce, ct, nil
}

func DecryptGCM(key, nonce, aad, ct []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("aes: invalid key length")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm.Open(nil, nonce, ct, aad)
}
