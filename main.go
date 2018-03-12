package main

import (
    "strings"
    "bufio"
    "fmt"
    "os"
)

func printInfo() string {
    infoText := " \t [*] Windows Setup Program\n\t [*] This program is written to speed up the process of de-bloating Windows by downloading and running powershell scripts\n\t [*] Note: All powershell scripts are properties of their original author."
    

    return infoText
}

func main() {
    fmt.Println(printInfo())

    // Get user authorization to run scripts
    fmt.Println("\t [+] Do you wish to continue? y/n")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {

        switch strings.ToLower(scanner.Text()) {
        case "y":
            fmt.Println("\t [*] Running scripts. This may take a while")
        case "n":
            fmt.Println("\t [*] Program has been cancelled")
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(os.Stderr, "reading standard input: ", err)
    }
}
