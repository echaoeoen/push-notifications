package notification

import "time"

type UserData struct {
	Application string `json:"application"`
	Username    string `json:"username"`
	FCMToken    string `json:"fcm_token"`
}

type Content struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	SubTitle string    `json:"sub_title"`
	Message  string    `json:"message"`
	Action   string    `json:"action"`
	Param    string    `json:"param"`
	Readed   string    `json:"readed"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

type Paginatation struct {
	Page int64  `json:"page"`
	Size int64  `json:"size"`
	Next string `json:"next,omitempty"`
}

type PaginatedContent struct {
	Paginatation
	Data []Content `json:"data"`
}
