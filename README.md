Shuffle
=======

This library provides a function for shuffling a list of items in O(n) time.

# Usage

Anything that conforms to the `shuffle.Interface` interface can be shuffled.

```Go
type Interface interface {
	Len() int
	Swap(i, j int)
}
```

Notice that this interface is a subset of the `sort.Interface` interface.

```Go
func Shuffle(deck Interface, rand func() int)
```

The `shuffle.Shuffle` function takes as its second argument a function that returns a random integer. It is used to select random items to be swapped. One could use the `rand.Int` function from the `math/rand` package for this.

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

Since the `sort.Interface` is a superset of the `shuffle.Interface` one can re-write this sample as:

```Go
package main

import (
	"fmt"
	"github.com/gonutz/shuffle"
	"math/rand"
	"sort"
)

func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8}
	shuffle.Shuffle(sort.IntSlice(deck), rand.Int)
	fmt.Println(deck)
}
```
