package database

import (
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/pem"
)

type KeyDatabase struct {
}

//go:embed publickey.txt
var publicKey string

//go:embed privatekey.txt
var privateKey string

func (keyDatabase KeyDatabase) GetPublicKey() rsa.PublicKey { // This file should only handle getting the file, probably not converting
	publicKeyBytes := []byte(publicKey)
	block, _ := pem.Decode(publicKeyBytes)
	pub, _ := x509.ParsePKIXPublicKey(block.Bytes)
	rsaPubKey := pub.(*rsa.PublicKey)
	return *rsaPubKey
}

func (keyDatabase KeyDatabase) GetPrivateKey() rsa.PrivateKey { // This file should only handle getting the file, probably not converting
	privateKeyBytes := []byte(privateKey)
	block, _ := pem.Decode(privateKeyBytes)
	pub, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	//rsaPrivKey := pub.(*rsa.PrivateKey)
	return *pub
}
