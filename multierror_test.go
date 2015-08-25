package multierror

import (
	"errors"
	"testing"
)

func TestAppend(t *testing.T) {
	tests := []struct {
		Err  error
		Errs []error
		Want string
	}{
		{},
		{
			Err:  errors.New("A"),
			Want: "1 error(s): A",
		},
		{
			Errs: []error{errors.New("A")},
			Want: "1 error(s): A",
		},
		{
			Err:  errors.New("A"),
			Errs: []error{errors.New("B")},
			Want: "2 error(s): A, B",
		},
		{
			Errs: []error{errors.New("A"), errors.New("B")},
			Want: "2 error(s): A, B",
		},
		{
			Errs: []error{errors.New("A"), nil, errors.New("B")},
			Want: "2 error(s): A, B",
		},
	}
	for i, test := range tests {
		err := Append(test.Err, test.Errs...)
		var got string
		if err != nil {
			got = err.Error()
		}
		if got != test.Want {
			t.Errorf("test %d: got=%q want=%q", i, got, test.Want)
		}
	}
}
