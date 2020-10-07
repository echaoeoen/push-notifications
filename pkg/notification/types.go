package notification

import "time"

type FCMToken struct {
	FCMToken string `json:"fcm_token"`
}
type UserData struct {
	Application string `json:"application"`
	Username    string `json:"username"`
	FCMToken
}

type Content struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	SubTitle string    `json:"subtitle"`
	Message  string    `json:"message"`
	Action   string    `json:"action"`
	Param    string    `json:"param"`
	Readed   bool      `json:"readed"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

type Pagination struct {
	Page int64  `json:"page"`
	Size int64  `json:"size"`
	Next string `json:"next,omitempty"`
}

type PaginatedContent struct {
	Pagination
	Data []Content `json:"data"`
}
