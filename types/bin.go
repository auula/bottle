package types

import "bytes"

type Binary struct {
	buf bytes.Buffer
}

func (bin *Binary) ToBSON() ([]byte, error) {
	return nil, nil
}
