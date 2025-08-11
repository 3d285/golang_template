package aesx

import "errors"

func pkcs7Pad(src []byte, blockSize int) []byte {
	pad := blockSize - len(src)%blockSize
	out := make([]byte, len(src)+pad)
	copy(out, src)
	for i := len(src); i < len(out); i++ {
		out[i] = byte(pad)
	}
	return out
}

func pkcs7UnPadding(src []byte, blockSize int) ([]byte, error) {
	if len(src) == 0 || len(src)%blockSize != 0 {
		return nil, errors.New("pkcs7: invalid length")
	}
	pad := int(src[len(src)-1])
	if pad == 0 || pad > blockSize || pad > len(src) {
		return nil, errors.New("pkcs7: invalid padding")
	}
	for i := 0; i < pad; i++ {
		if src[len(src)-1-i] != byte(pad) {
			return nil, errors.New("pkcs7: invalid padding byte")
		}
	}
	return src[:len(src)-pad], nil
}
