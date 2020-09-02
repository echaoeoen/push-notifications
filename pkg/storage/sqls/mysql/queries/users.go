package queries

var InsertUserData = `
	INSERT INTO users(application, username, fcm_token) VALUES(?, ?, ?)
`

var DeleteUserData = `
	DELETE FROM users WHERE application = ? AND username = ?
`

var UpdateFCMToken = `
	UPDATE users SET fcm_token = ? WHERE application = ? AND username = ?
`
