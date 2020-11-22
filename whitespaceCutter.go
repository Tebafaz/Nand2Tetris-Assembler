package main

import(
    "strings"
)


// Trims all comments, empty lines and whitespaces.
// Also trims label declaration
// Input: one line
// Outputs: trimmed string, (is it empty?)
func trimmer(line string) (string, bool){
    index := 0
    if strings.Contains(line, "//"){
        index = strings.Index(line, "//")
        line = line[:index]
    }
    for strings.Contains(line, " "){
        index = strings.Index(line, " ")
        line = line[:index] + line[index + 1:]
    }
    if strings.Contains(line, "(") {
        insertIntoL(line)
        return "", true
    }
    if line == ""{
        return line, true
    }
    return line, false
}