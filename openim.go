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
	Text string `json:"Text"`
}
type TypeLocation struct {
	Desc      string  `json:"Desc"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}
type TypeFace struct {
	Index uint   `json:"Index"`
	Data  string `json:"Data"`
}
type TypeCustom struct {
	Data  string ` json:"Data"`
	Desc  string `json:"Desc"`
	Ext   string `json:"Ext"`
	Sound string `json:"Sound"`
}

type MsgBody struct {
	MsgType    string      `json:"MsgType"`
	MsgContent interface{} `json:"MsgContent"`
}

type SendMsgReq struct {
	SyncOtherMachine      uint            `json:"SyncOtherMachine,omitempty"`
	FromAccount           string          `json:"From_Account,omitempty"`
	ToAccount             string          `json:"To_Account"`
	MsgLifeTime           uint32          `json:"MsgLifeTime,omitempty"`
	MsgRandom             uint64          `json:"MsgRandom"`
	MsgTimeStamp          uint64          `json:"MsgTimeStamp,omitempty"`
	ForbidCallbackControl []string        `json:"ForbidCallbackControl,omitempty"`
	MsgBody               []MsgBody       `json:"MsgBody"`
	OfflinePushInfo       OfflinePushInfo `json:"OfflinePushInfo,omitempty"`
}

type SendMsgResp struct {
	Resp
	MsgTime uint64 `json:"MsgTime"`
	MsgKey  string `json:"MsgKey"`
}

type BatchSendMsgReq struct {
	SyncOtherMachine uint            `json:"SyncOtherMachine,omitempty"`
	FromAccount      string          `json:"From_Account,omitempty"`
	ToAccount        []string        `json:"To_Account"`
	MsgRandom        uint64          `json:"MsgRandom"`
	MsgBody          []MsgBody       `json:"MsgBody"`
	OfflinePushInfo  OfflinePushInfo `json:"OfflinePushInfo,omitempty"`
}

type BatchSendMsgErrorList struct {
	ToAccount string `json:"To_Account"`
	ErrorCode uint   `json:"ErrorCode"`
}

type BatchSendMsgResp struct {
	Resp
	MsgKey    string                  `json:"MsgKey"`
	ErrorList []BatchSendMsgErrorList `json:"ErrorList"`
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

// 发送单聊消息
func (o *openim) Sendmsg(req *SendMsgReq) (*SendMsgResp, error) {
	var resp SendMsgResp
	err := o.Sdk.request(OPENIM_SERVICENAME, "sendmsg", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, err
}

// 批量发送单聊消息
func (o *openim) BatchSendmsg(req *BatchSendMsgReq) (*BatchSendMsgResp, error) {
	var resp BatchSendMsgResp
	err := o.Sdk.request(OPENIM_SERVICENAME, "batchsendmsg", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, err
}
