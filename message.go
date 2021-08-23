package easemob

import (
	"github.com/go-resty/resty/v2"
)

type targetType string
type msgType string

const (
	TargetTypeUsers      targetType = "users"      // 给用户发消息
	TargetTypeChatGroups targetType = "chatgroups" // 给群发消息
	TargetTypeChatRooms  targetType = "chatrooms"  // 给聊天室发消息

	TypeTxt    msgType = "txt"    // 文本消息
	TypeImg    msgType = "img"    // 图片消息
	TypeLoc    msgType = "loc"    // 位置消息
	TypeAudio  msgType = "audio"  // 语音消息
	TypeVideo  msgType = "video"  // 视频消息
	TypeFile   msgType = "file"   // 文件消息
	TypeCustom msgType = "custom" // 自定义消息
)

type message struct {
	Base *Base
}

type MessageRequest struct {
	TargetType targetType  `json:"target_type"`
	Target     []string    `json:"target"`
	Msg        H           `json:"msg"`
	From       string      `json:"from"`
	Ext        interface{} `json:"ext,omitempty"`
}

// SendTxt 发送文本消息
// MessageRequest.Msg H{"msg":"发送第一条消息"}
func (m *message) SendTxt(params *MessageRequest) (Response, error) {
	params.Msg["type"] = TypeTxt
	return m.send(params)
}

// SendImage 发送图片消息
// MessageRequest.Msg H{"url":"成功上传文件返回的UUID","filename":"语音名称","secret":"成功上传文件后返回的secret","size":"图片尺寸；height：高度，width：宽度 map[string]int"}
func (m *message) SendImage(params *MessageRequest) (Response, error) {
	params.Msg["type"] = TypeImg
	return m.send(params)
}

// SendLoc 发送位置消息
// MessageRequest.Msg H{"lat":"纬度","lng":"经度","addr":"地址，例如：中国北京市海淀区中关村"}
func (m *message) SendLoc(params *MessageRequest) (Response, error) {
	params.Msg["type"] = TypeLoc
	return m.send(params)
}

// SendAudio 发送语音消息
// MessageRequest.Msg H{"url":"成功上传文件返回的UUID","filename":"语音名称","secret":"成功上传文件后返回的secret","length":"语音时间（单位：秒）"}
func (m *message) SendAudio(params *MessageRequest) (Response, error) {
	params.Msg["type"] = TypeAudio
	return m.send(params)
}

// SendVideo 发送视频消息
// MessageRequest.Msg H{
//		"filename":"文件名称", "thumb": "成功上传视频缩略图返回的UUID", "length": "视频播放长度 int",
//		"secret": "成功上传视频文件后返回的secret", "file_length":"视频文件大小（单位：字节）int",
//		"thumb_secret":"成功上传视频缩略图后返回的secret", "url": "成功上传视频文件返回的UUID"
//	}
func (m *message) SendVideo(params *MessageRequest) (Response, error) {
	params.Msg["type"] = TypeVideo
	return m.send(params)
}

// SendFile 发送文件消息
func (m *message) SendFile(params *MessageRequest) (Response, error) {
	params.Msg["type"] = TypeFile
	return m.send(params)
}

// SendCustom 发送自定义消息
// MessageRequest.Msg H{"customEvent":"用户自定义的事件类型", "customExts": "用户自定义的事件属性,类型必须是map[string]string"}
func (m *message) SendCustom(params *MessageRequest) (Response, error) {
	params.Msg["type"] = TypeCustom
	return m.send(params)
}

func (m *message) send(params *MessageRequest) (Response, error) {
	data := H{"target_type": params.TargetType, "target": params.Target, "msg": params.Msg}
	if params.From != "" {
		data["from"] = params.From
	}
	return m.Base.request(m.Base.config.URL+"/messages", resty.MethodPost, data)
}
