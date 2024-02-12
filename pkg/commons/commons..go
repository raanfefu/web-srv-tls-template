package commons

import (
	"flag"
	"fmt"
)

type TypeSettings string
type ModeType string

const (
	ArgumentSettings TypeSettings = "args"
	JsonSettings     TypeSettings = "json"
)

const (
	TLS  ModeType = "https"
	HTTP ModeType = "http"
)

func StringIsRequiered(v *string) error {

	if *v == "" {
		return flag.ErrHelp
	}
	return nil
}

func UintRequiered(v uint) error {
	if v == 0 {
		return flag.ErrHelp
	}
	return nil
}

func (c ModeType) String() string {
	return string(c)
}

func ParseMode(s *string) (c *ModeType, err error) {
	if s != nil {
		capabilities := map[ModeType]struct{}{
			HTTP: {},
			TLS:  {},
		}

		mode := ModeType(*s)
		_, ok := capabilities[mode]
		if !ok {
			return c, fmt.Errorf(`cannot parse:[%s] as mode`, s)
		}
		return &mode, nil
	} else {
		return nil, fmt.Errorf(`cannot parse:[%s] as mode`, s)
	}
}
