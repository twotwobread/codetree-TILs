package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)


func main() {
  scanner := bufio.NewScanner(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()
  
  scanner.Scan()
  var n int
  fmt.Sscanf(scanner.Text(), "%d", &n)
  
  scanner.Scan()
  fields := strings.Fields(scanner.Text())
  
  sizes := make([]int, len(fields))
  for i, field := range fields {
    sizes[i], _ = strconv.Atoi(field)
  }
  
  scanner.Scan()
  var t, p int
  fmt.Sscanf(scanner.Text(), "%d %d", &t, &p)
  
  maxShirtBundle := GetShirtBundle(sizes, t)
  bundleOfPen, piecesOfPen := GetPenCount(n, p)
  
  fmt.Fprintln(writer, maxShirtBundle)
  fmt.Fprintf(writer, "%d %d\n", bundleOfPen, piecesOfPen)
}

func GetShirtBundle(sizes []int, t int) int {
  totalBundles := 0
  
  for _, size := range sizes {
    bundles := size / t
    if size % t > 0 {
      bundles++
    }
    
    totalBundles += bundles
  }
  
  return totalBundles
}

func GetPenCount(n int, p int) (int, int) {
  bundles := n / p
  pieces := n % p
  
  return bundles, pieces
}