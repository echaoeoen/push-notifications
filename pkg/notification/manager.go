package notification

import "context"

type StorageManager interface {
	SetFCMToken(ctx context.Context, userdata UserData) error
	GetFCMToken(ctx context.Context, application, username string) (UserData, error)
	SaveNotification(ctx context.Context, application, username string, content Content) (UserData, error)
	FetchNotification(ctx context.Context, application, username string, filter ...[3][]string) ([]Content, error)
	ReadNotification(ctx context.Context, notificationID string) error
}

type Manager interface {
	StorageManager
	SendNotification(ctx context.Context, application, username string, content Content) error
}
