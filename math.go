package utils

type comparableNumbers interface {
  int|int8|int16|int32|int64|float32|float64  
}

func Max[T comparableNumbers](a T, b ...T) T {
  max := a
  for _, e := range b {
    if (e > max) {
      max = e
    }
  }
  return max
}

func Min[T comparableNumbers](a T, b ...T) T {
  min := a
  for _, e := range b {
    if (e < min) {
      min = e
    }
  }
  return min
}
