package easemob

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type user struct {
	Base *Base
}

// Deactivate 用户账号禁用
func (u *user) Deactivate(userId string) (Response, error) {
	url := fmt.Sprintf("%s/users/%s/deactivate", u.Base.config.URL, userId)
	return u.Base.request(url, resty.MethodPost, H{})
}

// Activate 用户账号解禁
func (u *user) Activate(userId string) (Response, error) {
	url := fmt.Sprintf("%s/users/%s/activate", u.Base.config.URL, userId)
	return u.Base.request(url, resty.MethodPost, H{})
}

// Disconnect 强制下线
func (u *user) Disconnect(userId string) (Response, error) {
	url := fmt.Sprintf("%s/users/%s/disconnect", u.Base.config.URL, userId)
	return u.Base.request(url, resty.MethodPost, H{})
}
