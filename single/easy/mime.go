package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // N: Number of elements which make up the association table.
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    // Q: Number Q of file names to be analyzed.
    var Q int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&Q)
    
    // Create map, AKA hash table
    mimeType := make(map[string]string)
    mimeTypeLower := make(map[string]string)
    
    // MIME types
    for i := 0; i < N; i++ {
        
        // Pull out information
        var ET, MT string
        scanner.Scan()
        fmt.Sscan(scanner.Text(),&ET, &MT)
        
        // Insert into map. Separate map for lowercase extensions
        mimeType[ET] = MT
        mimeTypeLower[strings.ToLower(ET)] = ET
    }
    
    
    for i := 0; i < Q; i++ {
        var fileExtension string
        result := "UNKNOWN"
        scanner.Scan()
        fmt.Sscan(scanner.Text(), &fileExtension)
        
        // Only check file extension if there is actually an extension
        if strings.Contains(fileExtension, ".") {
            split := strings.Split(fileExtension, ".")
            end := split[len(split) - 1]
         
            // Check if extension exists in map
            key, ok := mimeTypeLower[strings.ToLower(end)]
            if (ok) {
                // Get the original extension size
                result, _ = mimeType[key]
            }
        }
        fmt.Println(result)
    }
}
