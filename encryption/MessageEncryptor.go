package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

type MessageEncryptor struct {
	KeyDatabase IKeyDatabase
}

func (messageEncryptor MessageEncryptor) EncryptMessage(message string, label string) string {
	messageBytes := []byte(message)
	labelBytes := []byte(label)
	encryptedBytes := messageEncryptor.EncryptBytes(messageBytes, labelBytes)
	encryptedString := base64.StdEncoding.EncodeToString(encryptedBytes)
	return encryptedString
}

func (messageEncryptor MessageEncryptor) EncryptBytes(message []byte, label []byte) []byte {
	publicKey := messageEncryptor.KeyDatabase.GetPublicKey()
	encryptedBytes, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, message, label)
	return encryptedBytes
}

func (messageEncryptor MessageEncryptor) DecryptMessage(message string, label string) string {
	encryptedBytes, _ := base64.StdEncoding.DecodeString(message)
	decryptedBytes := messageEncryptor.DecryptBytes(encryptedBytes, []byte(label))
	decryptedString := string(decryptedBytes)
	return decryptedString
}

func (messageEncryptor MessageEncryptor) DecryptBytes(message []byte, label []byte) []byte {
	privateKey := messageEncryptor.KeyDatabase.GetPrivateKey()
	decryptedBytes, _ := rsa.DecryptOAEP(sha256.New(), rand.Reader, &privateKey, message, label)
	return decryptedBytes
}
