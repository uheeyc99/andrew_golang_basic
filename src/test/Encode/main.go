package main
//http://blog.csdn.net/qq_34645412/article/details/77305456
import (
	"fmt"
	"encoding/base64"
	"time"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
)



func base64encode(src []byte) []byte{
	t1:=time.Now()
	value:=base64.StdEncoding.EncodeToString(src)
	fmt.Println("encode time:",time.Now().Sub(t1))
	return []byte(value)
}
func base64decode(src []byte) []byte{
	t1:=time.Now()
	value,_:=base64.StdEncoding.DecodeString(string(src))
	fmt.Println("decode time:",time.Now().Sub(t1))
	return value
}

func test1()  {  //base64
	var str="hello"
	fmt.Println(str)
	encoded:=base64encode([]byte(str))
	fmt.Println(encoded)
	decode:=base64decode(encoded)
	fmt.Println(decode,string(decode))

}

type AesEncrypt struct {
}


func (this *AesEncrypt) getKey() []byte {
	key := md5.Sum([]byte("ssss"))
	//keyLen := len(strKey)
	//if keyLen < 16 {
	//	panic("res key 长度不能小于16")
	//}
	//arrKey := []byte(strKey)
	//if keyLen >= 32 {
	//	return arrKey[:32]
	//}
	//if keyLen >= 24 {
	//	return arrKey[:24]
	//}
	//return arrKey[:16]
	return key[0:16]
}

//加密字符串
func (this *AesEncrypt) Encrypt(strMesg string) ([]byte, error) {
	key := this.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipher.NewCFBEncrypter(aesBlockEncrypter, iv).XORKeyStream(encrypted, []byte(strMesg))
	return encrypted, nil
}

//解密字符串
func (this *AesEncrypt) Decrypt(src []byte) (strDesc string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	key := this.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	cipher.NewCFBDecrypter(aesBlockDecrypter, iv).XORKeyStream(decrypted,src)
	return string(decrypted), nil
}





func test2(str string)  { //aes

	aesEnc := AesEncrypt{}
	t1:=time.Now()
	arrEncrypt, err := aesEnc.Encrypt(str)
	if err != nil {
		fmt.Println(arrEncrypt)
		return
	}
	fmt.Println("encode time:",time.Duration(time.Now().Sub(t1)),len(arrEncrypt))

	t2:=time.Now()
	strMsg, err := aesEnc.Decrypt(arrEncrypt)
	if err != nil {
		fmt.Println(arrEncrypt)
		return
	}

	fmt.Println("decode time:",time.Duration(time.Now().Sub(t2)))
	fmt.Println(strMsg)


}

func test3(str string){  //md5
	 checkSum:=md5.Sum([]byte(str))
	 fmt.Println(checkSum,len(checkSum))
}


func main(){
	//test1()  //数据被转换成字符串而已
	//test2("abcdefg")  //对称加密
	test3("qwqqqqqqqqqwwqwqwqwqwqdsadsafdfdfdfdfd")
}
