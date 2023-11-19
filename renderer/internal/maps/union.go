package maps

import "maps"

func Union[M interface{ ~map[K]V }, K comparable, V any](mm ...M) M {
	result := make(map[K]V)
	for _, m := range mm {
		maps.Copy(result, m)
	}
	return result
}
