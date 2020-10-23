package tencentim

import (
	sig "github.com/tencentyun/tls-sig-api-golang"
	"time"
	"fmt"
	"math/rand"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
)

const (
	SIGN_MODE_KEYPAIR = 1

	API_DOMAIN = "console.tim.qq.com"
	API_VERSION = "v4"
)

type Resp struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int `json:"ErrorCode"`
}

type Sdk struct {
	sdkAppid   int
	identifier string
	publicKey  string
	privateKey string
	signMode   int
	// sig        string

	Group      *group
}

func NewTimByKeyPair(sdkAppid int, publicKey, privateKey, identifier string) (*Sdk, error) {
	// var err error

	t := &Sdk{
		sdkAppid: sdkAppid,
		identifier: identifier,
		publicKey: publicKey,
		privateKey: privateKey,
		signMode: SIGN_MODE_KEYPAIR,
	}

	// if t.sig, err = UserSig(t.privateKey, t.sdkAppid, t.identifier); err != nil {
	//	return nil, err
	// }

	t.Group = &group{
		Tim: t,
	}

	return t, nil
}

func (t *Sdk) request(servicename, command string, reqData interface{}, out interface{}) (error) {
	var err error

	reqBodyByte, err := json.Marshal(reqData)
	if err != nil {
		return err
	}
	reqBody := bytes.NewReader(reqBodyByte)

	url, err := t.combineUrl(servicename, command)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, reqBody)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body fail: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http erro status code is %d, response is: %s", resp.StatusCode, respBody)
	}

	json.Unmarshal(respBody, out)

	return err
}

func (t *Sdk) combineUrl(servicename, command string) (string, error) {
	rand.Seed(time.Now().Unix())
	userSig, err := UserSig(t.privateKey, t.sdkAppid, t.identifier)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s/%s/%s/%s?sdkappid=%d&identifier=%s&usersig=%s&random=%d&contenttype=json",
		API_DOMAIN, API_VERSION, servicename, command, t.sdkAppid, t.identifier, userSig, rand.Intn(4294967294)), nil
}

func UserSig(privateKey string, appid int, identifier string) (string, error) {
	return sig.GenerateUsersig(privateKey, appid, identifier)
}