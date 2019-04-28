package secapi

import (
	"encoding/xml"
)

const (
	BIND_APPID_URL = "https://api.mch.weixin.qq.com/secapi/mch/addsubdevconfig"
)

type configResponse struct {
	XMLName       xml.Name `xml:"xml"`
	Code          string   `xml:"return_code"`
	Message       string   `xml:"return_msg"`
	ResultCode    string   `xml:"result_code,omitempty"`
	ErrorCode     string   `xml:"err_code,omitempty"`
	ErrorDesption string   `xml:"err_code_des,omitempty"`
	Sign          string   `xml:"sign,omitempty"` //签名

}

/*
服务商给特约子商户配置绑定关系；
注意：
API只支持新增配置，不支持修改， 如银行需要修改请先登录微信服务商后台手工删除后重新配置。
可以绑定特约商户或渠道公司名字相同的公众号、小程序、开放平台应用的APPID; 如果提交绑定了subappid，支付接口就一定要传；

接口链接
URL地址: https://api.mch.weixin.qq.com/secapi/mch/addsubdevconfig
*/
type BindRequest struct {
	XMLName xml.Name `xml:"xml"`
	baseRequest
	SubMerchant string `xml:"sub_mch_id"` //银行服务商报备的特约商户识别码
	SubAppId    string `xml:"sub_appid"`  //绑定特约商户或渠道公众号、小程序、APP支付等对应的APPID
}

func (this *BindRequest) ToMap() map[string]string {
	r := make(map[string]string)
	r["appid"] = this.AppId
	r["mch_id"] = this.Merchant
	r["sub_mch_id"] = this.SubMerchant
	r["sub_appid"] = this.SubAppId
	return r

}

//绑定appid和子商户号
func (this *Client) BindSubMerchantApp(sub, app string) configResponse {
	bind := BindRequest{}
	bind.SubMerchant = sub
	bind.SubAppId = app
	r := configResponse{}
	this.CallApi(BIND_APPID_URL, &bind, &r)

	return r
}
