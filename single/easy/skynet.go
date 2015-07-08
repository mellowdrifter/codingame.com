package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // road: the length of the road before the gap.
    var road int
    fmt.Scan(&road)
    
    // gap: the length of the gap.
    var gap int
    fmt.Scan(&gap)
    
    // How fast do we need to go?
    var speedNeeded = gap + 1
    
    // When does the platform start?
    var platformStart = road + gap
    
    // platform: the length of the landing platform.
    var platform int
    fmt.Scan(&platform)
    
    for {
        // speed: the motorbike's speed.
        var speed int
        fmt.Scan(&speed)
        
        // coordX: the position on the road of the motorbike.
        var coordX int
        fmt.Scan(&coordX)
        
        if (coordX == road - 1) {
            fmt.Println("JUMP")
        } else if (coordX < platformStart && speed < speedNeeded) {
            fmt.Println("SPEED")
        } else if (coordX < platformStart && speed > speedNeeded) {
            // Need to catch cases where starts too fast
            fmt.Println("SLOW")
        } else if (coordX < platformStart) {
            fmt.Println("WAIT")
        } else {
            fmt.Println("SLOW")
        }
    }
}
