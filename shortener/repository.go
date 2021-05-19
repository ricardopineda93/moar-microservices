package shortener

// Defining the interface the repository(s) must implement and satisfy
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
