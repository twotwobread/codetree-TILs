package main
import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  
  scanner.Scan()
  _ = scanner.Text()
  
  scanner.Scan()
  parts := strings.Fields(scanner.Text())
  
  result := cntPrimeNumber(parts)
  fmt.Println(result)
}

func cntPrimeNumber(parts []string) int {
  total := 0
  for _, p := range parts {
    num, _ := strconv.Atoi(p)
    
    if isPrimeNumber(num) {
      total++
    } 
  }
  
  return total
}

func isPrimeNumber(num int) bool {
  if num <= 1 {
    return false
  }
  
  for n := 2; n < num; n++ {
    if num % n == 0 {
      return false
    } 
  }
  
  return true
}