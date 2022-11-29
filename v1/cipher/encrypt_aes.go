package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func EncryptAES(key string, initVector string, plaintext []byte) ([]byte, error) {

	bKey := []byte(key)
	bIV := []byte(initVector)
	bPlaintext := PKCS5Padding(plaintext, aes.BlockSize)
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return ciphertext, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
