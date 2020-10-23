package tencentim

import (
	"testing"
	"fmt"
)

func TestCreateGroup(test *testing.T) {
	req := &CreateGroupReq{
		Owner: "liyifei",
		Type: GROUP_TYPE_PRIVATE,
		Name: "test api",

	}

	resp, err := s.Group.CreateGroup(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Resp)
	fmt.Println(resp.ActionStatus)
	fmt.Println(resp.ErrorCode)
}

func TestModifyGroupBaseInfo(test *testing.T) {
	req := &ModifyGroupBaseInfoReq{
		GroupId: "@TGS#12F33JYGF",
		ShutUpAllMember: "On",
	}

	resp, err := s.Group.ModifyGroupBaseInfo(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
