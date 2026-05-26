package jwtx

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// LoadPrivateKey 从 PKCS#8 PEM 文件加载 RSA 私钥。
// 文件格式与 Java Nimbus 写入的 PKCS#8 PEM 兼容。
func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("jwtx: read private key %q: %w", path, err)
	}
	return ParsePrivateKey(data)
}

// LoadPublicKey 从 X.509 SubjectPublicKeyInfo PEM 文件加载 RSA 公钥。
func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("jwtx: read public key %q: %w", path, err)
	}
	return ParsePublicKey(data)
}

// ParsePrivateKey 解析 PKCS#8（推荐）或 PKCS#1 PEM。
func ParsePrivateKey(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("jwtx: invalid PEM block (private key)")
	}
	if k, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		rk, ok := k.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("jwtx: PKCS#8 key is not RSA")
		}
		return rk, nil
	}
	// 兼容 PKCS#1
	rk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("jwtx: parse RSA private key: %w", err)
	}
	return rk, nil
}

// ParsePublicKey 解析 X.509 SubjectPublicKeyInfo（推荐）或 PKCS#1 PEM。
func ParsePublicKey(pemBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("jwtx: invalid PEM block (public key)")
	}
	if k, err := x509.ParsePKIXPublicKey(block.Bytes); err == nil {
		rk, ok := k.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("jwtx: PKIX key is not RSA")
		}
		return rk, nil
	}
	rk, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("jwtx: parse RSA public key: %w", err)
	}
	return rk, nil
}
