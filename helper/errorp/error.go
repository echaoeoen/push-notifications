package errorp

func FetchError(message string) Error {
	return NewNotificationError(
		500,
		"fetch_err",
		"Fetch error, please contact administrator",
		message,
	)
}

func NotFoundError(message string) Error {
	return NewNotificationError(
		404,
		"not_found",
		"Data not found",
		message,
	)
}

func ConflictError(message string) Error {
	return NewNotificationError(
		409,
		"err_conflict",
		"Data conflict, please check your request",
		message,
	)
}

func IncorrectPayloadError(message string) Error {
	return NewNotificationError(
		400,
		"err_incorrect_payload",
		"Please check your body request",
		message,
	)
}

func DBConnectionError(message string) Error {
	return NewNotificationError(
		500,
		"err_db_conn",
		"Pelase check your db connection config",
		message,
	)
}
