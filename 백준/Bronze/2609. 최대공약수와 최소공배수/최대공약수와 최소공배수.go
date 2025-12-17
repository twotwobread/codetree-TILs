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
  parts := strings.Fields(scanner.Text())
  
  x, _ := strconv.Atoi(parts[0])
  y, _ := strconv.Atoi(parts[1])
  
  gcd := GetGreatestCommonDivisorUsingEuclidean(x, y)
  lcm := GetLeastCommonMultiple(x, y, gcd)
  
  fmt.Println(gcd)
  fmt.Println(lcm)
}

func GetGreatestCommonDivisor(x int, y int) int {
  var maxValue int
  if x > y {
    maxValue = x
  } else {
    maxValue = y
  }
  
  gcd := 1
  for i := 1; i <= maxValue; i++ {
    if x % i == 0 && y % i == 0{
      gcd = i
    }
  }
  
  return gcd
}

func GetGreatestCommonDivisorUsingEuclidean(x int, y int) int {
  for y != 0 {
    x, y = y, x % y
  }
  return x
}

func GetLeastCommonMultiple(x int, y int, gcd int) int {
  return (x * y) / gcd
}
