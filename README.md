## Usage
环信IM简单封装 
[环信官方文档](https://docs-im.easemob.com/)

### Simple 
```go
import "github.com/Daniel66666666/easemob"

client := easemob.NewEasemob(easemob.Config{
    ClientId:     "XXXXXXXXX",
    ClientSecret: "XXXXXXXXXXXXXX",
    URL:          "https://a1.easemob.com/XXXXXXXXXXX/demo",
})

// 发送文本消息
_, err := client.Message.SendTxt(&easemob.MessageRequest{TargetType: easemob.TargetTypeUsers, Target: []string{"user id"}, Msg: easemob.H{"msg": "first msg!"}, From: "user id", Ext: ""})

// 获取群成员
users, err := client.Group.GetUsers("group id")
if err != nil {
    if e, ok := err.(*easemob.Error); ok {
        fmt.Println(e.GetData()) //获取环信返回的错误消息
    }
} else {
    fmt.Println(users.Data)
}

// 账号禁用
_, err := client.User.Deactivate("user id")
```

