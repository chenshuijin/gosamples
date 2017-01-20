package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("crypto")
	sha256Sample()
	rsaSample()
}

func sha256Sample() {
	fmt.Println("sha256 sample")
	s := "come on for test sha256"
	fmt.Printf("[%s] sha256 is [%x]\n", s, sha256.Sum256([]byte(s)))
}

func rsaSample() {
	fmt.Println("rsa sample")
	keySize := 1024
	pri, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		fmt.Println("error:", err)
	}
	if pri.N.BitLen() != keySize {
		fmt.Printf("key too short wants %s but get %s\n", keySize, pri.N.BitLen())
	}
	fmt.Printf("private key:[%x]\n", pri)
	fmt.Printf("public key:[%x]\n", pri.Public())
}
