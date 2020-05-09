package apiv3

import (
	"encoding/json"
	"fmt"
)

const (
	//已实名
	Authorize_authorized = "AUTHORIZE_STATE_AUTHORIZED"
	//未实名
	Authorize_unauthorized = "AUTHORIZE_STATE_UNAUTHORIZED"
)

type authorizeResponse struct {
	State string `json:"authorize_state"`
}

//获取特约商户的实名状况
func MerchantApplymentState(v3 ApiV3, merchant string) bool {
	url := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/apply4subject/applyment/merchants/%s/state", merchant)
	body, err := v3.DoGet(url)
	if err != nil {
		return false
	}
	response := authorizeResponse{}
	json.Unmarshal([]byte(body), &response)
	if response.State == Authorize_authorized {
		return true
	}
	return false
}
