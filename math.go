package utils

type comparableNumbers interface {
  int|int8|int16|int32|int64|float32|float64  
}

func Max[T comparableNumbers](st ...T) T {
  if len(st) == 0 {
    return 0
  }
  max := st[0]
  for _, e := range st {
    if (e > max) {
      max = e
    }
  }
  return max
}

func Min[T comparableNumbers](st ...T) T {
  if len(st) == 0 {
    return 0
  }
  min := st[0]
  for _, e := range st {
    if (e < min) {
      min = e
    }
  }
  return min
}
