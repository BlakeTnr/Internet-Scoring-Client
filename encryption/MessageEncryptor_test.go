package encryption

import (
	"testing"
)

func TestEncryptionAndDecryption(test *testing.T) {
	messageToEncrypt := "testing"

	keyDatabase := GetKeyDatabase()
	messageEncryptor := MessageEncryptor{keyDatabase}
	encryptedMessage := messageEncryptor.EncryptMessage(messageToEncrypt, "")
	decryptedMessage := messageEncryptor.DecryptMessage(encryptedMessage, "")

	if messageToEncrypt != decryptedMessage {
		test.Fatalf("Decrypted message not the same as text before encryption")
	}
}
