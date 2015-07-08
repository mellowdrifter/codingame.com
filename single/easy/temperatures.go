package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func abs(x int) int {
	if (x < 0) {
		return -x
	}
	return x
}


func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // N: the number of temperatures to analyse
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    answer := 10000
    positive := false // IF only negative values, don't use absolute values
    
    scanner.Scan()
    TEMPS := scanner.Text() // the N temperatures expressed as integers ranging from -273 to 5526
    split_TEMPS := strings.Split(TEMPS, " ")
    

    // As we only care about positives, use absolute values
    for i := 0; i < N; i++ {
        actual, _ := strconv.Atoi(split_TEMPS[i])
        if (actual > 0) {
            positive = true   
        }
        
        number := abs(actual)
        
        if (number < answer) {
            answer = number
        }
    }
    
    
    if (answer == 10000) {
        fmt.Printf("0\n")
    } else if (positive) {
        fmt.Printf("%d\n", answer)
    } else {
        fmt.Printf("%d\n", -answer)
    }
}
