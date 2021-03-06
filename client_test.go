package pangcrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientCryptoGood(t *testing.T) {
	tests := []struct {
		key           uint8
		plain, cipher []byte
	}{
		{
			key: 0,
			plain: []byte{
				0x00, 0x00, 0x02, 0x0D, 0x00, 0x68, 0x65, 0x79,
				0x6C, 0x6C, 0x6F, 0x77, 0x20, 0x77, 0x6F, 0x72,
				0x6C, 0x64, 0x3D, 0x1F, 0x00, 0x00, 0x09, 0x00,
				0x31, 0x32, 0x37, 0x2E, 0x30, 0x2E, 0x30, 0x2E,
				0x31,
			},
			cipher: []byte{
				0x34, 0x22, 0x00, 0x00, 0xF0, 0x00, 0x00, 0x02,
				0xC6, 0x00, 0x68, 0x67, 0x74, 0x6C, 0x04, 0x0A,
				0x0E, 0x4C, 0x1B, 0x00, 0x05, 0x4C, 0x13, 0x52,
				0x6D, 0x6C, 0x64, 0x34, 0x1F, 0x31, 0x32, 0x3E,
				0x2E, 0x01, 0x1C, 0x07, 0x00, 0x01,
			},
		},
		{
			key: 5,
			plain: []byte{
				0x01, 0x00, 0x04, 0x00, 0x6a, 0x6f, 0x68, 0x6e,
				0x20, 0x00, 0x30, 0x39, 0x38, 0x46, 0x36, 0x42,
				0x43, 0x44, 0x34, 0x36, 0x32, 0x31, 0x44, 0x33,
				0x37, 0x33, 0x43, 0x41, 0x44, 0x45, 0x34, 0x45,
				0x38, 0x33, 0x32, 0x36, 0x32, 0x37, 0x42, 0x34,
				0x46, 0x36, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00,
			},
			cipher: []byte{
				0x07, 0x3c, 0x00, 0x00, 0x5b, 0x01, 0x00, 0x04,
				0x47, 0x6b, 0x6f, 0x6c, 0x6e, 0x4a, 0x6f, 0x58,
				0x57, 0x18, 0x46, 0x06, 0x7b, 0x7b, 0x02, 0x02,
				0x74, 0x71, 0x75, 0x70, 0x05, 0x05, 0x02, 0x07,
				0x72, 0x73, 0x76, 0x77, 0x04, 0x7c, 0x76, 0x06,
				0x73, 0x0a, 0x04, 0x70, 0x02, 0x74, 0x01, 0x42,
				0x34, 0x46, 0x36, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			key: 5,
			plain: []byte{
				0x03, 0x00, 0xea, 0x4e, 0x00, 0x00,
			},
			cipher: []byte{
				0x11, 0x07, 0x00, 0x00, 0x6e, 0x03, 0x00, 0xea,
				0xcf, 0x03, 0x00,
			},
		},
	}

	for _, test := range tests {
		decrypted, err := ClientDecrypt(test.cipher, test.key)
		assert.Nil(t, err)
		assert.Equal(t, test.plain, decrypted, "client decrypt")

		encrypted, err := ClientEncrypt(test.plain, test.key, test.cipher[0])
		assert.Nil(t, err)
		assert.Equal(t, test.cipher, encrypted, "client encrypt")
	}
}

func TestInvalidClientKey(t *testing.T) {
	var err error

	_, err = ClientEncrypt([]byte{}, 0x10, 0x00)
	assert.EqualError(t, err, "key 0x10 is too large (maximum key is 0x0f)")

	_, err = ClientDecrypt([]byte{0x00, 0x01, 0x00, 0x00, 0x00}, 0x10)
	assert.EqualError(t, err, "key 0x10 is too large (maximum key is 0x0f)")
}

func TestInvalidClientBuffer(t *testing.T) {
	var err error

	_, err = ClientDecrypt([]byte{}, 0x00)
	assert.EqualError(t, err, "buffer too small (have 0 bytes, need at least 5.)")
}
