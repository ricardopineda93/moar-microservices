package shortener

// Defining the interface that the serializer(s) must implement and satisfy
type RedirectSerializer interface {
	Decode(input []byte) (*Redirect, error)
	Encode(input *Redirect) ([]byte, error)
}
