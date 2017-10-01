package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/sha3"
)

func main() {
	fmt.Println("crypto")
	//	sha256Sample()
	//	rsaSample()
	ed25519Sample()
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

func ed25519Sample() {
	fmt.Println("ed25519 sample")
	signData := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTVUWXYZ0123456789"

	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("pub length:", len(pub))
	fmt.Println("pri length:", len(pri))
	fmt.Printf("pub key:%v\n", pub)
	fmt.Printf("pub key:%x\n", pub)
	fmt.Printf("pri key:%v\n", pri)
	fmt.Printf("pri key:%x\n", pri)
	fmt.Printf("sign [%s]\n", signData)
	sig := ed25519.Sign(pri, []byte(signData))
	fmt.Println("sign result:", sig)
	fmt.Println("verify result:", ed25519.Verify(pub, []byte(signData), sig))
	fmt.Printf("%x\n", sha3.Sum256([]byte(signData)))
}
