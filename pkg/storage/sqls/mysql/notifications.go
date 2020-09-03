package mysql

import (
	"context"

	"github.com/oeoen/push-notifications/helper/errorp"
	"github.com/oeoen/push-notifications/pkg/notification"
	"github.com/oeoen/push-notifications/pkg/storage/sqls/mysql/queries"
)

func (m *MYSQLManager) FetchNotification(ctx context.Context, filter ...[3]string) ([]*notification.Content, error) {
	query, k := constructWhereQuery(queries.FetchNotification, filter)
	ns := []*notification.Content{}
	sp := startSpan(ctx, "fetch-notification", query)
	defer finishSpan(sp)
	r, err := m.DBService().Query(query, k...)
	if err != nil {
		return nil, errorp.FetchError(err.Error())
	}
	defer r.Close()
	for r.Next() {
		n := notification.Content{}
		err = r.Scan(
			&n.ID,
			&n.Title,
			&n.SubTitle,
			&n.Message,
			&n.Action,
			&n.Param,
			&n.Readed,
			&n.Created,
			&n.Updated,
		)
		if err != nil {
			return nil, errorp.FetchError(err.Error())
		}
		ns = append(ns, &n)
	}
	if len(ns) == 0 {
		return nil, errorp.NotFoundError("No notification data")
	}
	return ns, nil
}

func (m *MYSQLManager) SaveNotification(ctx context.Context, application, username string, content notification.Content) error {
	query := queries.InsertNotification
	stmt, err := m.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		application,
		username,
		content.Title,
		content.SubTitle,
		content.Message,
		content.Action,
		content.Param,
		content.Readed,
	)
	if err != nil {
		return errorp.InsertError(err.Error())
	}
	return nil
}
func (m *MYSQLManager) ReadNotification(ctx context.Context, application, username, notificationID string) error {
	query := queries.ReadNotification
	stmt, err := m.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		application,
		username,
		notificationID,
	)
	if err != nil {
		return errorp.UpdateError(err.Error())
	}
	return nil
}
