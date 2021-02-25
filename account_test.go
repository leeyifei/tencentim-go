package tencentim

import (
	"fmt"
	"testing"
)

func TestAccount_Kick(t *testing.T) {
	req := &KickReq{
		Identifier: "user_1",
	}
	resp, err := s.Account.Kick(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Resp)
	fmt.Println(resp.ActionStatus)
	fmt.Println(resp.ErrorCode)
}
