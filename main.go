package main

import "fmt"

type Dual struct {
    first        int
    second       int
    strengthDiff int
}

func SortDesc(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr) - 1; j++ {
            if arr[i] > arr[j] {
                tmp := arr[i]
                arr[i] = arr[j]
                arr[j] = tmp
            }
        }
    }

    return arr
}

func main() {
    var horsesNumber int

    fmt.Println("How much horses do you want to add ?")
    fmt.Scanf("%d", &horsesNumber)

    if horsesNumber <= 1 {
        fmt.Println("Please, add more than 1!")

        return
    }

    strengthArr := make([]int, horsesNumber)

    for key, _ := range strengthArr {
        var strength int

        fmt.Printf("Add strength to %d horse\n", key+1)
        fmt.Scanln(&strength)

        strengthArr[key] = strength
    }

    strengthArr = SortDesc(strengthArr)

    dual := Dual{}

    for i := 0; i < len(strengthArr) - 1; i++ {
        first := strengthArr[i]
        second := strengthArr[i+1]
        diff := first - second

        if i == 0 || diff <= dual.strengthDiff {
            dual = Dual{
                first:        first,
                second:       second,
                strengthDiff: diff,
            }
        }
    }

    fmt.Printf(
        "The two closest strengths are %d and %d, with a difference of %d.\n",
        dual.first,
        dual.second,
        dual.strengthDiff,
    )
}