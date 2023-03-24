package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Competitor struct {
	Nr int
	Pi int
}

type Pair struct {
	fNr   int
	sNr   int
	delta int
}

/*
|--------------------------------------------------------------------------
| Task description
|--------------------------------------------------------------------------
|
| The Goal
| Casablanca’s hippodrome is organizing a new type of horse racing: duals. During a dual,
| only two horses will participate in the race.
| In order for the race to be interesting,
| it is necessary to try to select two horses with similar strength.
|
| Write a program which, using a given number of strengths,
| identifies the two closest strengths and shows their difference with an integer (≥ 0).
|
| Input
| Line 1: Number N of horses
|
| The N following lines: the strength Pi of each horse. Pi is an integer.
|
| Output
| The difference D between the two closest strengths. D is an integer greater than or equal to 0.
|
*/
func main() {
	input, err := getUserInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	// make separate slice for the keep competitor Nrs
	// will use competitor Nrs for pretty look output
	competitors := make([]Competitor, 0, len(input))

	// fill it
	for i, pi := range input {
		competitors = append(competitors, Competitor{Nr: i + 1, Pi: pi})
	}

	// sort should guarantee that closest numbers will be in right order
	sort.SliceStable(competitors, func(i, j int) bool {
		return competitors[i].Pi < competitors[j].Pi
	})

	// reverse loop for the get strongest competitors
	// with same delta
	pair := Pair{}

	for i := len(competitors) - 1; i > 0; i-- {
		fNr := competitors[i-1].Nr
		sNr := competitors[i].Nr
		delta := competitors[i].Pi - competitors[i-1].Pi

		// if it's first iteration delta < than prev set new pair
		if i == len(competitors)-1 || delta < pair.delta {
			pair = Pair{fNr: fNr, sNr: sNr, delta: delta}

			if delta == 0 {
				break // bestPairFound
			}
		}
	}

	fmt.Printf("There is best match: No: #%d, No: #%d [delta: %d]. \n",
		pair.fNr,
		pair.sNr,
		pair.delta)
}

func getUserInput() ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter a number of competitors [default: 10] [min: 2]:")
	scanner.Scan()

	inputNr := strings.TrimSpace(scanner.Text())
	number := 10

	if inputNr != "" {
		n, err := strconv.Atoi(inputNr)

		if err != nil {
			return nil, err
		}

		if n > 1 {
			number = n
		} else {
			return nil, fmt.Errorf("invalid number [%d] please provide correct number greater than 2", n)
		}
	}

	input := make([]int, number)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please set competitors Pi between: 1-100:")

	for i := 0; i < number; {
		fmt.Printf("#%d -> ", i+1)
		piInput, err := reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		pi, err := strconv.Atoi(strings.TrimSpace(piInput))

		if err != nil {
			fmt.Println(err)
		}

		// validation
		if pi > 0 && pi <= 100 {
			input[i] = pi
			i++
		}

		fmt.Println("Competitor Pi is required, please specify it between: 1-100.")
	}

	return input, nil
}
