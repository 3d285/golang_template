package rsax

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func EncryptOAEP(pubPEM, msg, label []byte) ([]byte, error) {
	pub, err := parseRSAPublicKeyFromPEM(pubPEM)
	if err != nil { return nil, err }
	h := sha256.New()
	return rsa.EncryptOAEP(h, rand.Reader, pub, msg, label)
}

func DecryptOAEP(privPEM, ct, label []byte) ([]byte, error) {
	priv, err := parseRSAPrivateKeyFromPEM(privPEM)
	if err != nil { return nil, err }
	h := sha256.New()
	return rsa.DecryptOAEP(h, rand.Reader, priv, ct, label)
}
