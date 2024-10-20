package constants

import "errors"

var (
	ERR_NO_TOKENS             = errors.New("NO AVAILABLE TOKEN")
	ERR_TOKEN_ALREADY_DELETED = errors.New("TOKEN ALREADY DELTED OR DOES NOR EXIST")
	DB_OPERATION_ERR          = errors.New("DB OPERATION ERROR")
)
