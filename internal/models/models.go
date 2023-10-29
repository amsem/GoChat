package models

import "time"


type User struct {
    ID int64  `json:"id"`
    Email  string `json:"email"`
    Name   string `json:"name"`
}

type RoomChat struct {
    ID int64 `json:"id"`
    Name string `json:"name"`
}

type Message struct {
    ID int64 `json:"id"`
    Msg_from int64 `json:"msg_from"`
    Msg_to int64 `json:"msg_to"`
    UserID int64 `json:"user_id"`
    Content string `json:"content"`
    Created time.Time `json:"created"`
}

type MessageGroup struct {
    ChannelID int64 `json:"channel_id"`
    MessageID int64 `json:"message_id"`
    UserID int64 `json:"user_id"`
    Content string `json:"content"`
    Created time.Time `json:"created"`
}
