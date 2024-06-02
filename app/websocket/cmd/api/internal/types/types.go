// Code generated by goctl. DO NOT EDIT.
package types

type GetInboxReq struct {
}

type GetInboxResp struct {
	Messages []*Message `json:"messages"`
}

type Message struct {
	ToUserId   string `json:"toUserId"`
	FromUserId string `json:"fromUserId"`
	Content    string `json:"content"`
	Type       string `json:"type"`
	CreateTime int64  `json:"createTime"`
}

type User struct {
	Id          int64   `json:"id"`
	Mobile      string  `json:"mobile"`
	Nickname    string  `json:"nickname"`
	Sex         int64   `json:"sex"`
	Avatar      string  `json:"avatar"`
	Info        string  `json:"info"`
	HomeStayIds []int64 `json:"homestayIds"`
}
