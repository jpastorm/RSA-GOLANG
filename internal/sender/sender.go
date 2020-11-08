package sender

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func Encrypt(publicKeyPath, plainText string) (string, error) {
	bytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return "", err
	}

	publicKey, err := convertBytesToPublicKey(bytes)
	if err != nil {
		return "", err
	}

	cipher, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, publicKey, []byte(plainText), nil)
	if err != nil {
		return "", err
	}

	return cipherToPemString(cipher), nil
}

func convertBytesToPublicKey(keyBytes []byte) (*rsa.PublicKey, error) {
	var err error

	block, _ := pem.Decode(keyBytes)
	blockBytes := block.Bytes
	ok := x509.IsEncryptedPEMBlock(block)

	if ok {
		blockBytes, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}

	publicKey, err := x509.ParsePKCS1PublicKey(blockBytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func cipherToPemString(cipher []byte) string {
	return string(
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "MESSAGE",
				Bytes: cipher,
			},
		),
	)
}
