package queries

var FetchNotification = `
	SELECT
		id,
		title,
		subtitle,
		message,
		action,
		param,
		readed,
		created,
		updated
	FROM 
		notifications
	:where 
	ORDER BY created DESC LIMIT ?, ?
`
var InsertNotification = `
	INSERT INTO 
	notifications(
		application,
		username,
		title,
		subtitle,
		message,
		action,
		param,
		readed,
		created,
		updated
	)
	VALUES(
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?
	)
`
var ReadNotification = `
	UPDATE  
	notifications
	SET
		readed = 1
	WHERE
		application = ? AND
		username = ? AND
		id = ?
`
