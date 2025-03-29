package cryptoUtil

//aes加密

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type AesEncrypt struct {
}

func (tl *AesEncrypt) getKey(strKey string) []byte {
	//strKey := "1234567890123456"
	keyLen := len(strKey)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

// 加密字符串
func (tl *AesEncrypt) Encrypt(strKey, strMesg string) ([]byte, error) {
	key := tl.getKey(strKey)
	fmt.Println("key", string(key))
	var iv = []byte(key)[:aes.BlockSize]
	fmt.Println("key", string(iv))
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	return encrypted, nil
}

// 解密字符串
func (tl *AesEncrypt) Decrypt(strKey string, src []byte) (strDesc []byte, err error) {
	defer func() {
		//错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	key := tl.getKey(strKey)
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return decrypted, nil
}
