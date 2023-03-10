package main

import "sort"

type Horse struct {
	Name  string
	Power int
}

type HorsePair struct {
	Horse1 Horse
	Horse2 Horse
}

func main() {
	horses := make([]Horse, 0)
	horses = append(horses, Horse{Name: "Лошадь 0", Power: 9})
	horses = append(horses, Horse{Name: "Лошадь 1", Power: 10})
	horses = append(horses, Horse{Name: "Лошадь 2", Power: 17})
	horses = append(horses, Horse{Name: "Лошадь 3", Power: 20})
	horses = append(horses, Horse{Name: "Лошадь 4", Power: 11})
	horses = append(horses, Horse{Name: "Лошадь 5", Power: 10})
	horses = append(horses, Horse{Name: "Лошадь 6", Power: 20})
	horses = append(horses, Horse{Name: "Лошадь 7", Power: 22})

	pairs := FindSameHorse(horses)
	for _, pair := range pairs {
		println(pair.Horse1.Power, pair.Horse2.Power)
	}
}

func FindSameHorse(horses []Horse) []HorsePair {
	sort.Slice(horses, func(i, j int) bool {
		return horses[i].Power < horses[j].Power
	})

	pairs := []HorsePair{}
	for i := 0; i < len(horses)-1; i++ {
		if horses[i+1].Power-horses[i].Power <= 1 {
			pair := HorsePair{Horse1: horses[i], Horse2: horses[i+1]}
			pairs = append(pairs, pair)
		}
	}

	return pairs
}
