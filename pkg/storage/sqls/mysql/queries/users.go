package queries

var InsertUserData = `
	INSERT INTO users(application, username, fcm_token) VALUES(?, ?, ?)
`

var GetUserData = `
	SELECT application, username, fcm_token from  users WHERE application = ? AND username = ?
`

var DeleteUserData = `
	DELETE FROM users WHERE application = ? AND username = ?
`

var UpdateFCMToken = `
	UPDATE users SET fcm_token = ? WHERE application = ? AND username = ?
`
