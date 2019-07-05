package cryptoaes_test
import (
	"testing"
	"gitlab.saobang.vn/vimo-aritime/go-billing-core/crypto/cryptoaes"
)

func TestEncrpt(t *testing.T) {
    passphrase := "KEY_ENCRYPT"
    text := "Ví điện tử VIMO đã được nạp tiền thành công PY9177000741"
    strout := cryptoaes.Encrypt(text, passphrase)
    // t.Errorf("got: %s", strout)
    if strout != "" {
       t.Errorf("Sum was incorrect, got: %s, want: %s.", strout, "PY9177000741")
    }
}



func TestDecrpt1(t *testing.T) {    
   // strout := cryptoaes.Decrypt("U2FsdGVkX1+/aNrJgD5clcicAYoMrRnkJDo3XJ5QimCzk4uyEntwONKseZtxROAwmjpaBuYO43WcbG9MrV5wdfRUtlOCDcLAL8F2brRRIXq50MjGPzUf/ypHIwnhtIrAdWEVaxOKxHaWmydhEFNixj3l/kdMY6/Oa/FiA93t1sn5fbHf6QUKYytL4Zdm8KGjZ3+yHhS/AckBErhUB7hN6Q==", "dskfadasde4324324")
   passphrase := "KEY_ENCRYPT"
   strout := cryptoaes.Decrypt("U2FsdGVkX18sTER5maewp2sEd7WUr7SY3e9ROPRjwTMsWFx8r11A6i/HONwgdNeW/uZzihY5B4+DloRkPIkk6g==", passphrase)
    t.Logf("Output, got: %s", strout)
    if strout != "" {
       t.Errorf("Output, got: %s, want: %s.", strout, "PY9177000741")
    }
        
}