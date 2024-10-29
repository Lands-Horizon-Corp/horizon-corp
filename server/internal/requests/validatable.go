package requests

type Validatable interface {
	Validate() error
}
