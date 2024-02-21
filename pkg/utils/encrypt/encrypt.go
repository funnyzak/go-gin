package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

type Encrypt struct {
	key []byte
}

func NewEncrypt(key string) (*Encrypt, error) {
	if len(key) < 16 {
		return nil, errors.New("key length must be at least 16")
	}

	keyBytes := []byte(key)

	return &Encrypt{key: keyBytes}, nil
}

func (e *Encrypt) Encrypt(data []byte) ([]byte, error) {
	aesBlock, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aesBlock.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(aesBlock, iv)

	encryptedData := make([]byte, len(data)+aesBlock.BlockSize())
	copy(encryptedData[:aesBlock.BlockSize()], iv)
	stream.XORKeyStream(encryptedData[aesBlock.BlockSize():], data)

	return []byte(base64.StdEncoding.EncodeToString(encryptedData)), nil
}

func (e *Encrypt) Decrypt(data []byte) ([]byte, error) {
	decodedData, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	aesBlock, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	iv := decodedData[:aesBlock.BlockSize()]

	stream := cipher.NewCFBDecrypter(aesBlock, iv)

	decryptedData := make([]byte, len(decodedData)-aesBlock.BlockSize())
	stream.XORKeyStream(decryptedData, decodedData[aesBlock.BlockSize():])

	return decryptedData, nil
}

func Example() {
	encrypt, err := NewEncrypt("my-secret-key")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := []byte("This is a secret message.")
	encryptedData, err := encrypt.Encrypt(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	decryptedData, err := encrypt.Decrypt(encryptedData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(decryptedData))
}
