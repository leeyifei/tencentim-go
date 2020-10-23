package tencentim

import (
	"os"
	"fmt"
	"testing"
	"strconv"
)

var s *Sdk

func init() {
	var err error

	sdkAppid, err := strconv.Atoi(os.Getenv("TIM_SDKAPPID"))
	if err != nil {
		panic(err)
	}

	s, err = NewTimByKeyPair(sdkAppid, os.Getenv("TIM_PUBLICKEY"), os.Getenv("TIM_PRIVATEKEY"), os.Getenv("TIM_IDENTIFIER"))
	if err != nil {
		panic(err)
	}
}

func TestCombineUrl(test *testing.T) {
	url, _ := s.combineUrl("servicename", "command")
	fmt.Println(url)
}

func TestRequest(test *testing.T) {
	var out Resp
	err := s.request("test", "test", nil, &out)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}