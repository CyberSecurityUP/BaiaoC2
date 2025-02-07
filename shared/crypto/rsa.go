package crypto

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"encoding/base64" // Adicione esta linha
	"errors"
)

// GenerateRSAKeyPair gera um par de chaves RSA
func GenerateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}


func decodePublicKey(pubKeyStr string) (*rsa.PublicKey, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(pubKeyStr)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(decodedKey)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pub, nil
}

// EncryptWithPublicKey criptografa dados com a chave pública RSA
func EncryptWithPublicKey(data []byte, pubKey *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
}

// DecryptWithPrivateKey descriptografa dados com a chave privada RSA
func DecryptWithPrivateKey(ciphertext []byte, privKey *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privKey, ciphertext)
}

// PublicKeyToBytes converte a chave pública para bytes
func PublicKeyToBytes(pubKey *rsa.PublicKey) []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		panic(err)
	}
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})
}

// BytesToPublicKey converte bytes para chave pública
func BytesToPublicKey(pubKeyBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pubKeyBytes)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pub.(*rsa.PublicKey), nil
}
