package pagespeed

import (
	"encoding/json"
)

type RuntimeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r RuntimeError) Error() string {
	return r.Message
}

func (r *RuntimeError) UnmarshalJSON(data []byte) error {

	v := struct {
		Err struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}{}

	err := json.Unmarshal(data, &v)

	*r = *(*RuntimeError)(&(v.Err))

	return err
}
