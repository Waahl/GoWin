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
      fmt.Printf("\t [*] Failed to delete apps, Reason: %v \n", err)
    }

    err = cmd.Wait()
    fmt.Printf("\t [*] Finished deleting apps with error: %v \n", err)
}

func runExes() {
    var exitKey string

    files, err := ioutil.ReadDir(".")

    if err != nil {
        fmt.Printf("\t [*] Error reading current directory occured.\n")
        fmt.Printf("\t [*] Press Enter to exit ...")
        fmt.Scanln(&exitKey)
        os.Exit(1)
    }

    for _, file := range files {
        cmd := exec.Command(file.Name())
        err := cmd.Start()

        if err != nil {
            fmt.Printf("\t [*] Error running program, skipping file: %s, Reason: %v \n", file.Name(), err)
            continue
        }

        fmt.Printf("\t [*] Waiting for install to finish. \n")
        err = cmd.Wait()
        fmt.Printf("\t [*] Status: %s \n", err)
    }
}

func checkDir() {
  var exitKey string

  if curr_dir, err := os.Getwd(); err != nil {
      fmt.Printf("\t [*] Unexpected error occured.")
      fmt.Printf("\t [*] Press Enter to exit ...")
      fmt.Scanln(&exitKey)
      os.Exit(1)
  } else if hostname, err := os.Hostname(); err != nil {
      fmt.Printf("\t [*] Error fetching windows user.")
  } else if curr_dir != "c:\\users\\" + hostname + "\\Downloads" {
      os.Chdir("c:\\users\\" + hostname + "\\Downloads")
  }
}

func printInfo() string {
    infoText := "\t [*] Windows Setup Program\n\t [*] This program runs exes downloaded in the download folder to speed up the process as well as deleting bloatware.\n"
    return infoText
}

func main() {
    var exitKey string
    fmt.Printf(printInfo())

    // Check user platform to ensure a windows platform
    if runtime.GOOS != "windows" {
      fmt.Printf("\t [*] This script was designed for windows.")
      fmt.Printf("\t [*] Press Enter to exit ...")
      fmt.Scanln(&exitKey)
      os.Exit(1);
    }

    // Check if running in administrator mode
    _, err := exec.Command("net", "session").Output()
    if err != nil {
      fmt.Printf("\t [*] Please run the exe in administrator mode.\n")
      fmt.Printf("\t [*] Press Enter to exit ...")
      fmt.Scanln(&exitKey)
      os.Exit(1)
    }

    // Get user authorization to run scripts
    fmt.Printf("\t [+] Do you wish to continue? y/n: ")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {

        switch strings.ToLower(scanner.Text()) {
        case "y":
            fmt.Printf("\t [*] Running scripts. This may take a while.\n")
            checkDir()
            runExes()
            deleteWinApps()
            fmt.Printf("\t [*] Script finished.\n")
        case "n":
            fmt.Printf("\t [*] Program has been cancelled.\n")
            fmt.Printf("\t [*] Press Enter to exit ...")
            fmt.Scanln(&exitKey)
            os.Exit(0)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("reading standard input: \n", err)
        fmt.Printf("\t [*] Press Enter to exit ...")
        fmt.Scanln(&exitKey)
        os.Exit(1)
    }

    os.Exit(0)
}
