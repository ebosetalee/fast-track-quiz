package quiz

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var ErrEmptyBody = errors.New("body must not be empty")

func ReadJSON(req *http.Request, dst interface{}) error {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	err := json.NewDecoder(req.Body).Decode(dst)
	if err != nil {
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %v", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return ErrEmptyBody

		default:
			return err
		}

	}

	return nil
}
