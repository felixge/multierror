package multierror

import (
	"errors"
	"testing"
)

func TestAppend(t *testing.T) {
	tests := []struct {
		Err     error
		Errs    []error
		WantErr string
		WantNil bool
	}{
		{
			WantNil: true,
		},
		{
			Errs:    []error{nil},
			WantNil: true,
		},
		{
			Errs:    []error{nil, nil},
			WantNil: true,
		},
		{
			Err:     errors.New("A"),
			WantErr: "1 error(s): A",
		},
		{
			Errs:    []error{errors.New("A")},
			WantErr: "1 error(s): A",
		},
		{
			Err:     errors.New("A"),
			Errs:    []error{errors.New("B")},
			WantErr: "2 error(s): A, B",
		},
		{
			Errs:    []error{errors.New("A"), errors.New("B")},
			WantErr: "2 error(s): A, B",
		},
		{
			Errs:    []error{errors.New("A"), nil, errors.New("B")},
			WantErr: "2 error(s): A, B",
		},
	}
	for i, test := range tests {
		err := Append(test.Err, test.Errs...)
		if (err == nil) != test.WantNil {
			t.Errorf("test %d: got=%#v want=nil", i, err)
		}
		var got string
		if err != nil {
			got = err.Error()
		}
		if got != test.WantErr {
			t.Errorf("test %d: got=%q want=%q", i, got, test.WantErr)
		}
	}
}
