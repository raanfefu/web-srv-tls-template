package server

type ModeType string

type ServerParams struct {
	Mode ModeType
	Port int64
	Crt  string
	Key  string
}

const (
	TLS  ModeType = "https"
	HTTP ModeType = "http"
)

func (c ModeType) String() string {
	return string(c)
}
