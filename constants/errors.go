package constants

import "errors"

var (
	ERR_NO_TOKENS             = errors.New("NO_AVAILABLE_TOKEN")
	ERR_TOKEN_ALREADY_DELETED = errors.New("TOKEN_ALREADY_DELTED_OR_DOES_NOR_EXIST")
	DB_OPERATION_ERR          = errors.New("DB_OPERATION_ERROR")
	TOKEN_UNBLOCK_ERR         = errors.New("TOKEN_UNVBLOCK_ERROR")
)
