package easemob

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/guonaihong/gout"
)

type user struct {
	Base *Base
}

// Deactivate 用户账号禁用
func (u *user) Deactivate(userId string) (Response, error) {
	s := ""
	var code int
	err := gout.POST(fmt.Sprintf("%s/users/%s/deactivate", u.Base.Config.URL, userId)).
		SetHeader(u.Base.GetHeader()).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return u.response(code, s)
}

// Activate 用户账号解禁
func (u *user) Activate(userId string) (Response, error) {
	s := ""
	var code int
	err := gout.POST(fmt.Sprintf("%s/users/%s/activate", u.Base.Config.URL, userId)).
		SetHeader(u.Base.GetHeader()).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return u.response(code, s)
}

// Disconnect 强制下线
func (u *user) Disconnect(userId string) (Response, error) {
	s := ""
	var code int
	err := gout.GET(fmt.Sprintf("%s/users/%s/disconnect", u.Base.Config.URL, userId)).
		SetHeader(u.Base.GetHeader()).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return u.response(code, s)
}

func (u *user) response(code int, response string) (Response, error) {
	resp := Response{}
	if code != http.StatusOK {
		return resp, NewEasemobError(code, response)
	}
	if err := json.Unmarshal([]byte(response), &resp); err != nil {
		return resp, NewEasemobError(code, response)
	}
	return resp, nil
}
