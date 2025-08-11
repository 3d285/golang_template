package rsax

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func SignPSS(privPEM, msg []byte) ([]byte, error) {
	priv, err := parseRSAPrivateKeyFromPEM(privPEM)
	if err != nil { return nil, err }
	h := sha256.Sum256(msg)
	return rsa.SignPSS(rand.Reader, priv, crypto.SHA256, h[:], nil)
}

func VerifyPSS(pubPEM, msg, sig []byte) error {
	pub, err := parseRSAPublicKeyFromPEM(pubPEM)
	if err != nil { return err }
	h := sha256.Sum256(msg)
	return rsa.VerifyPSS(pub, crypto.SHA256, h[:], sig, nil)
}
