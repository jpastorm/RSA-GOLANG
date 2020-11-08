package reciever

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func Decrypt(privateKeyPath, encryptedMessage string) (string, error) {
	bytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return "", err
	}

	privateKey, err := convertBytesToPrivateKey(bytes)
	if err != nil {
		return "", err
	}

	plainMessage, err := rsa.DecryptOAEP(
		sha512.New(),
		rand.Reader,
		privateKey,
		pemStringToCipher(encryptedMessage),
		nil,
	)

	return string(plainMessage), err
}

func convertBytesToPrivateKey(keyBytes []byte) (*rsa.PrivateKey, error) {
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

	privateKey, err := x509.ParsePKCS1PrivateKey(blockBytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func pemStringToCipher(encryptedMessage string) []byte {
	b, _ := pem.Decode([]byte(encryptedMessage))

	return b.Bytes
}
