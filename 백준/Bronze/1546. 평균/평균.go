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
  
  scores := GetManipulatedScores(parts)
  avg := GetAverage(scores)
  
  result := FormatFloat(avg)
  fmt.Println(result)
}

func GetManipulatedScores(parts []string) []float64 {
  max := getMaxScore(parts)

  newScores := make([]float64, len(parts))

  for i, p := range parts {
    n, _ := strconv.Atoi(p)
    
    newScores[i] = manipulateScore(max, n)
  }
  
  return newScores
}

func getMaxScore(parts []string) int {
  max := 0
  for _, p := range parts {
    n, _ := strconv.Atoi(p)
    
    if max < n {
      max = n
    }
  }
  
  return max
}

func manipulateScore(max int, org int) float64 {
  return (float64(org) / float64(max)) * 100
}

func GetAverage(scores []float64) float64 {
  sum := 0.0
  for _, s := range scores {
    sum += s
  }
  
  return float64(sum) / float64(len(scores))
}


func FormatFloat(f float64) string {
  s := fmt.Sprintf("%.6f", f)
  
  s = strings.TrimRight(s, "0")
  
  if strings.HasSuffix(s, ".") {
    s += "0"
  }
  
  return s
}
