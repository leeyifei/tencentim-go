package tencentim

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type customMsgData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Link    string `json:"link"`
	Mt      string `json:"mt"`
	Mtc     string `json:"mtc"`
	E1      string `json:"e1"`
	E2      string `json:"e2"`
	E3      string `json:"e3"`
	E4      string `json:"e4"`
}

func TestOpenim_Querystate(t *testing.T) {
	req := &QuerystateReq{
		IsNeedDetail: 1,
		ToAccount:    []string{"user_1", "fdsf"},
	}

	resp, err := s.Openim.Querystate(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Resp)
	fmt.Println(resp.ActionStatus)
	fmt.Println("ErrorCode ===> ", resp.ErrorCode)
	fmt.Println(resp.QueryResult)
	fmt.Println(resp.ErrorList)
}

func TestOpenim_Sendmsg(t *testing.T) {
	req := &SendMsgReq{
		ToAccount: "user_1",
		MsgRandom: uint64(time.Now().Unix()),
		MsgBody:   make([]MsgBody, 0),
	}
	msg := MsgBody{
		MsgType: MsgTypeText,
		MsgContent: TypeText{
			Text: "hello world",
		},
	}
	req.MsgBody = append(req.MsgBody, msg)

	cm := customMsgData{
		Title:   "test",
		Content: "test",
	}
	cmb, _ := json.Marshal(&cm)
	msg = MsgBody{
		MsgType: MsgTypeCustom,
		MsgContent: TypeCustom{
			Data: string(cmb),
			Desc: "test custom message",
		},
	}
	req.MsgBody = append(req.MsgBody, msg)
	fmt.Println(req.MsgBody)
	resp, err := s.Openim.Sendmsg(req)
	fmt.Println(resp)
	fmt.Println(err)
}
