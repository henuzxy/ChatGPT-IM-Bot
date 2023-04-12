package qqbot

type MessageSegment struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}
type QQGroupMessage struct {
	PostType    string           `json:"post_type"`
	MessageType string           `json:"message_type"`
	Time        int              `json:"time"`
	SelfID      int              `json:"self_id"`
	SubType     string           `json:"sub_type"`
	GroupID     int              `json:"group_id"`
	Message     []MessageSegment `json:"message"`
	Sender      Sender           `json:"sender"`
	Anonymous   interface{}      `json:"anonymous"`
	Font        int              `json:"font"`
	MessageSeq  int              `json:"message_seq"`
	RawMessage  string           `json:"raw_message"`
	UserID      int              `json:"user_id"`
	MessageID   int              `json:"message_id"`
}
type Sender struct {
	Age      int    `json:"age"`
	Area     string `json:"area"`
	Card     string `json:"card"`
	Level    string `json:"level"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
	Sex      string `json:"sex"`
	Title    string `json:"title"`
	UserID   int    `json:"user_id"`
}
