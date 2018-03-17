package main

import (
    "io/ioutil"
    "os/exec"
    "runtime"
    "strings"
    "bufio"
    "fmt"
    "os"
)

func deleteWinApps() {
    cmd := exec.Command("powershell", "get-appxpackage | remove-appxpackage")
    err := cmd.Start()

    if err != nil {
      fmt.Println(fmt.Sprintf("\t [*] Failed to delete apps, Reason: %v", err))
    }

    err = cmd.Wait()
    fmt.Println(fmt.Sprintf("\t [*] Finished deleting apps with error: %v", err))
}

func runExes() {
    files, err := ioutil.ReadDir(".")

    if err != nil {
        fmt.Println("\t [*] Error reading current directory occured")
        os.Exit(1)
    }

    for _, file := range files {
        cmd := exec.Command(file.Name())
        err := cmd.Start()

        if err != nil {
            fmt.Println(fmt.Sprintf("\t [*] Error running program, skipping file: %s, Reason: %v", file.Name(), err))
            continue
        }

        fmt.Println("\t [*] Waiting for install to finish")
        err = cmd.Wait()
        fmt.Println(fmt.Sprintf("\t [*] Status: %s", err))
    }
}

func checkDir() {
  if curr_dir, err := os.Getwd(); err != nil {
    fmt.Println("\t [*] Unexpected error occured.")
    os.Exit(1)
  } else if hostname, err := os.Hostname(); err != nil {
    fmt.Println("\t [*] Error fetching windows user.")
  } else if curr_dir != "c:\\users\\" + hostname + "\\Downloads" {
    os.Chdir("c:\\users\\" + hostname + "\\Downloads")
  }
}

func printInfo() string {
    infoText := "\t [*] Windows Setup Program\n\t [*] This program is written to speed up the process of de-bloating Windows by downloading and running powershell scripts\n\t [*] Note: All powershell scripts are properties of their original author."
    return infoText
}

func main() {
    fmt.Println(printInfo())

    // Check user platform to ensure a windows platform
    if runtime.GOOS != "windows" {
      fmt.Println("\t [*] This script was designed for windows.")
      os.Exit(1);
    }

    // Get user authorization to run scripts
    fmt.Println("\t [+] Do you wish to continue? y/n")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {

        switch strings.ToLower(scanner.Text()) {
        case "y":
            fmt.Println("\t [*] Running scripts. This may take a while")
            checkDir()
            runExes()
            deleteWinApps()
            fmt.Printf(fmt.Sprintf("\t [*] Script finished."))
        case "n":
            fmt.Println("\t [*] Program has been cancelled")
            os.Exit(0)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(os.Stderr, "reading standard input: ", err)
        os.Exit(1)
    }

    os.Exit(0)
}
