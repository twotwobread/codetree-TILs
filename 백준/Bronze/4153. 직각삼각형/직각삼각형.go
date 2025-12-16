package main

import (
  "bufio"
  "fmt"
  "os"
)


func main() {
  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()
  
  var x, y, z int
  
  for {
    fmt.Fscanf(reader, "%d %d %d\n", &x, &y, &z)
    if x == 0 && y == 0 && z == 0 {
      break
    }
    
    if isRightTriangle(x, y, z) {
      fmt.Fprintln(writer, "right")
    } else {
      fmt.Fprintln(writer, "wrong")
    }
  }
}

func isRightTriangle(x int, y int, z int) bool {
  x2, y2, z2 := x*x, y*y, z*z
  
  return x2+y2 == z2 || x2+z2 == y2 || y2+z2 == x2
}