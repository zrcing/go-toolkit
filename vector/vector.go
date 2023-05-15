package vector

func InVectors[T comparable](x T, vectors Vector[T]) bool {
	return vectors.Contains(x)
}

type Vector[T comparable] []T

func (vectors Vector[T]) Contains(x T) bool {
	for _, v := range vectors {
		if v == x {
			return true
		}
	}
	return false
}
