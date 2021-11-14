package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"log"
)

func DoRSASigning(sent, received string) {
	fmt.Println("Исходный текст:", sent)
	text := []byte(sent)
	rng := rand.Reader
	pk, _ := rsa.GenerateKey(rng, 512)
	hash := sha256.Sum256(text)
	fmt.Printf("Хэш исходного текста:\n%x\n", hash)
	sign, err := rsa.SignPKCS1v15(nil, pk, crypto.SHA256, hash[:])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Подпись исходного текста:\n%x\n", sign)
	fmt.Println("Исходный текст:", sent)
	afterhash := sha256.Sum256([]byte(received))
	fmt.Printf("Хэш полученного текста:\n%x\n", afterhash)
	if rsa.VerifyPKCS1v15(&pk.PublicKey, crypto.SHA256, afterhash[:], sign) != nil {
		fmt.Println("НЕ СОВПАЛО")
	} else {
		fmt.Println("СОВПАЛО")
	}
}
