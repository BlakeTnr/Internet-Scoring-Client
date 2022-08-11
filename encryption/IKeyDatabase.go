package encryption

import "crypto/rsa"

type IKeyDatabase interface {
	GetPublicKey() rsa.PublicKey
	GetPrivateKey() rsa.PrivateKey // Unclear as database always have to have a private key
}
