Shuffle
=======

This library provides a function for shuffling a list of items in O(n) time.

# Example

```Go
package main

import (
	"fmt"
	"github.com/gonutz/shuffle"
	"math/rand"
)

func main() {
	d := deck{1, 2, 3, 4, 5, 6, 7, 8}
	shuffle.Shuffle(d, rand.Int)
	fmt.Println(d)
}

type deck []int

func (d deck) Len() int      { return len(d) }
func (d deck) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
```

This program shuffles an integer slice, sample output: `[2 6 3 5 1 8 7 4]`
