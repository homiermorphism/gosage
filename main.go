package main

import (
  "fmt"
  "log"
  "os/exec"
)

func main() {
  cmd := exec.Command("sage", "test.sage")

  out, err := cmd.CombinedOutput()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(out))
}
