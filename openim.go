package tencentim

const (
	MsgTypeText     = "TIMTextElem"
	MsgTypeLocation = "TIMLocationElem"
	MsgTypeFace     = "TIMFaceElem"
	MsgTypeCustom   = "TIMCustomElem"
	//MsgTypeSound     = "TIMSoundElem"
	//MsgTypeImage     = "TIMImageElem"
	//MsgTypeFIle      = "TIMFileElem"
	//MsgTypeVideoFile = "TIMVideoFileElem"

	OPENIM_SERVICENAME = "openim"
)

type OfflinePushInfo struct {
	PushFlag    uint        `json:"PushFlag,omitempty"`
	Title       string      `json:"Title,omitempty"`
	Desc        string      `json:"Desc,omitempty"`
	Ext         interface{} `json:"Ext,omitempty"`
	AndroidInfo struct {
		Sound           string `json:"Sound,omitempty"`
		HuaWeiChannelID string `json:"HuaWeiChannelID,omitempty"`
		XiaoMiChannelID string `json:"XiaoMiChannelID,omitempty"`
		OPPOChannelID   string `json:"OPPOChannelID,omitempty"`
		GoogleChannelID string `json:"GoogleChannelID,omitempty"`
	} `json:"AndroidInfo,omitempty"`
	ApnsInfo struct {
		BadgeMode uint   `json:"BadgeMode,omitempty"`
		Title     string `json:"Title,omitempty"`
		SubTitle  string `json:"SubTitle,omitempty"`
		Image     string `json:"Image"`
	} `json:"ApnsInfo,omitempty"`
}

type TypeText struct {
}

type SendMsgReq struct {
	SyncOtherMachine      uint     `json:"SyncOtherMachine,omitempty"`
	FromAccount           string   `json:"From_Account,omitempty"`
	ToAccount             string   `json:"To_Account"`
	MsgLifeTime           uint32   `json:"MsgLifeTime,omitempty"`
	MsgRandom             uint64   `json:"MsgRandom"`
	MsgTimeStamp          uint64   `json:"MsgTimeStamp,omitempty"`
	ForbidCallbackControl []string `json:"ForbidCallbackControl,omitempty"`
	MsgBody               []struct {
		MsgType    string      `json:"MsgType"`
		MsgContent interface{} `json:"MsgContent"`
	} `json:"MsgBody"`
	OfflinePushInfo OfflinePushInfo `json:"OfflinePushInfo,omitempty"`
}

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

// 查询用户账号状态
func (o *openim) Querystate(req *QuerystateReq) (*QuerystateResp, error) {
	var resp QuerystateResp

	err := o.Sdk.request(OPENIM_SERVICENAME, "querystate", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (o *openim) Sendmsg() {

}
