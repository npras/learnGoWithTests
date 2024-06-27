package iteration


import (
  "testing"
  "fmt"
)


func TestRepeater(t *testing.T) {
  got := Repeater("a", 5)
  want := "aaaaa"
  if got != want {
    t.Errorf("want: %q, got: %q", want, got)
  }

  got = Repeater("sally", 6)
  want = "sallysallysallysallysallysally"
  if got != want {
    t.Errorf("want: %q, got: %q", want, got)
  }
}


func BenchmarkRepeater(b *testing.B) {
  for range b.N { Repeater("sally", 1000) }
}


func BenchmarkRepeater2(b *testing.B) {
  for range b.N { Repeater2("sally", 1000) }
}


func ExampleRepeater() {
	str := Repeater("Sally", 4)
	fmt.Println(str)
	// Output: SallySallySallySally
}
