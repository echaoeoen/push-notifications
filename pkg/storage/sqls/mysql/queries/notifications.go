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
	ORDER BY id DESC LIMIT ?, ?
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
		param
	)
	VALUES(
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

var UnreadCountNotification = `
	SELECT count(*)
	FROM
	notifications
	WHERE
	application = ? AND
	username = ? AND
	readed = 0
`
