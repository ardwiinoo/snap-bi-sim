package signer

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"
	"time"
)

type Service struct {
	privateKey *rsa.PrivateKey
}

func NewService(privateKeyPath string) (*Service, error) {
	keyData, err := os.ReadFile(privateKeyPath)

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)

	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return &Service{
		privateKey: privKey,
	}, nil
}

func (s *Service) GenerateSignature(req SignRequest) (*SignResponse, error) {
	timestamp := time.Now().UTC().Format(time.RFC3339)

	data := req.PartnerID + "|" + timestamp + "|" + req.Body
	hash := sha256.Sum256([]byte(data))

	sig, err := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, hash[:])

	if err != nil {
		return nil, err
	}

	return &SignResponse {
		Timestamp: timestamp,
		Signature: base64.StdEncoding.EncodeToString(sig),
	}, nil
}