package main

import "fmt"

func main() {
	i := 0
	for {
		if i == 5 {
			fmt.Println("Skipping 5")
			i += 1
			continue
		}

		fmt.Println(i)

		if i >= 100 {
			fmt.Println("Breaking out of the loop")
			break
		}

		i += 1
	}
}
