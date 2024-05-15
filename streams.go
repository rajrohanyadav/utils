package utils

func Contains[T comparable](val T, itr []T) bool {
	for _, e := range itr {
		if e == val {
			return true
		}
	}
	return false
}

func Map[T any](apply func(T) T, itr []T) []T{
	res := make([]T, len(itr))
	for i, e := range itr {
		res[i] = apply(e)
	}
	return res
}

func Filter() {}

func Reduce() {}

func Reverse[T any](itr []T) {
  for i, j := 0, len(itr) - 1; i < j; i, j = i+1, j-1 {
    itr[i], itr[j] = itr[j], itr[i]
  }
}
