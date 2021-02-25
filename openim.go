package tencentim

const (
	OPENIM_SERVICENAME = "openim"
)

type QuerystateReq struct {
	IsNeedDetail int      `json:"IsNeedDetail,omitempty"`
	ToAccount    []string `json:"To_Account"`
}

type QuerystateResp struct {
	Resp
	QueryResult []struct {
		ToAccount string `json:"To_Account"`
		Status    string `json:"Status"`
		Detail    []struct {
			Platform string `json:"Platform"`
			Status   string `json:"Status"`
		} `json:"Detail,omitempty"`
	} `json:"QueryResult"`
	ErrorList []struct {
		ToAccount string `json:"To_Account"`
		ErrorCode int    `json:"ErrorCode"`
	} `json:"ErrorList"`
}

type openim struct {
	Sdk *Sdk
}

func (o *openim) Querystate(req *QuerystateReq) (*QuerystateResp, error) {
	var resp QuerystateResp

	err := o.Sdk.request(OPENIM_SERVICENAME, "querystate", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
