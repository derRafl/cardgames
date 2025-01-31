package card

import "fmt"

func ExampleCard_String() {
	c1 := New(Ace, Spades)
	c2 := New(Queen, Hearts)
	c3 := New(Ten, Hearts)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)

	// Output:
	// ┌───────┐
	// │A      │
	// │       │
	// │   ♠   │
	// │       │
	// │      A│
	// └───────┘
	// ┌───────┐
	// │Q      │
	// │       │
	// │   ♥   │
	// │       │
	// │      Q│
	// └───────┘
	// ┌───────┐
	// │10     │
	// │       │
	// │   ♥   │
	// │       │
	// │     10│
	// └───────┘
}
