package crypto


import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"encoding/hex"
	"crypto/md5"
	"encoding/base64"
)
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}


func TestTripleDES() {
	// because we are going to use TripleDES... therefore we Triple it!
	//triplekey := "12345678" + "12345678" + "12345678"
	key := "v6dfdb4eba53131"

	 // Step 1: Convert it to a rune
	 keyCode := []rune(GetMD5Hash(key))
	 // Step 2: Grab the num of chars you need
	 triplekey := string(keyCode[0:24])	
	fmt.Println(triplekey)
	//plaintext := []byte(`{"pg_user_code":"VIMO","channel_name":"EVN","fnc":"GetBill","data":"{\"customerCode\":\"FAILED123456\",\"serviceCode\":\"BILL_ELECTRIC\",\"billType\":\"TD\",\"zoneCode\":\"12312313\",\"branchCode\":\"1500000\"}","checksum":"8c6c1603ace02edc0866db8181cf2a05"}`)
	plaintext := []byte(`encrypts the given message with`)
	// encrypt
	crypted, err := TripleDesEncrypt(plaintext, []byte(triplekey))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s encrypt to %x \n", string(plaintext[:]), string(crypted[:]))

	fmt.Println(base64.StdEncoding.EncodeToString(crypted))




	decrypted, err := TripleDesDecrypt(crypted, []byte(triplekey))
	fmt.Printf("%x decrypt to %s\n", crypted, decrypted)


	//decrypt
	crypted1 :=`JrHEYGbruRM=`
	ciphertext,_:= base64.StdEncoding.DecodeString(crypted1)
	
	decrypted1, err := TripleDesDecrypt(ciphertext, []byte(triplekey))
	fmt.Printf("%x decrypt to %s\n", crypted, decrypted1)
}

func TripleDesEncrypt(data, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := key
	iv := ciphertext[:des.BlockSize]
	origData := PKCS5Padding(data, block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(origData))
	mode.CryptBlocks(encrypted, origData)
	return encrypted, nil
}




func TripleDesDecrypt(data, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := key
	iv := ciphertext[:des.BlockSize]

	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(data))
	decrypter.CryptBlocks(decrypted, data)
	decrypted = PKCS5UnPadding(decrypted)
	return decrypted, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}