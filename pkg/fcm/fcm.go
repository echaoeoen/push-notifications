package fcm

import (
	"encoding/json"

	"github.com/appleboy/go-fcm"
	"github.com/oeoen/push-notifications/driver/config"
	"github.com/oeoen/push-notifications/helper/errorp"
	"github.com/oeoen/push-notifications/pkg/notification"
)

type FCM struct {
	c config.Provider
}

func New(c config.Provider) *FCM {
	return &FCM{
		c: c,
	}
}
func structJSONToMap(c interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return nil, errorp.ParseError(err.Error())
	}
	var r = map[string]interface{}{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, errorp.ParseError(err.Error())
	}
	return r, nil
}
func (f *FCM) Send(destToken string, content notification.Content) error {
	// Create the message to be sent.
	data, err := structJSONToMap(content)
	if err != nil {
		return err
	}
	msg := &fcm.Message{
		To:   destToken,
		Data: data,
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient(f.c.FCMServerKey())
	if err != nil {
		return errorp.FCMConnError(err.Error())
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		return errorp.FCMConnError(err.Error())
	}
	f.c.Logger().Info(response)
	return nil
}
