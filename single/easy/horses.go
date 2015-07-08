package main

import (
    "fmt"
    "math"
    "sort"
    )

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    var N int
    fmt.Scan(&N)
    
    // Create slice of INTs
    list := make([]int, N)
    
    // Smallest difference
    difference := 100000
    
    
    for i := 0; i < N; i++ {
        //var Pi int
        fmt.Scan(&list[i])
    }
    
    // Initial sort of list
    sort.Sort(sort.IntSlice(list))
    
    for i := 1; i < len(list); i++ {
        currentDiff := int(math.Abs(float64(list[i] - list[i - 1])))
        if (currentDiff < difference) {
                difference = currentDiff
            }
    }
    fmt.Println(difference)
}
