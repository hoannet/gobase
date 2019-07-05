/*
 * Genarate rsa keys.
 */

 package crypto

 import (	
	"io/ioutil"
	"encoding/base64"
	//"encoding/hex"
	"crypto/sha256"
	 "crypto/rand"
	 "crypto/rsa"
	 "crypto/x509"
	 "encoding/asn1"
	 "encoding/gob"
	 "encoding/pem"
	 "fmt"
	 "log"
	 "os"
 )
 
 func TestRSA() {
	private_name :="private_mctest.pem";
	public_name :="public_mctest.pem";
	
	
	
	//  reader := rand.Reader
	//  //bitSize := 2048
	//  bitSize := 1024 
	 
	//  key, err := rsa.GenerateKey(reader, bitSize)
	 
	//  checkError(err)

	 
	//  log.Println(string(x509.MarshalPKCS1PublicKey(&key.PublicKey)))


	// //  log.Println(string(x509.MarshalPKCS1PrivateKey(key)))
	// //  log.Println(key.PublicKey.N)

	 
	//  //saveGobKey("private.key", key)
	//  savePEMKey(private_name, key) 
	//  //saveGobKey("public.key", publicKey)
	//  savePublicPEMKey(public_name, key.PublicKey)

	// panic("END")

	
	publicKeyData, err := ioutil.ReadFile(public_name)
	if err != nil {
		log.Println("Load private key error")
		//panic(err)
	}

	log.Println(string(publicKeyData))



	publicKeyBlock, _ := pem.Decode([]byte(publicKeyData))
	var prub *rsa.PublicKey
	prub, parseErr1 := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
	if parseErr1 != nil {
		fmt.Println("Load public key error")
		panic(parseErr1)
	}
	
	encryptedData := EncryptOAEP(prub)
	log.Println(encryptedData)
	
	
	
	//encryptedData = `YfT2k0hWQ1Drni1SzrNDUBWHXnzL6v+KFPh47pp28H15kE/eC0d+t/c4VlDUL39nUzvcHeJgvbciPkDGOc/SHRw8QazasvQiHfb/YVqZ8bP6k91Q+Q1WFYru3LkgdYKPikNhDk0V+VqhThc88iTW86/j8SPuf6MiWfnckfz/KBYhWYi+tH6Yn8zT+x+FAo6zS1LO6J1Ef2Al3X5gQBrWhsD1/ezPnugANilJWV0ELz3z8V/erVhUfiG6jOITzGSe6MfV2dcUqcUF0eGqU1kANd+nyznK/heaxQcOB56F0rDZFIecBv8MpnRK6OD7yKpjuPVCtC5fUAf+e5rj/K1P3g==`

	privateKeyData, err := ioutil.ReadFile(private_name)
	
	if err != nil {
		log.Println("Load private key error")
		//panic(err)
	}
	privateKeyBlock, _ := pem.Decode([]byte(privateKeyData))
	var pri *rsa.PrivateKey
	
	pri, parseErr := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if parseErr != nil {
		fmt.Println("Load private key error")
		panic(parseErr)
	}
	DecryptOAEP(encryptedData,pri)



	
 }
 

 func EncryptOAEP(publicKey *rsa.PublicKey) string{
	secretMessage := []byte("rypto/rand.Reader is a good source of entropy for randomizing the")
	label := []byte("orders")

	// crypto/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, publicKey, secretMessage, label)
	if err != nil {
			fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
			return ""
	}

	// Since encryption is a randomized function, ciphertext will be
	// different each time.
	//fmt.Printf("Ciphertext: %x\n", ciphertext)
	return base64.StdEncoding.EncodeToString(ciphertext) //hex.EncodeToString(ciphertext)
 }


 func DecryptOAEP(value string,privateKey *rsa.PrivateKey){
	//ciphertext, _ := hex.DecodeString("4d1ee10e8f286390258c51a5e80802844c3e6358ad6690b7285218a7c7ed7fc3a4c7b950fbd04d4b0239cc060dcc7065ca6f84c1756deb71ca5685cadbb82be025e16449b905c568a19c088a1abfad54bf7ecc67a7df39943ec511091a34c0f2348d04e058fcff4d55644de3cd1d580791d4524b92f3e91695582e6e340a1c50b6c6d78e80b4e42c5b4d45e479b492de42bbd39cc642ebb80226bb5200020d501b24a37bcc2ec7f34e596b4fd6b063de4858dbf5a4e3dd18e262eda0ec2d19dbd8e890d672b63d368768360b20c0b6b8592a438fa275e5fa7f60bef0dd39673fd3989cc54d2cb80c08fcd19dacbc265ee1c6014616b0e04ea0328c2a04e73460")
	
	ciphertext,_:= base64.StdEncoding.DecodeString(value)
	// ciphertext, err := hex.DecodeString(data)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error from DecodeString: %s\n", err)
	// 	return
	// }	
	label := []byte("orders")
	
	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	rng := rand.Reader
	
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, privateKey, ciphertext, label)
	if err != nil {
			fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
			return
	}
	
	fmt.Printf("Plaintext: %s\n", string(plaintext))
 }





 func saveGobKey(fileName string, key interface{}) {
	 outFile, err := os.Create(fileName)
	 checkError(err)
	 defer outFile.Close()
 
	 encoder := gob.NewEncoder(outFile)
	 err = encoder.Encode(key)
	 checkError(err)
 }
 



 func savePEMKey(fileName string, key *rsa.PrivateKey) {
	 outFile, err := os.Create(fileName)
	 checkError(err)
	 defer outFile.Close()
 
	 var privateKey = &pem.Block{
		 Type:  "PRIVATE KEY",
		 Bytes: x509.MarshalPKCS1PrivateKey(key),
	 }
 
	 err = pem.Encode(outFile, privateKey)
	 checkError(err)
 }
 
 func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
	 asn1Bytes, err := asn1.Marshal(pubkey)
	 checkError(err)

	 var pemkey = &pem.Block{
		 Type:  "PUBLIC KEY",
		 Bytes: asn1Bytes,
	 }
 
	 pemfile, err := os.Create(fileName)
	 checkError(err)
	 defer pemfile.Close()
 
	 err = pem.Encode(pemfile, pemkey)
	 checkError(err)
 }
 
 func checkError(err error) {
	 if err != nil {
		 fmt.Println("Fatal error ", err.Error())
		 os.Exit(1)
	 }
 }