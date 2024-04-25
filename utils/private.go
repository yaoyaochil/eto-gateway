package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"gateway/global"
	"go.uber.org/zap"
	"os"
)

func LoadPrivateKey(path string) (string, error) {
	pemData, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(pemData), nil
}

// EncryptBase64WithUID 使用 RSA 公钥加密 UID 并返回 Base64 编码的字符串
func EncryptBase64WithUID(uid int32) (string, error) {
	uidStr := fmt.Sprintf("%08x010101010101010101010101010101010101010101010101010101010101010155914510010403030101", uid)
	data := hex2bin(uidStr)
	privateKey, err := LoadPrivateKey("private.pem")
	if err != nil {
		global.GatewayLog.Error("Failed to load private key:", zap.Error(err))
		return "", err
	}

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		global.GatewayLog.Error("no key found")
		return "", fmt.Errorf("no key found")
	}

	if block.Type != "RSA PRIVATE KEY" {
		global.GatewayLog.Error("unsupported key type", zap.String("type", block.Type))
		return "", fmt.Errorf("unsupported key type %q", block.Type)
	}

	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// 使用私钥签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.Hash(0), data)
	if err != nil {
		global.GatewayLog.Error("Failed to sign", zap.Error(err))
		return "", err
	}

	// 返回签名的 Base64 编码
	return base64.StdEncoding.EncodeToString(signature), nil
}

func hex2bin(hexStr string) []byte {
	data, _ := hex.DecodeString(hexStr)
	return data
}
