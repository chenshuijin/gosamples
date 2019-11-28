package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/nacl/box"
	"golang.org/x/crypto/sha3"
)

func main() {
	fmt.Println("crypto")
	//	sha256Sample()
	//	rsaSample()
	//	ed25519Sample()
	//	ed25519Exec()
	ecdhSample()
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
		fmt.Printf("key too short wants %d but get %d\n", keySize, pri.N.BitLen())
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

func ed25519Exec() {
	pk := ""
	sk := ""
	signature := ""
	data := ""
	fmt.Println("input pub")
	fmt.Scanln(&pk)
	fmt.Println("input sk")
	fmt.Scanln(&sk)
	fmt.Println("input data")
	fmt.Scanln(&data)
	fmt.Println("input data:", data)
	fmt.Println("input signature")
	fmt.Scanln(&signature)
	pkbs, err := base64.StdEncoding.DecodeString(pk)
	if err != nil {
		fmt.Println("decode pk err:", err)
	}
	fmt.Println("decode pk:", pkbs)

	signaturebs, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println("decode signature err:", err)
	}
	fmt.Println("decode signature:", signaturebs)
	fmt.Println("verify result:", ed25519.Verify(pkbs, []byte(data), signaturebs))
	for {
		fmt.Println("input data")
		fmt.Scanln(&data)
		fmt.Println("input signature")
		fmt.Scanln(&signature)
		signaturebs, err := base64.StdEncoding.DecodeString(signature)
		if err != nil {
			fmt.Println("decode signature err:", err)
		}
		fmt.Println("decode signature:", signaturebs)
		fmt.Println("verify result:", ed25519.Verify(pkbs, []byte(data), signaturebs))
	}
}

func ecdhSample() {
	senderPublicKey, senderPrivateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	recipientPublicKey, recipientPrivateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	// You must use a different nonce for each message you encrypt with the
	// same key. Since the nonce here is 192 bits long, a random value
	// provides a sufficiently small probability of repeats.
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	msg := []byte("Alas, poor Yorick! I knew him, Horatio")
	// This encrypts msg and appends the result to the nonce.
	encrypted := box.Seal(nonce[:], msg, &nonce, recipientPublicKey, senderPrivateKey)

	// The recipient can decrypt the message using their private key and the
	// sender's public key. When you decrypt, you must use the same nonce you
	// used to encrypt the message. One way to achieve this is to store the
	// nonce alongside the encrypted message. Above, we stored the nonce in the
	// first 24 bytes of the encrypted text.
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])
	decrypted, ok := box.Open(nil, encrypted[24:], &decryptNonce, senderPublicKey, recipientPrivateKey)
	if !ok {
		panic("decryption error")
	}
	fmt.Println(string(decrypted))
	// Output: Alas, poor Yorick! I knew him, Horatio
}
