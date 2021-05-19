package json

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/rjjp5294/url-shortener/shortener"
)

// Just the struct that will house the interface method implementations. This file
// essentially just is a class declaration of with methods that satisfy the interface expected.
type Redirect struct{}

// The methods defined here are just fulfilling the shortener/serializer interface
// in the JSON flavor
func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
