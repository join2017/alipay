package alipay

type ZmAuthParams struct {
	BuckleAppId      string `json:"buckle_app_id"`      // 商户在芝麻端申请的appId
	BuckleMerchantId string `json:"buckle_merchant_id"` // 商户在芝麻端申请的 merchantId
}

type ProdParams struct {
	AuthBizParams string `json:"auth_biz_params"` // 预授权业务信息
}

type AccessParams struct {
	Channel string `json:"channel"` // 目前支持以下值：1.ALIPAYAPP（钱包h5页面签约）2.QRCODE(扫码签约)3.QRCODEORSMS(扫码签约或者短信签约)
}

type SubMerchantParams struct {
	SubMerchantId                 string `json:"sub_merchant_id"`
	SubMerchantName               string `json:"sub_merchant_name"`
	SubMerchantServiceName        string `json:"sub_merchant_service_name"`
	SubMerchantServiceDescription string `json:"sub_merchant_service_description"`
}

type DeviceParams struct {
	DeviceId   string `json:"device_id"`
	DeviceName string `json:"device_name"`
	DeviceType string `json:"device_type"` // 设备类型，目前有四种值：VR一体机：VR_MACHINE、电视：TV、身份证：ID_CARD、工牌：WORK_CARD
}

type IdentityParams struct {
	UserName     string `json:"user_name"`
	CertNo       string `json:"cert_no"`
	IdentityHash string `json:"identity_hash"`
	SignUserId   string `json:"sign_user_id"`
}

type PeriodRuleParams struct {
	PeriodType    string `json:"period_type"`
	Period        string `json:"period"`
	ExecuteTime   string `json:"execute_time"`
	SingleAmount  string `json:"single_amount"`
	TotalAmount   string `json:"total_amount"`
	TotalPayments int    `json:"total_payments"`
}

// --------------------------------------------------------------------------------
// https://docs.open.alipay.com/api_2/alipay.user.agreement.page.sign
// 支付宝个人协议页面签约接口
type AgreementPageSign struct {
	AppAuthToken        string            `json:"-"`
	ReturnURL           string            `json:"-"`
	NotifyURL           string            `json:"-"`
	SignValidityPeriod  string            `json:"sign_validity_period"`  // 当前用户签约请求的协议有效周期。整形数字加上时间单位的协议有效期，从发起签约请求的时间开始算起。目前支持的时间单位：1.d：天2.m：月如果未传入，默认为长期有效
	ProductCode         string            `json:"product_code"`          // 销售产品码，商户签约的支付宝合同所对应的产品码
	ExternalLogonId     string            `json:"external_logon_id"`     // 用户在商户网站的登录账号，用于在签约页面展示，如果为空，则不展示
	PersonalProductCode string            `json:"personal_product_code"` // 个人签约产品码，商户和支付宝签约时确定，商户可咨询技 术支持
	SignScene           string            `json:"sign_scene"`            // 协议签约场景，商户和支付宝 签约时确定，商户可咨询技术支持。当传入商户签约号external_agreement_no时，场景不能为默认值DEFAULT|DEFAULT
	ExternalAgreementNo string            `json:"external_agreement_no"` // 商户签约号，代扣协议中标示 用户的唯一签约号（确保在商 户系统中唯一）。
	ThirdPartyType      string            `json:"third_party_type"`      // 签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。1.PARTNER（平台商户）;2.MERCHANT（集团商户）;默认为PARTNER
	ZmAuthParams        ZmAuthParams      `json:"zm_auth_params"`        // 芝麻授权信息，针对于信用代扣签约。json格式。
	ProdParams          ProdParams        `json:"prod_params"`           // 签约产品属性
	PromoParams         string            `json:"promo_params"`          // 签约营销参数，此值为json格式；具体的key需与营销约定
	AccessParams        AccessParams      `json:"access_params"`
	SubMerchantParams   SubMerchantParams `json:"sub_merchant_params"`
	DeviceParams        DeviceParams      `json:"device_params"`
	MerchantProcessUrl  string            `json:"merchant_process_url"`
	IdentityParams      IdentityParams    `json:"identity_params"`
	AgreementEffectType string            `json:"agreement_effect_type"`
	UserAgeRange        string            `json:"user_age_range"`
	PeriodRuleParams    PeriodRuleParams  `json:"period_rule_params"`
}

type AgreementPageSignRsp struct {
	Content struct {
		Code                string `json:"code"`
		Msg                 string `json:"msg"`
		SubCode             string `json:"sub_code"`
		SubMsg              string `json:"sub_msg"`
		ExternalAgreementNo string `json:"external_agreement_no"` // 代扣协议中标示用户的唯一签约号(确保在商户系统中唯一)
		PersonalProductCode string `json:"personal_product_code"` // 协议产品码，商户和支付宝签约时确定，不同业务场景对应不同的签约产品码
		ValidTime           string `json:"valid_time"`            // 协议生效时间，格式为 yyyyMM-dd HH:mm:ss
		SignScene           string `json:"sign_scene"`            // 签约协议的场景
		AgreementNo         string `json:"agreement_no"`          // 用户签约成功后的协议号
		ZmOpenId            string `json:"zm_open_id"`            // 用户的芝麻信用 openId，供商户查询用户芝麻信用使用。
		InvalidTime         string `json:"invalid_time"`          // 协议失效时间，格式为 yyyyMM-dd HH:mm:ss
		SignTime            string `json:"sign_time"`             // 协议签约时间，格式为 yyyyMM-dd HH:mm:ss
		AlipayUserId        string `json:"alipay_user_id"`        // 用户的支付宝账号对应的支付宝唯一用户号，以2088开头的16位纯数字组成;本参数与alipay_logon_id不可同时为空，若都填写，则以本参数为准，优先级高于alipay_logon_id
		Status              string `json:"status"`                // 协议当前状态 1.TEMP：暂存，协议未生效过；2.NORMAL：正常；3.STOP：暂停
		ForexEligible       string `json:"forex_eligible"`        // 是否海外购汇身份。值：T/F（只有在签约成功时才会返回）
		ExternalLogonId     string `json:"external_logon_id"`     // 外部登录Id
		AlipayLogonId       string `json:"alipay_logon_id"`       // 返回脱敏的支付宝账号
	} `json:"agreement_page_sign_response"`
	Sign string `json:"sign"`
}

func (this AgreementPageSign) APIName() string {
	return "alipay.user.agreement.page.sign"
}

func (this AgreementPageSign) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	m["return_url"] = this.ReturnURL
	return m
}

// --------------------------------------------------------------------------------
// https://docs.open.alipay.com/api_2/alipay.user.agreement.query
// 支付宝个人代扣协议查询接口
type AgreementQuery struct {
	PersonalProductCode string `json:"personal_product_code"`  // 协议产品码，商户和支付宝签约时确定，商户可咨询技术支持
	AlipayUserId        string `json:"alipay_user_id"`         // 用户的支付宝账号对应的支付宝唯一用户号，以2088开头的16位纯数字组成;本参数与alipay_logon_id不可同时为空，若都填写，则以本参数为准，优先级高于alipay_logon_id
	AlipayLogonId       string `json:"alipay_logon_id"`        // 用户的支付宝登录账号，支持邮箱或手机号码格式。本参数与alipay_user_id不可同时为空，若都填写，则以alipay_user_id为准
	SignScene           string `json:"sign_scene"`             // 签约协议场景，商户和支付宝签约时确定，商户可咨询技术支持
	ExternalAgreementNo string `jsson:"external_agreement_no"` // 代扣协议中标示用户的唯一签约号(确保在商户系统中 唯一)。
	ThirdPartyType      string `jsson:"third_party_type"`      // 签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约
	AgreementNo         string `json:"agreement_no"`           // 支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号），如果传了该参数，其他参数会被忽略
	AppAuthToken        string `json:"-"`
}

type AgreementQueryRsp struct {
	Content struct {
		Code                string `json:"code"`
		Msg                 string `json:"msg"`
		SubCode             string `json:"sub_code"`
		SubMsg              string `json:"sub_msg"`
		ValidTime           string `json:"valid_time"`            // 协议生效时间，格式为 yyyyMM-dd HH:mm:ss
		AlipayLogonId       string `json:"alipay_logon_id"`       // 返回脱敏的支付宝账号
		InvalidTime         string `json:"invalid_time"`          // 协议失效时间，格式为 yyyyMM-dd HH:mm:ss
		PricipalType        string `json:"pricipal_type"`         // 签约主体类型。 CARD:支付宝账号 CUSTOMER:支付宝用户
		DeviceId            string `json:"device_id"`             // 设备Id
		PrincipalId         string `json:"principal_id"`          // 签约主体标识。当principal_type为CARD时，该字段为支付宝用户号;当principal_type为CUSTOMER时，该字段为支付宝用户标识。
		SignScene           string `json:"sign_scene"`            // 签约协议的场景
		AgreementNo         string `json:"agreement_no"`          // 用户签约成功后的协议号
		ThirdPartyType      string `json:"third_party_type"`      // 签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。 1.PARTNER（平台商户）;2.MERCHANT（集团商户），集团下子商户可共享用户签约内容;默认为PARTNER
		Status              string `json:"status"`                // 协议当前状态 1.TEMP：暂存，协议未生效过；2.NORMAL：正常；3.STOP：暂停
		SignTime            string `json:"sign_time"`             // 协议签约时间，格式为 yyyyMM-dd HH:mm:ss
		PersonalProductCode string `json:"personal_product_code"` // 协议产品码，商户和支付宝签约时确定，不同业务场景对应不同的签约产品码
		ExternalAgreementNo string `json:"external_agreement_no"` // 代扣协议中标示用户的唯一签约号(确保在商户系统中唯一)
		ZmOpenId            string `json:"zm_open_id"`            // 用户的芝麻信用 openId，供商户查询用户芝麻信用使用。
		ExternalLogonId     string `json:"external_logon_id"`     // 外部登录Id
	} `json:"alipay_user_agreement_query_response"`
	Sign string `json:"sign"`
}

func (this AgreementQuery) APIName() string {
	return "alipay.user.agreement.query"
}

func (this AgreementQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

// --------------------------------------------------------------------------------
// https://docs.open.alipay.com/api_2/alipay.user.agreement.unsign
// 支付宝个人代扣协议解约接口
type AgreementUnsign struct {
	AppAuthToken        string `json:"-"`
	NotifyURL           string `json:"-"`
	AlipayUserId        string `json:"alipay_user_id"`        // 用户的支付宝账号对应的支付宝唯一用户号，以2088开头的16位纯数字组成;本参数与alipay_logon_id不可同时为空，若都填写，则以本参数为准，优先级高于alipay_logon_id
	AlipayLogonId       string `json:"alipay_logon_id"`       // 返回脱敏的支付宝账号
	PersonalProductCode string `json:"personal_product_code"` // 协议产品码，商户和支付宝签约时确定，不同业务场景对应不同的签约产品码
	SignScene           string `json:"sign_scene"`            // 签约协议的场景
	ExternalAgreementNo string `json:"external_agreement_no"` // 代扣协议中标示用户的唯一签约号(确保在商户系统中唯一)
	ThirdPartyType      string `jsson:"third_party_type"`     // 签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约
	AgreementNo         string `json:"agreement_no"`          // 支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号），如果传了该参数，其他参数会被忽略
	ExtendParams        string `json:"extend_params"`
	OperateType         string `json:"operate_type"`
}
type AgreementUnsignRsp struct {
	Content struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"agreement_unsign_response"`
	Sign string `json:"sign"`
}

func (this AgreementUnsign) APIName() string {
	return "alipay.user.agreement.unsign"
}

func (this AgreementUnsign) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}
