package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func DecryptAES(key string, seed string, encryptedData []byte) ([]byte, error) {

	if len(encryptedData) == 0 {
		return nil, nil
	}

	bKey := []byte(key)
	bSeed := []byte(seed)

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, bSeed)

	if len(encryptedData)%mode.BlockSize() != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	mode.CryptBlocks(encryptedData, encryptedData)
	return pkcs5UnPadding(encryptedData), nil
}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
