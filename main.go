package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func main11() {

	// Generate RSA Keys
	miryanPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	miryanPublicKey := &miryanPrivateKey.PublicKey

	raulPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	raulPublicKey := &raulPrivateKey.PublicKey

	fmt.Println("Private Key : ", miryanPrivateKey)
	fmt.Println("Public key ", miryanPublicKey)
	fmt.Println("Private Key : ", raulPrivateKey)
	fmt.Println("Public key ", raulPublicKey)

	//Encrypt Miryan Message
	message := []byte(`{"time":"2019-06-21T11:34:36.9857706+07:00","id":"","remote_ip":"10.5.21.31","host":"10.5.21.31:8089","method":"POST","uri":"/merchantapi/getbalance","user_agent":"MERCHANT","status":401,"error":"code=401, message=Unauthorized","latency":0,"latency_human":"0s","bytes_in":511,"bytes_out":27}`)
	label := []byte("")
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, raulPublicKey, message, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(message), ciphertext)
	fmt.Println()

	// Message - Signature
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	PSSmessage := message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, miryanPrivateKey, newhash, hashed, &opts)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("PSS Signature : %x\n", signature)

	// Decrypt Message
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, raulPrivateKey, ciphertext, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("OAEP decrypted [%x] to \n[%s]\n", ciphertext, plainText)

	//Verify Signature
	err = rsa.VerifyPSS(miryanPublicKey, newhash, hashed, signature, &opts)

	if err != nil {
		fmt.Println("Who are U? Verify Signature failed")
		os.Exit(1)
	} else {
		fmt.Println("Verify Signature successful")
	}

}