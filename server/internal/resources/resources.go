package resources

type Transformable[T any] interface {
	ToResource() T
}

type TransformableList[T any] interface {
	ToResourceList() []T
}
