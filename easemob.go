package easemob

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type H map[string]interface{}

var header = map[string]string{
	"Content-Type": "application/json",
	"Accept":       "application/json",
}

type Easemob struct {
	*Base
	Message *message
	Group   *group
	User    *user
}

func NewEasemob(config Config) *Easemob {
	base := Base{config: config, client: resty.New()}
	easemob := Easemob{
		Base:    &base,
		Message: &message{Base: &base},
		Group:   &group{Base: &base},
		User:    &user{Base: &base},
	}
	return &easemob
}

func (e *Easemob) SetDrive(drive cacheDrive) {
	e.Base.drive = drive
}

type Base struct {
	config Config
	client *resty.Client
	drive  cacheDrive
}

func (b *Base) GetToken() string {
	if b.drive.isValid() {
		return b.drive.getToken()
	}
	// 重新获取token
	b.drive.refresh(b.getToken())
	return b.drive.getToken()
}

func (b *Base) getToken() token {
	body := H{
		"grant_type":    "client_credentials",
		"client_id":     b.config.ClientId,
		"client_secret": b.config.ClientSecret,
	}
	resp, _ := b.client.R().
		SetHeaders(header).
		SetBody(body).
		Post(b.config.URL + "/token")
	m := token{}
	_ = json.Unmarshal(resp.Body(), &m)
	return m
}

func (b *Base) request(url string, method string, body H) (Response, error) {
	client := b.client.R().SetHeaders(header).SetAuthToken(b.GetToken())
	if len(body) > 0 {
		client.SetBody(body)
	}
	var resp *resty.Response
	var err error
	switch method {
	case resty.MethodGet:
		resp, err = client.Get(url)
	case resty.MethodPost:
		resp, err = client.Post(url)
	case resty.MethodPut:
		resp, err = client.Put(url)
	case resty.MethodDelete:
		resp, err = client.Delete(url)
	case resty.MethodOptions:
		resp, err = client.Options(url)
	case resty.MethodHead:
		resp, err = client.Head(url)
	case resty.MethodPatch:
		resp, err = client.Patch(url)
	default:
		resp, err = client.Get(url)
	}
	if err != nil {
		return Response{}, err
	}
	return b.response(resp)
}

func (b *Base) response(resp *resty.Response) (Response, error) {
	if resp.StatusCode() != http.StatusOK {
		return Response{}, NewEasemobError(resp.StatusCode(), resp.Body())
	}
	response := Response{}
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return response, errors.New("解析失败")
	}
	return response, nil
}
