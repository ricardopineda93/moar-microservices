package msgpack

import (
	"github.com/vmihailenco/msgpack/v5"

	"github.com/pkg/errors"
	"github.com/rjjp5294/url-shortener/shortener"
)

// Again, this file is essentially a class declaration with methods to satisfy
// the shortener/serializer interface
type Redirect struct{}

// Simply implement the methods expected by the shortener/serializer interface in the
// msgpack flavor
func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
