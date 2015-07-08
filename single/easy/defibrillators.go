package main

import "fmt"
import "os"
import "bufio"
import "math"
import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
 
type defib struct {
    ID int64
    Name string
    Address string
    Phone string
    LON float64
    LAT float64
}

func GetInfo(input string) (info defib) {
	result := strings.Split(input, ";")
	
	info.ID, _ = strconv.ParseInt(result[0], 0, 64)
	info.Name = result[1]
	info.Address = result[2]
	info.Phone = result[3]
	
	// replace comma with period
	result[4] = strings.Replace(result[4], ",", ".", -1)
	result[5] = strings.Replace(result[5], ",", ".", -1)
	
	info.LON, _ = strconv.ParseFloat(result[4], 64)
	info.LAT, _ = strconv.ParseFloat(result[5], 64)
	
	return
}


func location (LON, LAT float64, info defib) (distance float64) {
      x := (info.LON - LON) * math.Cos((LAT + info.LAT) / 2)
      y := info.LAT - LAT
      distance = math.Sqrt(x * x + y * y) * 6371
      return 
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var LON string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&LON)
    
    var LAT string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&LAT)
    
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    // Fix Longitude and Latitude values
    LON = strings.Replace(LON, ",", ".", -1)
    LAT = strings.Replace(LAT, ",", ".", -1)
    
    LONINT, _ := strconv.ParseFloat(LON, 64)
    LATINT, _ := strconv.ParseFloat(LAT, 64)
    
    
    //var lines []string
    
    // Variable to hold closest location
    var closeLocation float64 = 10000
    
    // Variable to hold name of closest
    var Name string
    
    // If only single value, use it
    // otherwise figure out which is the closest
    if (N == 1) {
        scanner.Scan()
        ucl := GetInfo(scanner.Text())
        Name = ucl.Name
    } else {
        for i := 0; i < N; i++ {
            scanner.Scan()
            ucl := GetInfo(scanner.Text())
            if (location(LONINT, LATINT, ucl) < closeLocation) {
                closeLocation = location(LONINT, LATINT, ucl)
                Name = ucl.Name
            }
        }
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    fmt.Println(Name)
    //fmt.Println(LAT)
}
