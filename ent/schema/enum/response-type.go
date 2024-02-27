package enum

type ResponseType string

// enumeration for Linux process response type

const (
	CORRECT               ResponseType = "CORRECT"
	WRONG_ANSWER          ResponseType = "WRONG_ANSWER"
	TIME_LIMIT_EXCEED     ResponseType = "TIME_LIMIT_EXCEED"
	COMPILE_ERROR         ResponseType = "COMPILE_ERROR"
	RUNTIME_ERROR_SIGSEGV ResponseType = "RUNTIME_ERROR_SIGSEGV"
	RUNTIME_ERROR_SIGXFSZ ResponseType = "RUNTIME_ERROR_SIGXFSZ"
	RUNTIME_ERROR_SIGFPE  ResponseType = "RUNTIME_ERROR_SIGFPE"
	RUNTIME_ERROR_SIGABRT ResponseType = "RUNTIME_ERROR_SIGABRT"
	RUNTIME_ERROR_NZEC    ResponseType = "RUNTIME_ERROR_NZEC"
	RUNTIME_ERROR         ResponseType = "RUNTIME_ERROR"
	INTERNAL_ERROR        ResponseType = "INTERNAL_ERROR"
	EXEC_FORMAT_ERROR     ResponseType = "EXEC_FORMAT_ERROR"
	MEMORY_LIMIT_EXCEED   ResponseType = "MEMORY_LIMIT_EXCEED"
)

// Implementation of ent EnumValues
// https://pkg.go.dev/entgo.io/ent/schema/field#EnumValues
func (ResponseType) Values() (types []string) {
	for _, s := range []ResponseType{
		CORRECT,
		WRONG_ANSWER,
		TIME_LIMIT_EXCEED,
		COMPILE_ERROR,
		RUNTIME_ERROR_SIGSEGV,
		RUNTIME_ERROR_SIGXFSZ,
		RUNTIME_ERROR_SIGFPE,
		RUNTIME_ERROR_SIGABRT,
		RUNTIME_ERROR_NZEC,
		RUNTIME_ERROR,
		INTERNAL_ERROR,
		EXEC_FORMAT_ERROR,
		MEMORY_LIMIT_EXCEED,
	} {
		types = append(types, string(s))
	}
	return
}
