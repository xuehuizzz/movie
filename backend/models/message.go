package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id int
	Content string
	AddTime time.Time
}


type MessageUser struct {
	Id int
	MessageId int64
	AddTime time.Time
	Status int
	UserId int
}

func init(){
	orm.RegisterModel(new(Message), new(MessageUser))
}


//保存通知消息
func SendMessageDo(content string) (int64, error) {
	o := orm.NewOrm()
	var message Message
	message.Content = content
	message.AddTime = time.Now()
	messageId,err := o.Insert(&message)
	return messageId, err
}


//保存消息接收人
func SaveReceiver(userId int, messageId int64)  error {
	o := orm.NewOrm()
	var messageuser MessageUser
	messageuser.UserId = userId
	messageuser.MessageId = messageId
	messageuser.Status = 1
	messageuser.AddTime = time.Now()
	_,err := o.Insert(&messageuser)
	return  err
}