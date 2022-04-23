package lib

type Result[T any, E any] struct {
	value *T
	error *E
}

func (result Result[T, E]) Ok() (bool, *T) {
	return result.value != nil, result.value
}

func (result Result[T, E]) Err() (bool, *E) {
	return result.error != nil, result.error
}

func Try[T any, E any](try func() Result[T, E]) Result[T, E] {
	return try()
}

func (result Result[T, E]) Catch(catch func(err E)) Result[T, E] {
	if result.error != nil {
		catch(*result.error)
	}

	return result
}

func (result Result[T, E]) Finally(finally func(value T) Result[T, E]) Result[T, E] {
	if result.error == nil {
		return finally(*result.value)
	}

	return result
}

func ResultOk[T any, E any](value T) Result[T, E] {
	return Result[T, E]{
		value: &value,
		error: nil,
	}
}

func ResultErr[T any, E any](err E) Result[T, E] {
	return Result[T, E]{
		value: nil,
		error: &err,
	}
}
