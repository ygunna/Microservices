package microservicebroker

import (
	"encoding/base64"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type CredentialsGenerator interface {
	RandomGenerate(string) Credentials
	Generate(string) Credentials
}

type RandomCredentialsGenerator struct{}

type Credentials struct {
	Username string
	Password string
	Password2 string
}

// https://www.ietf.org/rfc/rfc4648.txt - use "Base 64 Encoding with URL and Filename Safe Alphabet"
// No need to worry about '+' and '/' according to the above RFC
func generateString(size int) string {
	rb := make([]byte, size)
	_, err := rand.Read(rb)

	if err != nil {
		fmt.Println(err)
	}
	return base64.URLEncoding.EncodeToString(rb)
}

func (s RandomCredentialsGenerator) RandomGenerate(prefix string) Credentials {
	return Credentials{
		Username: prefix + generateString(12),
		Password: generateString(24),
		Password2: generateString(24),
	}
}


func (s RandomCredentialsGenerator) Generate(data string) Credentials {
	return Credentials{
		Username: data,
		Password: generateHash(data, "password"),
		Password2: generateString(24), // random password
	}
}

func generateHash(data, key string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
		return nil
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
	//return fmt.Sprintf("%x", ciphertext)
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return plaintext
	//return fmt.Sprintf("%s", plaintext)
}