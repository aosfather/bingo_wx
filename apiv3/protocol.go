package apiv3

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/**
  api v3 协议
*/
type ApiV3 struct {
	//服务商商户号
	MerchantId string
	//服务商证书序号
	merchantCertificateSerialNo string
	//服务商私钥
	merchantPrivatekey *rsa.PrivateKey
	//平台证书序号
	flatCertificateSerialNo string
	//平台证书-公钥
	flatCertificate *rsa.PublicKey
	Nonce           string
	aes             *AesUtils
	client          *http.Client
}

//设置商户秘钥信息
func (this *ApiV3) SetMerchantKey(cno string, privatekey []byte) error {
	this.merchantCertificateSerialNo = cno
	block, _ := pem.Decode(privatekey)
	if block == nil {
		return fmt.Errorf("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	this.merchantPrivatekey = priv.(*rsa.PrivateKey)
	return nil
}

//设置平台秘钥信息
func (this *ApiV3) SetFlatKey(cno string, publickey []byte) error {
	this.flatCertificateSerialNo = cno
	block, _ := pem.Decode(publickey)
	if block == nil {
		return fmt.Errorf("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}
	this.flatCertificate = pubInterface.PublicKey.(*rsa.PublicKey)

	return nil
}

//设置APIV3的秘钥
func (this *ApiV3) SetApiV3Key(key string) {
	if key != "" {
		this.aes = &AesUtils{}
		err := this.aes.Init(key)
		if err != nil {
			fmt.Println(err.Error())
			this.aes = nil
		}
	}
}
func (this *ApiV3) Call(method string, url string, parameter string) ([]byte, error) {
	if this.client == nil {
		this.client = &http.Client{}
	}

	req, err := http.NewRequest(method, url, strings.NewReader(parameter))
	if err != nil {
		return nil, err
	}
	//设置签名头
	token := this.getToken(method, url, parameter)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "WECHATPAY2-SHA256-RSA2048 "+token)

	resp, err := this.client.Do(req)
	if !this.validate(resp.Header) {
		return nil, fmt.Errorf("微信平台返回的消息校验不通过，请检查网络是否被劫持！")
	}
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println(resp)
	//	return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", url, resp.StatusCode)
	//}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

//检查消息头
func (this *ApiV3) validate(h http.Header) bool {
	nonce := h["Wechatpay-Nonce"]
	sign := h["Wechatpay-Signature"]
	serial := h["Wechatpay-Serial"]
	if this.aes != nil {
		_, err := this.aes.DecryptToString(serial[0], nonce[0], sign[0])
		if err != nil {
			return false
		}
	}

	return true

}

func (this *ApiV3) DoGet(url string) ([]byte, error) {
	return this.Call("GET", url, "")
}

func (this *ApiV3) DoPost(url string, body string) ([]byte, error) {
	return this.Call("POST", url, "")
}

func (this *ApiV3) getToken(method string, url string, body string) string {
	var t int64 = time.Now().Unix()
	msg := this.buildSignatureMsg(method, url, t, body)
	signer := this.sign(msg)
	return fmt.Sprintf("mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\"", this.MerchantId, this.Nonce, t, this.merchantCertificateSerialNo, signer)
}

func (this *ApiV3) buildSignatureMsg(method string, urlstr string, timestamp int64, body string) string {
	u, err := url.Parse(urlstr)
	if err != nil {

	}
	canonicalUrl := u.RequestURI()
	return fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, canonicalUrl, timestamp, this.Nonce, body)

}

func (this *ApiV3) sign(msg string) string {
	h := sha256.New()
	h.Write([]byte(msg))
	digest := h.Sum(nil)
	s, err := rsa.SignPKCS1v15(rand.Reader, this.merchantPrivatekey, crypto.SHA256, digest)
	if err != nil {
		fmt.Println("rsaSign SignPKCS1v15 error", err.Error())
		return ""
	}
	return base64.StdEncoding.EncodeToString(s)
}

/**
  加密敏感信息
*/
func (this *ApiV3) EncryptText(secretMessage string) (string, error) {
	rng := rand.Reader
	cipherdata, err := rsa.EncryptOAEP(sha1.New(), rng, this.flatCertificate, []byte(secretMessage), nil)
	if err != nil {
		return "", err
	}

	ciphertext := base64.StdEncoding.EncodeToString(cipherdata)
	return ciphertext, nil
}

/**
  解密敏感信息
*/
func (this *ApiV3) DecryptText(ciphertext string) (string, error) {
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader

	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, this.merchantPrivatekey, cipherdata, nil)
	if err != nil {
		return "", fmt.Errorf("Error from decryption: %s\n", err)
	}

	return string(plaintext), nil

}

//获取平台秘钥，并保存
func (this *ApiV3) DownloadFlatPublicKey(filename string) error {
	data, err := this.DoGet("https://api.mch.weixin.qq.com/v3/certificates")
	if err != nil {
		return err
	}
	flatresponse := flatKeyResponse{}
	json.Unmarshal(data, &flatresponse)
	//保存文件

	return nil

}

//平台key返回结果对象
type flatKeyResponse struct {
	Data []certificateData
}
type certificateData struct {
	SerialNo           string             `json:"serialNo"`
	EffectiveTime      string             `json:"effectiveTime"`
	ExpireTime         string             `json:"expireTime"`
	EncryptCertificate encryptCertificate `json:"encrypt_certificate"`
}

type encryptCertificate struct {
	Algorithm      string `json:"algorithm"`
	Nonce          string `json:"nonce"`
	AssociatedData string `json:"associated_data"`
	Ciphertext     string `json:"ciphertext"`
}
