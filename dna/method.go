package dna

type Method int

const (
	GET Method = iota
	POST
	DELETE
	UPDATE
	NO_METHOD
)

func (m Method) String() string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case DELETE:
		return "DELETE"
	case UPDATE:
		return "UPDATE"
	}
	return ""
}

func Str2Method(s string) (Method, error) {
	switch s {
	case "GET":
		return GET, nil
	case "POST":
		return POST, nil
	case "DELETE":
		return DELETE, nil
	case "UPDATE":
		return UPDATE, nil
	default:
		return NO_METHOD, ErrInvalidMethod
	}
}
