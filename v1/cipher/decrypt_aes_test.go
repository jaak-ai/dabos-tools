package cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecryptAES(t *testing.T) {
	// arrange or given
	key := "thisis32bitlongpassphraseimusing"
	seed := "encryptionIntVec"
	cipherData := []byte{
		190, 76, 59, 84, 128, 123, 194, 2, 189, 202, 30, 20, 97, 217, 90, 83, 51, 242, 162, 64, 128, 128, 169, 229, 189, 204, 235, 210, 89, 38, 186, 185,
	}

	// act or when
	data, err := DecryptAES(key, seed, cipherData)

	// assertion or then
	assert.Nil(t, err)
	assert.Equal(t, "This is a secret", string(data))
}
