package encryption

import (
	"isc/database"
	"testing"
)

func TestEncryptionAndDecryption(test *testing.T) {
	messageToEncrypt := "testing"

	messageEncryptor := MessageEncryptor{database.GetPublicKey()}
	encryptedMessage := messageEncryptor.EncryptMessage(messageToEncrypt, "")
	decryptedMessage := messageEncryptor.DecryptMessage(encryptedMessage, "", database.GetPrivateKey())

	if messageToEncrypt != decryptedMessage {
		test.Fatalf("Decrypted message not the same as text before encryption")
	}
}
