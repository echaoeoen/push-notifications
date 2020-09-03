package mysql

import (
	"context"

	"github.com/oeoen/push-notifications/helper/errorp"
	"github.com/oeoen/push-notifications/pkg/notification"
	"github.com/oeoen/push-notifications/pkg/storage/sqls/mysql/queries"
)

func (m *MYSQLManager) SetFCMToken(ctx context.Context, application, username string, token notification.FCMToken) error {
	_, err := m.GetFCMToken(ctx, application, username)
	if err != nil {
		return m.insertFMCToken(ctx, application, username, token)
	}
	return m.updateFMCToken(ctx, application, username, token)
}

func (m *MYSQLManager) insertFMCToken(ctx context.Context, application, username string, token notification.FCMToken) error {
	query := queries.InsertUserData
	stmt, err := m.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		application,
		username,
		token.FCMToken,
	)
	if err != nil {
		return errorp.InsertError(err.Error())
	}
	return nil
}
func (m *MYSQLManager) updateFMCToken(ctx context.Context, application, username string, token notification.FCMToken) error {
	query := queries.UpdateFCMToken
	stmt, err := m.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		token.FCMToken,
		application,
		username,
	)
	if err != nil {
		return errorp.UpdateError(err.Error())
	}
	return nil
}
func (m *MYSQLManager) GetFCMToken(ctx context.Context, application, username string) (*notification.UserData, error) {
	query := queries.GetUserData
	row, err := m.DBService().Query(query, application, username)
	if err != nil {
		return nil, errorp.FetchError(err.Error())
	}
	if !row.Next() {
		return nil, errorp.NotFoundError("No data FCM with that user")
	}
	userdata := notification.UserData{}
	if err = row.Scan(
		&userdata.Application,
		&userdata.Username,
		&userdata.FCMToken.FCMToken,
	); err != nil {
		return nil, errorp.FetchError(err.Error())
	}
	return &userdata, nil
}
