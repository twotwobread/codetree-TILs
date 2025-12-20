package main
import (
  "fmt"
  "bufio"
  "os"
)

const MAX = 1000000

func main() {
  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()
  
  var n int
  fmt.Fscanf(reader, "%d\n", &n)
  
  numbers := make([]bool, MAX * 2 + 1)
  for i := 0; i < n; i++ {
    var num int
    fmt.Fscanf(reader, "%d\n", &num)
    
    numbers[num + MAX] = true
  }
  
  for i := 0; i <= MAX*2; i++ {
    if numbers[i] {
      fmt.Fprintln(writer, i - MAX)
    }
  }
}