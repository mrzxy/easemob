package easemob

import "time"

type token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Application string `json:"application"`
	LastTime    int64  `json:"last_time"`
}

// isValid 是否有效
func (t *token) isValid() bool {
	return t.ExpiresIn+t.LastTime-600 > time.Now().Unix()
}

// Refresh 刷新token
func (t *token) Refresh(token token) {
	t.AccessToken = token.AccessToken
	t.ExpiresIn = token.ExpiresIn
	t.Application = token.Application
	t.LastTime = time.Now().Unix()
}
