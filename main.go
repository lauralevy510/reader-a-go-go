package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main() {

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("hey girrrl!")
  fmt.Println("----------------------")

  for {
    fmt.Print("->")
    text, _ := reader.ReadString('\n')
    //convert CRLF to LF
    text = strings.Replace(text, "\n","",-1)

    if strings.Compare("hi",text) == 0 {
      fmt.Println("hello, yourself")
      break;
    }
    if strings.Compare("xx",text) ==0 {
      fmt.Println("goodbye")
    }

  }

}
