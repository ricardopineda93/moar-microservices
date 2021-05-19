package shortener

// Defining the interface that the RedirectService implementation must satisfy
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
