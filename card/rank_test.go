package card

import "fmt"

func ExampleRank_String() {
	r1 := Nine
	r2 := Rank(9)

	fmt.Println(r1)
	fmt.Println(r2)

	// Output:
	// 9
	// J
}
