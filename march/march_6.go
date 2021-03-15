package main

import "fmt"

func main() {
	n, k := 0, 0
	fmt.Scanf("%d %d\n", &n, &k)

	order, last := 0, 0
	count := 0
	for ; n > 0; n-- {
		p := 0
		fmt.Scanf("%d", &p)

		if p > 0 && (order < k || last == p) {
			last = p
			count++
		}

		order++
	}

	fmt.Printf("%d\n", count)
}
