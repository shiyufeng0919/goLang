package 算法

import (
	"os"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
)

/*
对称加密算法--高级加解密  ---AES  and  DEA

1.AES（Advanced Encryption standard）
crypto/aes包,又称Rijindael加密法
是美国联邦政府采用的一种区块加密标准
2.DEA（Data Encryption Algorithm）
crypto/des包，目前使用最广泛的密钥系统，特别是在保护金融数据的安全中
*/

/*
AES对称加密算法
*/
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
	0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
func AESDemo() {
	//需要去加密的字符串
	plaintext := []byte("My name is yufeng")
	//如果传入加密串的话，plaint就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}
	//aes的加密字符串
	key_text := "astaxie12798akljzmknm.ahkjkljl;k"
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}
	fmt.Println(len(key_text))
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}
	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)
	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
}