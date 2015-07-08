package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    for {
        var spaceX, spaceY int
        fmt.Scan(&spaceX, &spaceY)
        
        var mountainToShoot int = 0 // Which mountain number to shoot?
        var highestMountain int = 0 // Highest mountain
        
        for i := 0; i < 8; i++ {
            // mountainH: represents the height of one mountain, from 9 to 0. Mountain heights are provided from left to right.
            var mountainH int
        
            fmt.Scan(&mountainH)
            
            // check if current mountain is bigger than largest
            if (mountainH > highestMountain) {
                highestMountain = mountainH
                mountainToShoot = i // which Mountain number to shoot?
            }
        }

        
        if (spaceX == mountainToShoot) {
            fmt.Println("FIRE")
        } else {
            fmt.Println("HOLD")
        }
    }
}
