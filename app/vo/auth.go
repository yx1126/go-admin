package vo

type LoginVo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
	Uuid     string `json:"uuid"`
}

type RegisterVo struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	ConfirmPwd string `json:"confirmPwd"`
	Code       string `json:"code"`
	Uuid       string `json:"uuid"`
}
