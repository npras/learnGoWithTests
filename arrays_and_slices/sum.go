package main

func Sum(arr []int) int {
  result := 0
  for _, v := range arr {
    result += v
  }
  return result
}
