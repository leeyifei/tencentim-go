package tencentim

const (
	ACCOUNT_SERVICENAME = "im_open_login_svc"
)

type KickReq struct {
	Identifier string `json:"Identifier"`
}

type KickResp struct {
	Resp
}

type account struct {
	Sdk *Sdk
}

// 失效帐号登录状态
func (a *account) Kick(req *KickReq) (*KickResp, error) {
	var resp KickResp

	err := a.Sdk.request(ACCOUNT_SERVICENAME, "kick", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
