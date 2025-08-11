package crypto

import (
	"golang_template/pkg/crypto/aesx"
	"golang_template/pkg/crypto/rsax"
)

// ---- AES (GCM preferred) ----

type AESGCMConfig struct {
	Key   []byte // 16/24/32
	Nonce []byte // 12 bytes; if nil generated
	AAD   []byte // optional
}

func Encrypt(data []byte, cfg AESGCMConfig) (nonce, ciphertext []byte, err error) {
	return aesx.EncryptGCM(cfg.Key, cfg.Nonce, cfg.AAD, data)
}

func Decrypt(nonce, ciphertext []byte, cfg AESGCMConfig) ([]byte, error) {
	return aesx.DecryptGCM(cfg.Key, nonce, cfg.AAD, ciphertext)
}

// ---- RSA (OAEP encrypt, PSS sign) ----

func RSAEncryptOAEP(pubPEM []byte, msg []byte, label []byte) ([]byte, error) {
	return rsax.EncryptOAEP(pubPEM, msg, label)
}

func RSADecryptOAEP(privPEM []byte, ct []byte, label []byte) ([]byte, error) {
	return rsax.DecryptOAEP(privPEM, ct, label)
}

func RSASignPSS(privPEM []byte, msg []byte) ([]byte, error) {
	return rsax.SignPSS(privPEM, msg)
}

func RSAVerifyPSS(pubPEM []byte, msg, sig []byte) error {
	return rsax.VerifyPSS(pubPEM, msg, sig)
}
