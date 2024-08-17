package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func GenerateRsaKey() {
	// return if rsa key exists
	if _, err := os.Stat("./keys/app.rsa"); err == nil {
		return
	}

	keyPath := "./keys"
	filename := "app"
	bitSize := 4096

	// Generate RSA key.
	key, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		panic(err)
	}

	// Extract public component.
	pub := key.Public()

	// Encode private key to PKCS#1 ASN.1 PEM.
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)

	// Encode public key to PKCS#1 ASN.1 PEM.
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pub.(*rsa.PublicKey)),
		},
	)

	// Create directory if it doesn't exist.
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		if err := os.Mkdir(keyPath, 0755); err != nil {
			panic(err)
		}
	}

	// Write private key to file.
	if err := os.WriteFile(keyPath+"/"+filename+".rsa", keyPEM, 0700); err != nil {
		panic(err)
	}

	// Write public key to file.
	if err := os.WriteFile(keyPath+"/"+filename+".rsa.pub", pubPEM, 0755); err != nil {
		panic(err)
	}
}
