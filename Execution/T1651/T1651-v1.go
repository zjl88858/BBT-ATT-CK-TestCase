//usr/bin/env go run "$0" "$@"; exit

package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "io/ioutil"
    "path/filepath"
)

func main() {
    fmt.Println("BangBangTuan ATT&CK TestCase - T1651 - Cloud Administration Command")

    homeDir, err := os.UserHomeDir()
    if err != nil {
        log.Printf("[ERROR] Can't get user home directory: %v\n", err)
    }

    credentialsFile := filepath.Join(homeDir, ".aws", "credentials")
    credentialsDir := filepath.Dir(credentialsFile)
    dirCreated := false

    _, err = os.Stat(credentialsDir)
    if os.IsNotExist(err) {
        err = os.MkdirAll(credentialsDir, 0755)
        if err != nil {
            log.Printf("[ERROR] Can't creat .aws directory: %v\n", err)
        }
        dirCreated = true
    }

    cmd := exec.Command("aws", "s3", "ls", "s3://fake-bucket")
    fmt.Println("[INFO] Exec command:", cmd.String())

    fmt.Println("[INFO] CASE:Reading AWS credentials file")
    fileCreated := false

    if _, err := os.Stat(credentialsFile); os.IsNotExist(err) {
        fmt.Println("[INFO] Credentials file not found, creating a fake one to continue test progress:", credentialsFile)
        err := ioutil.WriteFile(credentialsFile, []byte("[default]\naws_access_key_id=FAKE_ACCESS_KEY\naws_secret_access_key=FAKE_SECRET_KEY"), 0600)
        if err != nil {
            log.Printf("[ERROR] Can't creat fake credentials file: %v\n", err)
        } else {
            fileCreated = true
        }
    }

    _, err = ioutil.ReadFile(credentialsFile)
    if err != nil {
        log.Printf("[ERROR] can't reading credentials file: %v\n", err)
    } else {
        fmt.Println("[INFO] Credentials file readed:", credentialsFile)
    }

    fmt.Println("[INFO] CASE:Reading AWS APIs ENV")
    accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
    secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

    if accessKeyID != "" && secretAccessKey != "" {
        fmt.Println("[INFO] Done")
    } else {
        fmt.Println("[INFO] AWS APIs ENV not found (not a error!)")
    }

    fmt.Println("[INFO] Cleaning test case temp files")

    if fileCreated {
        fmt.Printf("[INFO] Delete fake credentials file: %s\n", credentialsFile)
        err := os.Remove(credentialsFile)
        if err != nil {
            log.Printf("[ERROR] Can't Delete: %v\n", err)
        } else {
            fmt.Println("[INFO] Done")
        }
    }

    if dirCreated {
        fmt.Printf("[INFO] Delete .aws directory: %s\n", credentialsDir)
        err := os.RemoveAll(credentialsDir)
        if err != nil {
            log.Printf("[ERROR] Can't Delete: %v\n", err)
        } else {
            fmt.Println("[INFO] Done")
        }
    }

    fmt.Println("[INFO] Cleaning complete. Press any key (except power key and your door key) to exit.")

	fmt.Scanln()
}
