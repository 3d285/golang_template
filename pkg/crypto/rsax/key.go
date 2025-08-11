package rsax

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func parseRSAPrivateKeyFromPEM(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil { return nil, errors.New("rsa: invalid private pem") }
	if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil { return key, nil }
	k, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil { return nil, err }
	key, ok := k.(*rsa.PrivateKey)
	if !ok { return nil, errors.New("rsa: not private key") }
	return key, nil
}

func parseRSAPublicKeyFromPEM(pemBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil { return nil, errors.New("rsa: invalid public pem") }
	if pk, err := x509.ParsePKIXPublicKey(block.Bytes); err == nil {
		if k, ok := pk.(*rsa.PublicKey); ok { return k, nil }
		return nil, errors.New("rsa: not rsa public key")
	}
	pk, err := x509.ParsePKCS1PublicKey(block.Bytes)
	return pk, err
}
