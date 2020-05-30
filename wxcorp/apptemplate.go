package wxcorp

import "fmt"

type NavigationQuery struct {
	App     string `Field:"app"`
	CorpId  string `Field:"corp"`  //授权企业的企业号
	AgentId int    `Field:"agent"` //授权企业的应用agent
}

type LoginParamter struct {
	App     string `Field:"app"`
	CorpId  string `Field:"corp"`  //授权企业的企业号
	AgentId int    `Field:"agent"` //授权企业的应用agent
	Code    string `Field:"code"`
	State   string `Field:"state"`
}

/**
  应用模板
*/
type ApplicationTemplate struct {
	WxLoginUrl string
	Hosts      string
	AppHosts   string
	modules    map[string]string
	suit       WxCorpSuite
	stage      CorpDataStage
}

func (this *ApplicationTemplate) Init(option *WxCorpConfig, stage CorpDataStage) {
	this.modules = make(map[string]string)
	this.suit.Init(option, stage)

}
func (this *ApplicationTemplate) SetApplicationHandle(app CorpApplicationHandle, contact CorpOrgChangeHandle) {
	this.suit.SetHandle(app, contact)
}

func (this *ApplicationTemplate) SetReplyHandle(handle MessageReplyHandle) {
	this.suit.SetReplyHandle(handle)
}

func (this *ApplicationTemplate) AddModule(m string, url string) {
	this.modules[m] = url
}

//用户登录，返回企业id和用户信息
func (this *ApplicationTemplate) OnLogin(l *LoginParamter) (string, WxUserDetail) {
	//获取用户信息
	userDetail := this.suit.GetAuthCorpApi().WxOauthGetUser(l.App, l.CorpId, l.AgentId, WxRedirectParamter{l.Code, l.State})
	fmt.Println(userDetail)
	go this.suit.GetAuthCorpApi().WxInitContact(l.CorpId)
	return l.CorpId, userDetail

}

//导航，如果未登录返回登录跳转url，否则返回导航目标地址
func (this *ApplicationTemplate) OnNavigation(n *NavigationQuery, logined bool, state string) string {
	if logined {
		return fmt.Sprintf("%s%s", this.AppHosts, this.GetModuleUrl(n.App))
	} else {
		targeturl := fmt.Sprintf(this.Hosts+"/%s?app=%s&corp=%s&agent=%d", this.WxLoginUrl, n.App, n.CorpId, n.AgentId)
		return BuildRedirectUrl(n.CorpId, n.AgentId, targeturl, OAUTH_LEVEL_INFO, state)
	}
}

//管理员登录，如果是管理员返回管理员信息，否则返回nil
func (this *ApplicationTemplate) OnAdminLogin(code string) *LoginInfo {
	info := this.suit.GetLoginInfo(code)
	if info != nil && info.ErrCode == 0 {
		//检查信息
		if info.UserType != 5 { //不是普通成员
			return info
		}
	}
	return nil
}

func (this *ApplicationTemplate) GetModuleUrl(m string) string {
	return this.modules[m]
}
