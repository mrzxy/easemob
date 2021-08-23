package easemob

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type group struct {
	Base *Base
}

// GetUsers 分页获取群组成员
func (g *group) GetUsers(groupId string) (Response, error) {
	url := fmt.Sprintf("%s/chatgroups/%s/users", g.Base.config.URL, groupId)
	return g.Base.request(url, resty.MethodGet, H{})
}

// AddUser 添加单个群组成员
func (g *group) AddUser(groupId string, userId string) (Response, error) {
	url := fmt.Sprintf("%s/chatgroups/%s/users/%s", g.Base.config.URL, groupId, userId)
	return g.Base.request(url, resty.MethodPost, H{})
}

// AddUsers 批量添加群组成员
func (g *group) AddUsers(groupId string, userIds []string) (Response, error) {
	url := fmt.Sprintf("%s/chatgroups/%s/users", g.Base.config.URL, groupId)
	return g.Base.request(url, resty.MethodPost, H{"usernames": userIds})
}

// DeleteUser 移除单个群组成员
func (g *group) DeleteUser(groupId string, userId string) (Response, error) {
	url := fmt.Sprintf("%s/chatgroups/%s/users/%s", g.Base.config.URL, groupId, userId)
	return g.Base.request(url, resty.MethodDelete, H{})
}

// DeleteUsers 批量移除群组成员 len(userIds) <= 60
func (g *group) DeleteUsers(groupId string, userIds []string) (Response, error) {
	url := fmt.Sprintf("%s/chatgroups/%s/users/%s", g.Base.config.URL, groupId, strings.Join(userIds, ","))
	return g.Base.request(url, resty.MethodDelete, H{})
}

// GetAdmin 获取群管理员列表
func (g *group) GetAdmin(groupId string) (Response, error) {
	url := fmt.Sprintf("%s/chatgroups/%s/admin", g.Base.config.URL, groupId)
	return g.Base.request(url, resty.MethodGet, H{})
}

// AddAdmin 添加群管理员
func (g *group) AddAdmin(groupId string, userId string) (Response, error) {
	url := fmt.Sprintf("%s/chatgroups/%s/admin", g.Base.config.URL, groupId)
	return g.Base.request(url, resty.MethodGet, H{"newadmin": userId})
}

// DeleteAdmin 移除群管理员
func (g *group) DeleteAdmin(groupId string, userId string) (Response, error) {
	url := fmt.Sprintf("%s/chatgroups/%s/admin/%s", g.Base.config.URL, groupId, userId)
	return g.Base.request(url, resty.MethodDelete, H{})
}

// TransferGroup 转让群组
func (g *group) TransferGroup(groupId string, userId string) (Response, error) {
	url := fmt.Sprintf("%s/chatgroups/%s", g.Base.config.URL, groupId)
	return g.Base.request(url, resty.MethodPut, H{"newowner": userId})
}

// AddMute 添加禁言
// duration 禁言的时间，单位毫秒，如果是“-1”代表永久
func (g *group) AddMute(groupId string, duration int, userIds []string) (Response, error) {
	url := fmt.Sprintf("%s/chatgroups/%s/mute", g.Base.config.URL, groupId)
	return g.Base.request(url, resty.MethodPost, H{"mute_duration": duration, "usernames": userIds})
}

// RemoveMute 移除禁言
func (g *group) RemoveMute(groupId string, userIds []string) (Response, error) {
	url := fmt.Sprintf("%s/chatgroups/%s/mute/%s", g.Base.config.URL, groupId, strings.Join(userIds, ","))
	return g.Base.request(url, resty.MethodPost, H{})
}

// GetMuteUsers 获取禁言列表
func (g *group) GetMuteUsers(groupId string) (Response, error) {
	url := fmt.Sprintf("%s/chatgroups/%s/mute", g.Base.config.URL, groupId)
	return g.Base.request(url, resty.MethodGet, H{})
}
