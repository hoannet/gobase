package main

import (	
	"gitlab.saobang.vn/vimo-aritime/go-billing-core/logger"	
	//"gitlab.saobang.vn/vimo-aritime/go-billing-core/untils/crypto"
	//"gitlab.saobang.vn/vimo-aritime/go-billing-core/untils/crypto/xrsa"
	"gitlab.saobang.vn/vimo-aritime/go-billing-core/crypto/xrsa"
	"log"
	"bytes"
	"io/ioutil"
)

// var pri_key string
// var pub_key string

func main1() {
	log.Println("Start")
	logger.Infof("Test:"," Content ")

	//TestCreate()
	TestUse()

	

}

func TestCreate(){
	
	
	 //crypto.TestRSA()
	//crypto.TestOpenssl()
	//crypto.TestTripleDES()
	publicKey := bytes.NewBufferString("")
    privateKey := bytes.NewBufferString("")

    err := xrsa.CreateKeys(publicKey, privateKey, 2048)
    if err != nil {
        return
	}

	log.Println(publicKey.String())

	log.Println(privateKey.String())


    xrsa, err := xrsa.NewXRsa(publicKey.Bytes(), privateKey.Bytes())
    if err != nil {
        return
    }

    data := `{"pg_user_code":"VIMO","channel_name":"VPAY","fnc":"BuyCard","data":"{\"transaction_id\":\"dfacb83a-fd40-5238-aae6-966ff69876a7\",\"quantity\":3,\"product_code\":9}","checksum":"db40046630e8775feae37feaf1656919"}`
	encrypted, _ := xrsa.PublicEncrypt(data)
	log.Println(publicKey.String())
	decrypted, _ := xrsa.PrivateDecrypt(encrypted) 
	log.Println(decrypted)
    sign, err := xrsa.Sign(data)
	err = xrsa.Verify(data, sign)
}

func TestUse(){
	
	//private_mctest, err := ioutil.ReadFile("pri_new.pem") // just pass the file name
	//private_mctest, err := ioutil.ReadFile("private_mctest.pem") // just pass the file name
	private_mctest, err := ioutil.ReadFile("privateKey.pem") // just pass the file name
    if err != nil {
        log.Println(err)
	}

	//public_mctest, err := ioutil.ReadFile("pub_new.pem") // just pass the file name
	//public_mctest, err := ioutil.ReadFile("public_mctest.pem") // just pass the file name
	public_mctest, err := ioutil.ReadFile("public.pem") // just pass the file name
    if err != nil {
        log.Println(err)
	}


	private_mctest = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC0kJ5AtebuGF7E
TdPapwr7Z0fmim9lPyLM4cXrufv8pCsSSnA7lzukuEjrlTyOfDlypvESrgwFtTwR
xoTtBLs2x3QQgcb981B8FeQ99RP62NMUM6UfS2HjPLWJX1R0aoJ0Zk1Ir6XU1RzE
OjlsGVVAeF/9CiYm4iTXRg161q8bSXZsTyrW+mzLVv8A0patxJRRM25FUhgUzZvD
PaViF1jSnto+jNAk46nAExGn3OK7//HphBDM4iDIw7xfA+GZC5FQwn5jWCFzYYt4
ImNx0iVmE/hQP/ob978SwvtyWzIRe7iybvAVk9kJeME7s9u0zNBY3+JL/k5Utt32
36KSsnIVAgMBAAECggEAR9MC8DCrRGhzQdjHfcFWj/xcUGsetoJvwFnKleLsRmmz
LFpp3Hpi4jnf374EUZU+cMnROSH8bETUOSnjdDiek8Kw+2P0GM4xZLp8iNy07uJS
zEVi2Fju/ioG+DbwbEJ11AFhmpkr+MYw5ClY8BP/ol5/qDLiTDNGfk9CS2ucSdi6
0HRk3Bn5ySt7uGHXAtOqg+LNM2qoPovQDhvFKGxEOj6Yk+xaxpCwVLI1uEjm5NAY
eo9lZ7LIyCr7Rh96N3JTk+Ts1nihbBbdqpplZa0eBEBkeKRTZvL+D4lKsjg70G0+
S6FQ1mzfAQCO/7HqkfBvOTrbbe8WE4YaavuJy4Lf3QKBgQDf1LOAIyeeRJkwplk7
vPTCSib0gxJZ+mVAt93I/oJQYKcii4yh9xNPA6zw38snXNNsrG6hGmMu+SDvOm0I
CoVAd8gL0JYKxLht4bnnibyroxUE/1T4VylUEgrJBEDbdaDvn1lnUD8rXlEvNJT8
eD4Pl+qcDoHaYLiqjbGipFghmwKBgQDOhA3hl4YIzB4kawPxPY8aR0JbjM83rIUu
eMglzVcj77m5nm2jWmR91se8VjWBOFDfIBOPnxbqZbqDNNgNSivm5/+EcSOuhfab
u1nIiDSMpuWHaZ4V78FqhDUElbpHIDgWwLKGA5BFtt1KjgAaYmd7FV7unhdV2Ti9
LRTGKzYODwKBgEkENwYu8aaklCCVs3hUXLfgvrJ2646kq1Egad4+gD4+OWeAhID1
e7wD4++z4a7WjjXeTjyJwh85r++6bIT6AZwVNxH7mmaq48scnquUeBK5oYq+zHYy
4M2HETOgCpRM3BIvNlgIioLYQr2Wdp+hjgVCNotwBBY8BSHGFVOMLx/PAoGASE8e
mSW37a6iwzehSr+2MopO5sVX7POAMnRiCDJs141hp2eclfzZgOJKT1yUIWm5j1ao
4rneiTYTZ0uCdqzI0HHdGt+OkdYgMgkYd8mRNMpfRdmqW532SDiAY9mVPsA+Q2E3
YQswYeiUdT1hNdvMxeEGu2ApZYrIfhyh0H1i0OsCgYEArYAaWgHnK9UjrApa0sta
g7Kc4XorQeX+EoZgbUonaAtIjK8HZe7ODU4XoejypeXsmIs+JCjWaAINXothMT76
Fueu2DUSegWkn0yZFbgqzI5+JKjJruKFPw22ZW/T0guiCmyNN2zvmGXmAaTEIWh1
61W+6R4aFHx3SIaeOxNP+BM=
-----END PRIVATE KEY-----`)
public_mctest = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtJCeQLXm7hhexE3T2qcK
+2dH5opvZT8izOHF67n7/KQrEkpwO5c7pLhI65U8jnw5cqbxEq4MBbU8EcaE7QS7
Nsd0EIHG/fNQfBXkPfUT+tjTFDOlH0th4zy1iV9UdGqCdGZNSK+l1NUcxDo5bBlV
QHhf/QomJuIk10YNetavG0l2bE8q1vpsy1b/ANKWrcSUUTNuRVIYFM2bwz2lYhdY
0p7aPozQJOOpwBMRp9ziu//x6YQQzOIgyMO8XwPhmQuRUMJ+Y1ghc2GLeCJjcdIl
ZhP4UD/6G/e/EsL7clsyEXu4sm7wFZPZCXjBO7PbtMzQWN/iS/5OVLbd9t+ikrJy
FQIDAQAB
-----END PUBLIC KEY-----`)

	log.Println(string(public_mctest))

	log.Println(string(private_mctest))


	//xrsa, err := xrsa.NewXRsa(public_mctest, private_mctest)
	xrsa, err := xrsa.NewPubXRsa(public_mctest)
    if err != nil {
        log.Println(err.Error())
    }

    data := `{"pg_user_code":"VIMO","channel_name":"VPAY","fnc":"BuyCard","data":"{\"transaction_id\":\"dfacb83a-fd40-5238-aae6-966ff69876a7\",\"quantity\":3,\"product_code\":9}","checksum":"db40046630e8775feae37feaf1656919"}`
	encrypted, _ := xrsa.PublicEncrypt(data)
	log.Println(encrypted)
	// encrypted :=`Smp9N9zDarQC4UyoYviZFn3UIct8ujiRtqM5aJ8CB0ZrU4jMIpgOBSGXsBNE5bnwVFTNdvyAWl7CYUL80wp11vM9bL7wUzfbNuEsnSjm8acd-0HUUU_gMSqfRjgUR5KX-GsVk7vz8rpxgk1z4sjEK1-spR1IJQfnAu_4HaGRq9_miTOtGo5Yo83bsl3ixrvh9VS6fGansl6FEvuN5BZfNXJ3ye1XGPUjlCmqYMZskggExIz6yGpqGAfOdR9nuFYd18SdFHppwCvGO5ZWyqySdGEcFhu8UoS9eDjNfhWtvlxYL26EiO_1F3RTjRhwptHBt0qiWRA0Qrnv-VEULItbxA`
	//  decrypted, _ := xrsa.PrivateDecrypt(encrypted) 
	// log.Println(decrypted)
    // sign, err := xrsa.Sign(data)
    // err = xrsa.Verify(data, sign)
}
