package shuffle

type Interface interface {
	Len() int
	Swap(i, j int)
}

func Shuffle(deck Interface, rand func() int) {
	n := deck.Len()
	for i := 0; i < n-1; i++ {
		j := i + rand()%(n-i)
		if i != j {
			deck.Swap(i, j)
		}
	}
}
