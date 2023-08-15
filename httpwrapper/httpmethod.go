package httpwrapper

type httpMethod string

func (m httpMethod) isValid() bool {
	switch m {
	case "GET", "POST", "PATCH", "PUT", "DELETE":
		return true
	default:
		return false
	}
}
