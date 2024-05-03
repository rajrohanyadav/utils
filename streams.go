package utils


// TODO: Use generics here 
func Map(apply func(int) int, itr []int) []int {
  res := make([]int, len(itr))
  for i, e := range itr {
    res[i] = apply(e)
  }

  return res
}

func Filter() {}

func Reduce() {}
