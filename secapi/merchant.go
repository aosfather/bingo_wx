package secapi

import (
	"encoding/xml"
	"fmt"
)

const (
	SUB_MERCHANT_MANAGE  = "https://api.mch.weixin.qq.com/secapi/mch/submchmanage?action=add"
	MODIFY_MERCHANT_INFO = "https://api.mch.weixin.qq.com/secapi/mch/modifymchinfo"
	QUERY_MERCHANT_INFO  = "https://api.mch.weixin.qq.com/secapi/mch/submchmanage?action=query"
)

type merchantResponse struct {
	XMLName       xml.Name `xml:"xml"`
	Code          string   `xml:"return_code"`
	Message       string   `xml:"return_msg"`
	Merchant      string   `xml:"mch_id,omitempty"`     //	银行服务商的商户号
	SubMerchant   string   `xml:"sub_mch_id,omitempty"` //银行服务商报备的特约商户识别码
	ResultCode    string   `xml:"result_code,omitempty"`
	ResultMessage string   `xml:"result_msg,omitempty"`
	ErrorCode     string   `xml:"err_code,omitempty"`
	ErrorMessage  string   `xml:"err_code_des,omitempty"`
}

/*
银行特约商户录入API
应用场景
银行服务商接入微信支付前需要将下属特约商户基本资料信息报备给微信，在微信支付侧生成特约商户识别码后方可提交微信支付。特约商户识别码是区分子商户交易、结算和清分的标志。

银行服务商需先在管理后台录入渠道，渠道审核通过后，才可以录入特约商户。银行特约商户数量较多，一般采调用该API实现快速录入（注意：普通商户和普通服务商不可调用该接口 )。该接口有频率1次/秒限制。

银行服务商调用该接口将特约商户资料提交微信侧，微信根据提交的资料情况，判断商户资料正确性，返回识别码；如商户资料有误，返回相应的错误码。

接口链接
URL地址：https://api.mch.weixin.qq.com/secapi/mch/submchmanage?action=add

行业品类ID	一级类目	二级类目	三级类目
545	企业	餐饮	快餐
543	企业	餐饮	火锅
544	企业	餐饮	烧烤
548	企业	餐饮	小吃/熟食
539	企业	餐饮	其他中餐
541	企业	餐饮	日韩/东南亚菜
540	企业	餐饮	西餐
542	企业	餐饮	咖啡厅
550	企业	餐饮	甜品饮品
549	企业	餐饮	烘焙糕点
204	企业	线下零售	超市
205	企业	线下零售	便利店
207	企业	线下零售	百货
206	企业	线下零售	自动贩卖机
270	企业	线下零售	食品生鲜
310	企业	线下零售	数码电器/电脑办公
266	企业	线下零售	家具建材/家居厨具
271	企业	线下零售	服饰箱包
6	企业	线下零售	运动户外
267	企业	线下零售	美妆个护
19	企业	线下零售	母婴用品/儿童玩具
13	企业	线下零售	计生用品
9	企业	线下零售	黄金珠宝
272	企业	线下零售	礼品鲜花/农资绿植
284	企业	线下零售	宠物/宠物用品
268	企业	线下零售	钟表眼镜
585	企业	线下零售	批发业
315	企业	线下零售	图书音像/文具乐器
203	企业	网上购物	线上商超
70	企业	居民生活/商业服务	物流/快递公司
311	企业	居民生活/商业服务	家政/维修服务
273	企业	居民生活/商业服务	婚庆/摄影
564	企业	居民生活/商业服务	丧仪殡葬服务
565	企业	居民生活/商业服务	搬家/回收
538	企业	居民生活/商业服务	共享服务
566	企业	居民生活/商业服务	宠物医院
317	企业	居民生活/商业服务	苗木种植/园林绿化
289	企业	居民生活/商业服务	装饰/设计
312	企业	居民生活/商业服务	广告/会展/活动策划
42	企业	居民生活/商业服务	咨询/法律咨询/金融咨询等
93	企业	居民生活/商业服务	人才中介机构/招聘/猎头
94	企业	居民生活/商业服务	职业社交/婚介/交友
316	企业	居民生活/商业服务	房地产
281	企业	休闲娱乐	娱乐票务
572	企业	休闲娱乐	院线影城
573	企业	休闲娱乐	演出赛事
54	企业	休闲娱乐	运动健身场馆
574	企业	休闲娱乐	美发/美容/美甲店
280	企业	休闲娱乐	俱乐部/休闲会所
56	企业	休闲娱乐	游艺厅/KTV/网吧
571	企业	休闲娱乐	酒吧
77	企业	交通出行/票务旅游	租车
274	企业	交通出行/票务旅游	机票/机票代理
283	企业	交通出行/票务旅游	铁路客运
610	企业	交通出行/票务旅游	高速收费
288	企业	交通出行/票务旅游	城市公共交通
259	企业	交通出行/票务旅游	加油
287	企业	交通出行/票务旅游	停车缴费
269	企业	交通出行/票务旅游	汽车用品
563	企业	交通出行/票务旅游	汽车美容/维修保养
75	企业	交通出行/票务旅游	船舶/海运服务
354	企业	交通出行/票务旅游	景区
275	企业	交通出行/票务旅游	旅馆/酒店/度假区
23	企业	交通出行/票务旅游	旅行社
276	企业	网络媒体/计算机服务/游戏	在线图书/视频/音乐
104	企业	网络媒体/计算机服务/游戏	门户/资讯/论坛
501	企业	网络媒体/计算机服务/游戏	游戏
521	企业	网络媒体/计算机服务/游戏	网络直播
277	企业	网络媒体/计算机服务/游戏	软件/建站/技术开发
278	企业	网络媒体/计算机服务/游戏	网络推广/网络广告
2	企业	网上服务平台	团购服务平台
95	企业	网上服务平台	综合生活服务平台
24	企业	网上服务平台	旅游服务平台
522	企业	网上服务平台	订餐服务平台
52	企业	教育/医疗	教育/培训/考试缴费/学费
53	企业	教育/医疗	私立院校
314	企业	教育/医疗	保健器械/医疗器械/非处方药品
282	企业	教育/医疗	保健信息咨询平台
66	企业	教育/医疗	私立/民营医院/诊所
67	企业	教育/医疗	挂号平台
80	企业	生活缴费	电信运营商
81	企业	生活缴费	宽带收费
92	企业	生活缴费	话费通讯
58	企业	生活缴费	有线电视缴费
60	企业	生活缴费	物业管理费
57	企业	生活缴费	水电煤缴费/交通罚款等生活缴费
62	企业	生活缴费	其他生活缴费
96	企业	金融	财经资讯
97	企业	金融	股票软件类
318	企业	金融	保险业务
112	企业	金融	众筹
326	企业	金融	信用还款
31	企业	收藏/拍卖	非文物类收藏品
285	企业	收藏/拍卖	文物经营/文物复制品销售
325	企业	收藏/拍卖	拍卖/典当
111	企业	其他	其他行业
557	个体工商户	餐饮	快餐
555	个体工商户	餐饮	火锅
556	个体工商户	餐饮	烧烤
560	个体工商户	餐饮	小吃/熟食
551	个体工商户	餐饮	其他中餐
553	个体工商户	餐饮	日韩/东南亚菜
552	个体工商户	餐饮	西餐
554	个体工商户	餐饮	咖啡厅
562	个体工商户	餐饮	甜品饮品
561	个体工商户	餐饮	烘焙糕点
209	个体工商户	线下零售	便利店
292	个体工商户	线下零售	食品生鲜
319	个体工商户	线下零售	数码电器/电脑办公
293	个体工商户	线下零售	家具建材/家居厨具
297	个体工商户	线下零售	服饰箱包
116	个体工商户	线下零售	运动户外
294	个体工商户	线下零售	美妆个护
129	个体工商户	线下零售	母婴用品/儿童玩具
123	个体工商户	线下零售	计生用品
323	个体工商户	线下零售	图书音像/文具乐器
298	个体工商户	线下零售	钟表眼镜
305	个体工商户	线下零售	宠物/宠物用品
295	个体工商户	线下零售	礼品鲜花/农资绿植
586	个体工商户	线下零售	批发业
299	个体工商户	居民生活/商业服务	婚庆/摄影
568	个体工商户	居民生活/商业服务	丧仪殡葬服务
569	个体工商户	居民生活/商业服务	搬家/回收
570	个体工商户	居民生活/商业服务	宠物医院
143	个体工商户	居民生活/商业服务	咨询/法律咨询/金融咨询等
306	个体工商户	居民生活/商业服务	装饰/设计
320	个体工商户	居民生活/商业服务	家政/维修服务
321	个体工商户	居民生活/商业服务	广告/会展/活动策划
157	个体工商户	居民生活/商业服务	职业社交/婚介/交友
324	个体工商户	居民生活/商业服务	苗木种植/园林绿化
307	个体工商户	休闲娱乐	娱乐票务
578	个体工商户	休闲娱乐	院线影城
579	个体工商户	休闲娱乐	演出赛事
148	个体工商户	休闲娱乐	运动健身场馆
580	个体工商户	休闲娱乐	美发/美容/美甲店
300	个体工商户	休闲娱乐	俱乐部/休闲会所
149	个体工商户	休闲娱乐	游艺厅/KTV/网吧
577	个体工商户	休闲娱乐	酒吧
296	个体工商户	交通运输服务	汽车用品
567	个体工商户	交通运输服务	汽车美容/维修保养
502	个体工商户	网络媒体/计算机服务/游戏	游戏
302	个体工商户	网络媒体/计算机服务/游戏	软件/建站/技术开发
303	个体工商户	网络媒体/计算机服务/游戏	网络推广/网络广告
147	个体工商户	教育/医疗	教育/培训/考试缴费/学费
322	个体工商户	教育/医疗	保健器械/医疗器械/非处方药品
230	个体工商户	教育/医疗	私立/民营医院/诊所
155	个体工商户	生活缴费	话费通讯
309	个体工商户	生活缴费	生活缴费
242	个体工商户	金融	财经资讯
301	个体工商户	交通出行/票务旅游	旅馆/酒店/度假区
308	个体工商户	交通出行/票务旅游	铁路客运
340	个体工商户	交通出行/票务旅游	加油
158	个体工商户	其他	其他行业
164	党政、机关及事业单位	教育/医疗	公立院校
176	党政、机关及事业单位	教育/医疗	公立医院
177	党政、机关及事业单位	教育/医疗	挂号平台
290	党政、机关及事业单位	交通运输服务	停车缴费
165	党政、机关及事业单位	其他生活缴费	水电煤缴费/交通罚款等生活缴费
167	党政、机关及事业单位	其他生活缴费	事业单位
170	党政、机关及事业单位	其他生活缴费	物业管理费
172	党政、机关及事业单位	其他生活缴费	其他生活缴费
506	其他组织	教育/医疗	教育/培训/考试缴费/学费
517	其他组织	教育/医疗	私立院校
534	其他组织	教育/医疗	诊所/卫生站/卫生服务中心
523	其他组织	生活/咨询服务	咨询/法律咨询/金融咨询等
509	其他组织	交通运输/票务旅游	宗教
535	其他组织	交通运输/票务旅游	娱乐票务
537	其他组织	交通运输/票务旅游	机票/机票代理
510	其他组织	公益	公益
*/

type AddRequest struct {
	XMLName xml.Name `xml:"xml"`
	baseRequest
	Name         string `xml:"merchant_name"`           //
	ShortName    string `xml:"merchant_shortname"`      //商户简称，该名称是显示给消费者看的商户名称
	ServicePhone string `xml:"service_phone,omitempty"` //方便微信在必要时能联系上商家，会在支付详情展示给消费者
	contactInfo

	Channel  string `xml:"channel_id"`
	Business string `xml:"business"`
	Remark   string `xml:"merchant_remark"`
}

type contactInfo struct {
	Contact      string `xml:"contact,omitempty"`               //联系人	方便微信在必要时能联系上商家。
	ContactPhone string `xml:"contact_phone,omitempty"`         //联系电话
	ContactMail  string `xml:"contact_email,omitempty"`         //联系邮箱
	WebchatType  string `xml:"contact_wechatid_type,omitempty"` //联系人微信账号类型  微信号，值为type_wechatid,要传openid,值为type_openid
	WebchatId    string `xml:"contact_wechatid,omitempty"`      //联系人微信帐号
}

func (this *AddRequest) ToMap() map[string]string {
	r := make(map[string]string)
	r["appid"] = this.AppId
	r["mch_id"] = this.Merchant
	r["merchant_name"] = this.Name
	r["merchant_shortname"] = this.ShortName
	r["service_phone"] = this.ServicePhone
	r["channel_id"] = this.Channel
	r["business"] = this.Business
	r["contact"] = this.Contact
	r["contact_phone"] = this.ContactPhone
	r["contact_email"] = this.ContactMail
	r["contact_wechatid_type"] = this.WebchatType
	r["contact_wechatid"] = this.WebchatId
	return r
}

func (this *Client) AddMerchant(req AddRequest) merchantResponse {
	r := merchantResponse{}
	this.CallApi(SUB_MERCHANT_MANAGE, &req, &r)
	return r
}

/**
银行特约商户信息修改API
应用场景
1.目前支持修改简称和客服电话，子商户修改资料成功后，需要隔30天后才能修改。
2.商户简称是用户支付界面显示的商户名，必须是能清楚标识某个商户，不能乱传。

接口链接
URL地址：https://api.mch.weixin.qq.com/secapi/mch/modifymchinfo


*/
type ModifyRequest struct {
	XMLName xml.Name `xml:"xml"`
	baseRequest
	SubMerchant  string `xml:"sub_mch_id"`              //银行服务商报备的特约商户识别码
	ShortName    string `xml:"merchant_shortname"`      //商户简称，该名称是显示给消费者看的商户名称
	ServicePhone string `xml:"service_phone,omitempty"` //方便微信在必要时能联系上商家，会在支付详情展示给消费者
}

func (this *ModifyRequest) ToMap() map[string]string {
	r := make(map[string]string)
	r["appid"] = this.AppId
	r["mch_id"] = this.Merchant
	r["sub_mch_id"] = this.SubMerchant
	r["merchant_shortname"] = this.ShortName
	r["service_phone"] = this.ServicePhone
	return r

}

func (this *Client) ModifyMerchant(sub, sname, sphone string) merchantResponse {
	modify := ModifyRequest{}
	modify.SubMerchant = sub
	modify.ShortName = sname
	modify.ServicePhone = sphone
	r := merchantResponse{}
	this.CallApi(MODIFY_MERCHANT_INFO, &modify, &r)

	return r
}

/*
 应用场景
提供给银行服务商报备后的商户查询。通过MCHID（识别码），返回商户全部资料信息。

接口链接
URL地址：https://api.mch.weixin.qq.com/secapi/mch/submchmanage?action=query
*/
type queryRequest struct {
	XMLName xml.Name `xml:"xml"`
	baseRequest
	SubMerchant  string `xml:"sub_mch_id,omitempty"`    //银行服务商报备的特约商户识别码
	MerchantName string `xml:"merchant_name,omitempty"` //商户简称，该名称是显示给消费者看的商户名称
	PageIndex    int    `xml:"page_index"`
	PageSize     int    `xml:"page_size,omitempty"`
}

func (this *queryRequest) ToMap() map[string]string {
	r := make(map[string]string)
	r["appid"] = this.AppId
	r["mch_id"] = this.Merchant
	r["sub_mch_id"] = this.SubMerchant
	r["merchant_name"] = this.MerchantName
	r["page_index"] = fmt.Sprintf("%d", this.PageIndex)
	r["page_size"] = fmt.Sprintf("%d", this.PageSize)
	return r

}

func (this *Client) QueryMerchant(code, name string) merchantResponse {
	q := queryRequest{}
	q.PageIndex = 1
	q.PageSize = 100
	if code != "" {
		q.SubMerchant = code
	} else {
		q.MerchantName = name
	}

	r := merchantResponse{}
	this.CallApi(QUERY_MERCHANT_INFO, &q, &r)

	return r
}
