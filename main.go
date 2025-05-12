package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func stringToUppercaseHex(str string) string {
	hex := ""
	for i := 0; i < len(str); i++ {
		hex += fmt.Sprintf("%02X", str[i])
	}
	return hex
}

func decryptPayload(payloadBase64 string, rawKey string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(payloadBase64)
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(decodedBytes, &data)
	if err != nil {
		return "", err
	}

	ivB64, _ := data["iv"].(string)
	valueB64, _ := data["value"].(string)

	iv, err := base64.StdEncoding.DecodeString(ivB64)
	if err != nil {
		return "", err
	}

	cipherText, err := base64.StdEncoding.DecodeString(valueB64)
	if err != nil {
		return "", err
	}

	hexKey := stringToUppercaseHex(rawKey)
	key := []byte(hexKey)

	if len(key) != 32 {
		return "", fmt.Errorf("invalid key length: %d bytes. Expected 32", len(key))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(cipherText))
	mode.CryptBlocks(decrypted, cipherText)

	paddingLength := int(decrypted[len(decrypted)-1])
	if paddingLength > 0 && paddingLength <= 16 {
		decrypted = decrypted[:len(decrypted)-paddingLength]
	}

	return string(decrypted), nil
}

func main() {
	encryptedPayloadBase64 := "XXXXXXXXXXXXXXXX"
	rawKey := "XXXXXXXXXXXXXXXXXXXXX"

	fmt.Println("Encrypted Payload:" + encryptedPayloadBase64)

	plaintext, err := decryptPayload(encryptedPayloadBase64, rawKey)
	if err == nil {
		fmt.Println("✅ Decrypted Payload:")
		fmt.Println(plaintext)
	} else {
		fmt.Println("❌ Failed to decrypt:", err)
	}
}
