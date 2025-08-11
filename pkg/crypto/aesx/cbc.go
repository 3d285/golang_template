package aesx

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"golang_template/pkg/crypto/randx"
)

func EncryptCBC(key, iv, pt []byte) (usedIV, ct []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, nil, errors.New("aes: invalid key length")
	}
	if iv == nil {
		iv, err = randx.Bytes(aes.BlockSize)
		if err != nil {
			return nil, nil, err
		}
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	padded := pkcs7Pad(pt, aes.BlockSize)
	ct = make([]byte, len(padded))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ct, padded)
	return iv, ct, nil
}

func DecryptCBC(key, iv, ct []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("aes: invalid key length")
	}
	if len(ct)%aes.BlockSize != 0 {
		return nil, errors.New("aes-cbc: invalid ciphertext length")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	pt := make([]byte, len(ct))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(pt, ct)
	return pkcs7UnPadding(pt, aes.BlockSize)
}
