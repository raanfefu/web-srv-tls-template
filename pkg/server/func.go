package server

import (
	"fmt"
)

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
