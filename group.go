package tencentim

const (
	GROUP_TYPE_PRIVATE    = "Private"
	GROUP_TYPE_PUBLIC     = "Public"
	GROUP_TYPE_CHATROOM   = "ChatRoom"
	GROUP_TYPE_AVCHATROOM = "AVChatRoom"

	GROUP_SERVICENAME = "group_open_http_svc"
)

type CreateGroupReq struct {
	Owner           string `json:"Owner_Account"`
	Type            string `json:"Type"`
	GroupId         string `json:"GroupId,omitempty"`
	Name            string `json:"Name"`
	Introduction    string `json:"Introduction,omitempty"`
	Notification    string `json:"Notification,omitempty"`
	FaceuUrl        string `json:"FaceUrl,omitempty"`
	MaxMemberCount  int    `json:"MaxMemberCount,omitempty"`
	ApplyJoinOption string `json:"ApplyJoinOption,omitempty"`
}

type CreateGroupResp struct {
	Resp
	GroupId string `json:"GroupId"`
}

type group struct {
	Sdk *Sdk
}

type ModifyGroupBaseInfoReq struct {
	GroupId         string `json:"GroupId"`
	Name            string `json:"Name,omitempty"`
	Introduction    string `json:"Introduction,omitempty"`
	Notification    string `json:"Notification,omitempty"`
	FaceUrl         string `json:"FaceUrl,omitempty"`
	MaxMemberNum    string `json:"MaxMemberNum,omitempty"`
	ApplyJoinOption string `json:"ApplyJoinOption,omitempty"`
	ShutUpAllMember string `json:"ShutUpAllMember,omitempty"`
}

type ModifyGroupBaseInfoResp struct {
	Resp
}

func (g *group) CreateGroup(data *CreateGroupReq) (*CreateGroupResp, error) {
	var (
		resp CreateGroupResp
	)

	err := g.Sdk.request(GROUP_SERVICENAME, "create_group", data, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (g *group) ModifyGroupBaseInfo(data *ModifyGroupBaseInfoReq) (*ModifyGroupBaseInfoResp, error) {
	var resp ModifyGroupBaseInfoResp

	err := g.Sdk.request(GROUP_SERVICENAME, "modify_group_base_info", data, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
