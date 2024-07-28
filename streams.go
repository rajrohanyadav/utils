package utils

func Contains[T comparable](val T, itr []T) bool {
	for _, e := range itr {
		if e == val {
			return true
		}
	}
	return false
}

func Map[T any, U any](apply func(T) U, itr []T) []U{
	res := make([]U, len(itr))
	for i, e := range itr {
		res[i] = apply(e)
	}
	return res
}

func Filter[T any](fn func(T) bool, itr []T) []T{
	res := make([]T, 0)
	for _, e := range itr {
		if fn(e) {
			res = append(res, e)
		}
	}
	return res
}

func Reduce() {}

func Reverse[T any](itr []T) {
  for i, j := 0, len(itr) - 1; i < j; i, j = i+1, j-1 {
    itr[i], itr[j] = itr[j], itr[i]
  }
}
