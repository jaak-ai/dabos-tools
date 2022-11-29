package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func DecryptAES(key string, initVector string, cipherTextDecoded []byte) ([]byte, error) {

	if len(cipherTextDecoded) == 0 {
		return nil, nil
	}

	bKey := []byte(key)
	bIV := []byte(initVector)

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, bIV)

	if len(cipherTextDecoded)%mode.BlockSize() != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	mode.CryptBlocks(cipherTextDecoded, cipherTextDecoded)
	return PKCS5UnPadding(cipherTextDecoded), nil
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
