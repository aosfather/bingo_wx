package apiv3

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

/**
  AES GCM解密方式
*/
type AesUtils struct {
	key []byte
}

func (this *AesUtils) Init(key string) error {
	this.key = []byte(key)
	if len(this.key) != 32 {
		this.key = []byte{}
		return fmt.Errorf("无效的ApiV3Key，长度必须为32个字节")
	}
	return nil
}

//解密文本
func (this *AesUtils) DecryptToString(associatedData, nonce, ciphertext string) (string, error) {

	block := this.getBlock()
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	//base64解码
	cipherbytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	//解密
	plaintext, err := aesgcm.Open(nil, []byte(nonce), cipherbytes, []byte(associatedData))
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func (this *AesUtils) getBlock() cipher.Block {
	block, err := aes.NewCipher(this.key)
	if err != nil {
		panic(err.Error())
	}
	return block
}
