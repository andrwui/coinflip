package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// var airSequence = []string{"__", "\\", "|", "/"}
var airSequence = []string{"|", "/", "__", "\\"}

func main() {
	speed := flag.Int("s", 100, "The time it takes to render each frame of the coin animation.")
	noCoin := flag.Bool("n", false, "If present, skips the animation straight to the result.")

	flag.Parse()

	if !*noCoin {
		fmt.Print("\033[?25l")

		for i := 0; i < 10; i++ {
			fmt.Print("\n")
		}

		for i := 0; i < 20; i++ {
			parabolic := 1.0 - 1.5*float64((i-10)*(i-10))/float64(100)
			if parabolic < 0 {
				parabolic = 0
			}

			if i < 10 {
				fmt.Print("\033[A")
			}
			if i > 10 {
				fmt.Print("\033[B")
			}

			fmt.Print("\033[93m " + airSequence[i%len(airSequence)])
			time.Sleep(time.Duration((float64(*speed) * parabolic)) * time.Millisecond)
			EraseLines(1)
		}

	}

	random := rand.Intn(1000)

	if random == 0 {
		fmt.Print("\033[93m " + "|")
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\033[93m" + " Edge!")

	} else {

		fmt.Print("\033[93m " + "__")
		fmt.Print(" ")
		time.Sleep(500 * time.Millisecond)

		if (random % 2) == 0 {
			fmt.Print("\033[93m" + "Heads!")
		} else {
			fmt.Print("\033[93m" + "Tails!")
		}

	}
	fmt.Print("\033[?25h")
}

func EraseLines(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Print("\033[2K")
		if i < lines-1 {
			fmt.Print("\033[A")
		}
	}
	fmt.Print("\r")
}
