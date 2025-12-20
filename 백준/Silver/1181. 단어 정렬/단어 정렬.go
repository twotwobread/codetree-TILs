package main
import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "sort"
)


func main() {
  scanner := bufio.NewScanner(os.Stdin)
  
  scanner.Scan()
  n, _ := strconv.Atoi(scanner.Text())

  seen := make(map[string]bool)  
  words := []string{}
  
  for i:=0; i<n; i++ {
    scanner.Scan()
    name := scanner.Text()
    
    if !seen[name] {
      seen[name] = true
      words = append(words, name)
    }
  }
  
  //sort.Sort(ByWord(words))
  sort.Slice(words, func(i, j int) bool {
    if len(words[i]) == len(words[j]) {
      return words[i] < words[j]
    }
  
    return len(words[i]) < len(words[j])
  })
  
  for _, w := range words {
    fmt.Println(w)
  }
}