package easemob

import (
	"encoding/json"

	"github.com/guonaihong/gout"
)

type H map[string]interface{}

type Easemob struct {
	Base
	Message *message
	Group   *group
	User    *user
}

func NewEasemob(config Config) *Easemob {
	base := Base{Config: config, Token: &Token{}}
	easemob := Easemob{
		Base:    base,
		Message: &message{Base: &base},
		Group:   &group{Base: &base},
		User:    &user{Base: &base},
	}
	return &easemob
}

type Base struct {
	Config Config
	Token  *Token
}

func (b *Base) GetToken() string {
	if b.Token.isValid() {
		return b.Token.AccessToken
	}
	// 重新获取token
	b.Token.Refresh(b.getToken())
	return b.Token.AccessToken
}

func (b *Base) getToken() Token {
	s := ""
	_ = gout.POST(b.Config.URL + "/token").
		SetJSON(gout.H{
			"grant_type":    "client_credentials",
			"client_id":     b.Config.ClientId,
			"client_secret": b.Config.ClientSecret,
		}).
		SetHeader(gout.H{"Content-Type": "application/json"}).
		BindBody(&s).
		Do()
	m := Token{}
	_ = json.Unmarshal([]byte(s), &m)
	return m
}

func (b *Base) GetHeader() H {
	return H{
		"Content-Type":  "applicaiton/json",
		"Authorization": "Bearer " + b.GetToken(),
	}
}
