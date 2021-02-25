package tencentim

import (
	"fmt"
	"testing"
)

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
	fmt.Println(resp.ErrorCode)
	fmt.Println(resp.QueryResult)
	fmt.Println(resp.ErrorList)
}
