package error_code

const ERR_OK int32						= 0

const (
	ERR_SVR_INTERVAL					= 1001
	ERR_SVR_PUSH						= 1002
	ERR_SVR_HEADER_UNMARSHAL				= 1003
	ERR_SVR_UNKNOWN_MSG					= 1004
	ERR_SVR_MARSHAL_PB					= 1005
	ERR_SVR_UNMARSHAL_PB					= 1006
	ERR_SVR_CONN						= 1007
)

// err_code for db
const (
	ERR_DB_QUERY						= 2001
	ERR_DB_INSERT						= 2002
	ERR_DB_PACKAGE						= 2003
)
