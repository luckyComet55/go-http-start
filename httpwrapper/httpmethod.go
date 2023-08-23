package httpwrapper

type httpMethod string

const _SUPPORTED_METHOD_NUM = 5

func (m httpMethod) isValid() bool {
	switch m {
	case "GET", "POST", "PATCH", "PUT", "DELETE":
		return true
	default:
		return false
	}
}
