package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
)


type Track struct {
  Name    string
  Subjects []string
}

func (t Track) Contains(subject string) bool {
  parts := strings.Fields(subject)
  
  for _, p := range parts {
    for _, s := range t.Subjects {
      if p == s {
        return true
      }
    }
  }
  
  return false
}

var tracks = []Track{
  Track{
    Name: "digital humanities",
    Subjects: []string{
      "social", "history", "language", "literacy",
    },
  },
  Track{
    Name: "public bigdata",
    Subjects: []string {
      "bigdata", "public", "society",
    },
  },
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  
  scanner.Scan()
  subject := scanner.Text()
  
  trackNm := GetTrackName(subject)
  
  fmt.Println(trackNm)
}


func GetTrackName(subject string) string {
  for _, track := range tracks {
    if track.Contains(subject) {
      return track.Name
    }
  }
  return "None"
}