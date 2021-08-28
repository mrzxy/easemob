package easemob

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const tokenKey = "hxToken"

type token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Application string `json:"application"`
	LastTime    int64  `json:"last_time"`
}

type cacheDrive interface {
	isValid() bool
	refresh(token2 token)
	getToken() string
}

type RedisDrive struct {
	Client *redis.Client
}

func (r *RedisDrive) isValid() bool {
	_, err := r.Client.Get(context.Background(), tokenKey).Result()
	return err == nil
}

func (r RedisDrive) refresh(token2 token) {
	r.Client.Set(context.Background(), tokenKey, token2.AccessToken, time.Duration(token2.ExpiresIn)*time.Second-600)
}

func (r *RedisDrive) getToken() string {
	t, err := r.Client.Get(context.Background(), tokenKey).Result()
	if err != nil {
		fmt.Println(err)
	}
	return t
}
