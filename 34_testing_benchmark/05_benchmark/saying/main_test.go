package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Abhi")
	if s != "Welcome my dear Abhi" {
		t.Error("got", s, ", expected:", "Welcome my dear james")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Abhi")
	}
}
