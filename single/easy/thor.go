package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 * ---
 * Hint: You can use the debug stream to print initialTX and initialTY, if Thor seems not follow your orders.
 **/

func main() {
    // lightX: the X position of the light of power
    // lightY: the Y position of the light of power
    // initialTX: Thor's starting X position
    // initialTY: Thor's starting Y position
    var lightX, lightY, initialTX, initialTY int
    fmt.Scan(&lightX, &lightY, &initialTX, &initialTY)
    
    thorX := initialTX
    thorY := initialTY
    var goTo string
    
    for {
        
        var directionX string = ""
        var directionY string = ""
        
        if (thorX > lightX) {
            directionX = "W"
            thorX -= 1
        } else if (thorX < lightX) {
            directionX = "E"
            thorX += 1
        }
        
        if (thorY > lightY) {
            directionY = "N"
            thorY -= 1
        } else if (thorY < lightY) {
            directionY = "S"
            thorY += 1
        }
        
        goTo = fmt.Sprint(directionY + directionX)
        
        fmt.Println(goTo)
    }
}
