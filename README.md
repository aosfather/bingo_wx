# bingo_wx
微信公众号、企业微信快速开发框架，无需熟悉微信的接口规范就可以快速对接现有的业务。基于bingo  
包括
* 微信公众号
* 企业微信
* 新增了 apiv3

## api v3 的使用
### 导入包
```go
import (
 "github.com/aosfather/bingo_wx/apiv3"
)

```
### 构建api v3 sdk
apiv3 需要服务商商户私钥加密信息和微信平台沟通
```go
api:=apiv3.ApiV3{}
api.MerchantId="服务商商户号"
//读取私钥
pk,err:=ioutil.ReadFile("e:/opt/私钥.pem")
	if err!=nil {
		t.Error("私钥没找到")
	}
//读取平台公钥如果有的话
fpub,err:=ioutil.ReadFile("e:/opt/平台公钥.cert")
	if err!=nil {
		t.Error("没有公钥")
	}
//设置私钥编号和私钥内容，编号在下载的界面有
api.SetMerchantKey("EDFA23343",pk))
//设置私钥编号和私钥内容，编号在下载的界面有
api.SetFlatKey("PEDFA23343",fpub))

```
### 开始使用sdk 调用相应的接口
现在apiv3包中有两个通用的方法对应Get和Post请求，DoGet，DoPost。  
除此之外apiv3包还封装了和实名认证申请相关的接口。例如下面的“查询商户实名状态的”
```go
//调用查询商户实名状态的接口
MerchantApplymentState(api,"112323232")
```
### sdk提供了下载平台公钥的接口

