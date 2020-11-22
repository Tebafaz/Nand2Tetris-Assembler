package main

import (
	"regexp"
    "strings"
    "fmt"
    "strconv"
)

// Translates variable names into address in the RAM
// Input: line of string
// Output: number that corresponds to variable's name in the RAM
func transition(line string) int{
        line = line[1:]
        if number, err := strconv.Atoi(line); err == nil {
            return number
        }
        if _, found := lookupTable[line]; found {
            return lookupTable[line]
        }
        lookupTable[line] = globalAddress
        globalAddress++
        return (globalAddress - 1)
}


// Translates A instructions into byte code
// Input: line of string command
// Output: formatted string that represent byte code
func translateAInstruct(line string)string{
    number := transition(line)
    return fmt.Sprintf("%.16b", number)
}

// Regex that should split C instructions
var regexC = regexp.MustCompile(`[=;]`)

// Translates C instructions into byte code
// Input: line of string command
// Output: string that represents byte code
func translateCInstruct(line string)string{
    instr := regexC.Split(line, -1)
    binLine := "111"

    //dest=comp;jump - this statements is used very rarely
    if len(instr) == 3 {
        binLine += comp[instr[1]]
        binLine += dest[instr[0]]
        binLine += jump[instr[2]]
        return binLine
    }
    // dest=comp - jump should be 000
    if strings.Contains(line, "="){
        binLine += comp[instr[1]]
        binLine += dest[instr[0]]
        return binLine + "000"
    }
    // comp;jump - dest should be 000
    if strings.Contains(line, ";"){
        binLine += comp[instr[0]]
        binLine += "000"
        binLine += jump[instr[1]]
        return binLine
    }

    return binLine
}