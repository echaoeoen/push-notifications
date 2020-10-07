package notification

import "context"

type StorageManager interface {
	SetFCMToken(ctx context.Context, application, username string, token FCMToken) error
	GetFCMToken(ctx context.Context, application, username string) (*UserData, error)
	SaveNotification(ctx context.Context, application, username string, content Content) error
	FetchNotification(ctx context.Context, application, username string, filter ...[3]string) ([]*Content, error)
	ReadNotification(ctx context.Context, application, username, notificationID string) error
	UnreadCountNotification(ctx context.Context, application, username string) (int64, error)
}

type Manager interface {
	StorageManager() StorageManager
	SendNotification(ctx context.Context, application, username string, content Content) error
}
