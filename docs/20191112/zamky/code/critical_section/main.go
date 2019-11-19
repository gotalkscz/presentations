package main

import "time"

// START OMIT
func main() {
	var sum int

	for i := 1; i <= 1000; i++ {
		go func() {

			// kritická sekce
			sum++ // <-- tato operace není atomická // HL
			// kritická sekce

		}()
	}

	time.Sleep(time.Second * 2)
	println(sum)
}

// END OMIT
