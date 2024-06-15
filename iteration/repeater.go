package iteration

// Repeats `str` `n` times and returns the
// concatenated string.
func Repeater(str string, n int) string {
  result := ""
  for range n {
    result = result + str
  }
  return result
}

func Repeater2(str string, n int) (result string) {
  for i := 0; i < n; i++ {
    result += str
  }
  return
}
