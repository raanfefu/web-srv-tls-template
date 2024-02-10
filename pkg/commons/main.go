package commons

import "flag"

func StringIsRequiered(v *string) error {

	if *v == "" {
		return flag.ErrHelp
	}
	return nil
}

func Int64IsRequiered(v int64) error {
	if v == 0 {
		return flag.ErrHelp
	}
	return nil
}
