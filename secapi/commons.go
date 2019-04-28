package secapi

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"sort"
	"strings"
)

type Request interface {
	Init(appid, merchant string)
	SetSign(s string)
	ToMap() map[string]string
}

type baseRequest struct {
	AppId    string `xml:"appid"`  //银行服务商的公众账号 ID
	Merchant string `xml:"mch_id"` //	银行服务商的商户号
	Sign     string `xml:"sign"`   //签名
}

func (this *baseRequest) Init(appid, merchant string) {
	this.AppId = appid
	this.Merchant = merchant
}

func (this *baseRequest) SetSign(s string) {
	this.Sign = s
}

type Client struct {
	AppId    string
	Merchant string
	Key      string
	CerFile  string
	KeyFile  string
}

//呼叫API
func (this *Client) CallApi(url string, request Request, response interface{}) error {
	//
	request.Init(this.AppId, this.Merchant)
	p := request.ToMap()
	s := this.sign(p)
	request.SetSign(s)
	req, _ := xml.Marshal(request)
	fmt.Println(string(req))
	body, err := this.postHttps(url, req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		fmt.Println(string(body))
		xml.Unmarshal(body, response)
		fmt.Printf("%s", response)
	}

	return nil
}

func (this *Client) sign(m map[string]string) string {
	var array []string
	for key, _ := range m {
		array = append(array, key)
	}
	sort.Strings(array)

	var parames []string
	for _, v := range array {
		value := m[v]
		if value != "" {
			parames = append(parames, fmt.Sprintf("%s=%s", v, value))
		}

	}

	return this.buildSign(parames...)
}

func (this *Client) buildSign(parames ...string) string {
	s := strings.Join(parames, "&") + "&key=" + this.Key

	fmt.Println(s)

	h := md5.New()
	h.Write([]byte(s)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	str := fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果
	return strings.ToUpper(str)
}

func (this *Client) postHttps(url string, data []byte) ([]byte, error) {
	cer, err := tls.LoadX509KeyPair(this.CerFile, this.KeyFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	//跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: config,
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	resp, err := c.Post(url, "application/xml;charset=utf-8", strings.NewReader(string(data)))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil
}
