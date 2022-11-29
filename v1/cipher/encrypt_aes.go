package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func EncryptAES(key string, seed string, data []byte) ([]byte, error) {

	bKey := []byte(key)
	bSeed := []byte(seed)
	bData := pkcs5Padding(data, aes.BlockSize)
	block, err := aes.NewCipher(bKey)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, len(bData))
	mode := cipher.NewCBCEncrypter(block, bSeed)
	mode.CryptBlocks(ciphertext, bData)
	return ciphertext, nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
