package main
import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  
  for scanner.Scan() {
    s := scanner.Text()
    
    if s == "0" {
      break
    }
    
    if isPalindrome(s) {
      fmt.Println("yes")
    } else {
      fmt.Println("no")
    }
  }
}

func isPalindrome(s string) bool {
  size := len(s)
  for x := 0; x < size / 2; x++ {
    if s[x] != s[size-1-x] {
      return false
    }
  }
  
  return true
}
