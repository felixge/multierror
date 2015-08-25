package multierror

import (
	"fmt"
	"strings"
)

type Error []error

func (e Error) Error() string {
	return e.String()
}

func (e Error) String() string {
	if len(e) == 0 {
		return ""
	}
	var s []string
	for _, err := range e {
		s = append(s, err.Error())
	}
	return fmt.Sprintf("%d error(s): %s", len(s), strings.Join(s, ", "))
}

func Append(err error, errs ...error) error {
	var multiErr Error
	for _, err := range append([]error{err}, errs...) {
		if err != nil {
			multiErr = append(multiErr, err)
		}
	}
	return multiErr
}
