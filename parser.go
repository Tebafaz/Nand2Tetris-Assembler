package main

import(
    "strings"
    "fmt"
)

// Finds a Label declaration if there is any.
// Input: one line of string
// Outputs: string of label's name
func findLabel(line string)string{
        
        start := strings.Index(line, "(") + 1
        end := len(line) - 1
        label := line[start:end]

        return label
}


// Inserts  into lookup table
func insertIntoL(line string) string{

    label := findLabel(line)

    if _, found := lookupTable[label]; !found{
        lookupTable[label] = globalCounter
        return fmt.Sprint(lookupTable[label])
    }
    return fmt.Sprint(lookupTable[label])
}