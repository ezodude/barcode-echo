package main

import (
  "bufio"
  "fmt"
  "os"
)

func ScannerLines() {
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    fmt.Println("Scanned barcode", scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
  }
}

func main() {
  ScannerLines()
}