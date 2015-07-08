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

type mimeType struct {
    extension string
    mime string
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // N: Number of elements which make up the association table.
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    // Create slice of file types
    fileTypes := make([]mimeType, N)
    
    // Q: Number Q of file names to be analyzed.
    var Q int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&Q)
    
    // MIME types
    for i := 0; i < N; i++ {
        // Pull out information into a struct
        var checkType mimeType
        scanner.Scan()
        fmt.Sscan(scanner.Text(),&checkType.extension, &checkType.mime)
        
        // Append it to list
        fileTypes[i] = checkType
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
            end = strings.ToLower(end)
         
            // Check our list of acceptable extensions
            for j := 0; j < N; j++ {
                if (end == strings.ToLower(fileTypes[j].extension)) {
                    result = fileTypes[j].mime
                }
            }
        }
        fmt.Println(result)
    }
}
