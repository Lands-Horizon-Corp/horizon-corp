package requests

type ToModelConvertible[Model any] interface {
	ToModel() *Model
}
