package encription

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
)

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func AESECBEncryption(key string, text string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %v", err)
	}
	mode := ecb.NewECBEncrypter(block)
	pad := padding.NewPkcs7Padding(block.BlockSize())
	data, err := pad.Pad([]byte(text))
	if err != nil {
		return nil, fmt.Errorf("failed to pad: %v", err)
	}
	ciphertext := make([]byte, len(data))
	mode.CryptBlocks(ciphertext, data)
	return ciphertext, nil
}

func AESECBDecryption(key string, text []byte) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}
	mode := ecb.NewECBDecrypter(block)
	pad := padding.NewPkcs7Padding(block.BlockSize())
	plaintext := make([]byte, len(text))
	mode.CryptBlocks(plaintext, text)
	plaintext, err = pad.Unpad(plaintext)
	if err != nil {
		return "", fmt.Errorf("failed to unpad: %v", err)
	}
	return string(plaintext), nil
}
