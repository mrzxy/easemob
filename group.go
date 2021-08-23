package easemob

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/guonaihong/gout"
)

type group struct {
	Base *Base
}

// GetUsers 分页获取群组成员
func (g *group) GetUsers(groupId string) (Response, error) {
	s := ""
	var code int
	err := gout.GET(fmt.Sprintf("%s/chatgroups/%s/users", g.Base.Config.URL, groupId)).
		SetHeader(g.Base.GetHeader()).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return g.response(code, s)
}

// AddUser 添加单个群组成员
func (g *group) AddUser(groupId string, userId string) (Response, error) {
	s := ""
	var code int
	err := gout.POST(fmt.Sprintf("%s/chatgroups/%s/users/%s", g.Base.Config.URL, groupId, userId)).
		SetHeader(g.Base.GetHeader()).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return g.response(code, s)
}

// AddUsers 批量添加群组成员
func (g *group) AddUsers(groupId string, userIds []string) (Response, error) {
	s := ""
	var code int
	err := gout.POST(fmt.Sprintf("%s/chatgroups/%s/users", g.Base.Config.URL, groupId)).
		SetHeader(g.Base.GetHeader()).SetBody(gout.H{"usernames": userIds}).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return g.response(code, s)
}

// DeleteUser 移除单个群组成员
func (g *group) DeleteUser(groupId string, userId string) (Response, error) {
	s := ""
	var code int
	err := gout.DELETE(fmt.Sprintf("%s/chatgroups/%s/users/%s", g.Base.Config.URL, groupId, userId)).
		SetHeader(g.Base.GetHeader()).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return g.response(code, s)
}

// DeleteUsers 批量移除群组成员 len(userIds) <= 60
func (g *group) DeleteUsers(groupId string, userIds []string) (Response, error) {
	s := ""
	var code int
	err := gout.DELETE(fmt.Sprintf("%s/chatgroups/%s/users/%s", g.Base.Config.URL, groupId, strings.Join(userIds, ","))).
		SetHeader(g.Base.GetHeader()).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return g.response(code, s)
}

// GetAdmin 获取群管理员列表
func (g *group) GetAdmin(groupId string) (Response, error) {
	s := ""
	var code int
	err := gout.GET(fmt.Sprintf("%s/chatgroups/%s/admin", g.Base.Config.URL, groupId)).
		SetHeader(g.Base.GetHeader()).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return g.response(code, s)
}

// AddAdmin 添加群管理员
func (g *group) AddAdmin(groupId string, userId string) (Response, error) {
	s := ""
	var code int
	err := gout.POST(fmt.Sprintf("%s/chatgroups/%s/admin", g.Base.Config.URL, groupId)).
		SetHeader(g.Base.GetHeader()).SetBody(gout.H{"newadmin": userId}).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return g.response(code, s)
}

// DeleteAdmin 移除群管理员
func (g *group) DeleteAdmin(groupId string, userId string) (Response, error) {
	s := ""
	var code int
	err := gout.DELETE(fmt.Sprintf("%s/chatgroups/%s/admin/%s", g.Base.Config.URL, groupId, userId)).
		SetHeader(g.Base.GetHeader()).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return g.response(code, s)
}

// TransferGroup 转让群组
func (g *group) TransferGroup(groupId string, userId string) (Response, error) {
	s := ""
	var code int
	err := gout.PUT(fmt.Sprintf("%s/chatgroups/%s", g.Base.Config.URL, groupId)).
		SetHeader(g.Base.GetHeader()).SetBody(gout.H{"newowner": userId}).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return g.response(code, s)
}

// AddMute 添加禁言
// duration 禁言的时间，单位毫秒，如果是“-1”代表永久
func (g *group) AddMute(groupId string, duration int, userIds []string) (Response, error) {
	s := ""
	var code int
	err := gout.POST(fmt.Sprintf("%s/chatgroups/%s/mute", g.Base.Config.URL, groupId)).
		SetHeader(g.Base.GetHeader()).SetBody(gout.H{"mute_duration": duration, "usernames": userIds}).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return g.response(code, s)
}

// RemoveMute 移除禁言
func (g *group) RemoveMute(groupId string, userIds []string) (Response, error) {
	s := ""
	var code int
	err := gout.POST(fmt.Sprintf("%s/chatgroups/%s/mute/%s", g.Base.Config.URL, groupId, strings.Join(userIds, ","))).
		SetHeader(g.Base.GetHeader()).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return g.response(code, s)
}

// GetMuteUsers 获取禁言列表
func (g *group) GetMuteUsers(groupId string) (Response, error) {
	s := ""
	var code int
	err := gout.GET(fmt.Sprintf("%s/chatgroups/%s/mute", g.Base.Config.URL, groupId)).
		SetHeader(g.Base.GetHeader()).
		BindBody(&s).Code(&code).Do()
	if err != nil {
		return Response{}, err
	}
	return g.response(code, s)
}

func (g *group) response(code int, response string) (Response, error) {
	resp := Response{}
	if code != http.StatusOK {
		return resp, NewEasemobError(code, response)
	}
	if err := json.Unmarshal([]byte(response), &resp); err != nil {
		return resp, NewEasemobError(code, response)
	}
	return resp, nil
}
