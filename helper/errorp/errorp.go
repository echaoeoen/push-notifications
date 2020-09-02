package errorp

type Error interface {
	Error() string
	Code() string
	Status() int64
	Description() string
	Hint() string
	WithCode(c string) Error
	WithDescription(d string) Error
	WithHint(h string) Error
}

type NotificationError struct {
	C string `json:"error_code"`
	D string `json:"error_description"`
	H string `json:"error_hint"`
	S int64  `json:"-"`
}

func NewNotificationError(statusCode int64, code, hint, description string) *NotificationError {
	e := NotificationError{S: statusCode, H: hint, D: description, C: code}
	return &e
}

func (e *NotificationError) Error() string {
	return "Error " + e.C + ": " + e.D
}

func (e *NotificationError) Code() string {
	return e.C
}
func (e *NotificationError) Description() string {
	return e.D
}
func (e *NotificationError) Hint() string {
	return e.H
}
func (e *NotificationError) WithCode(c string) Error {
	e.C = c
	return e
}
func (e *NotificationError) WithDescription(d string) Error {
	e.D = d
	return e
}
func (e *NotificationError) WithHint(h string) Error {
	e.H = h
	return e
}
func (e *NotificationError) Status() int64 {
	return e.S
}
