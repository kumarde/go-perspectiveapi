package main

import (
	"encoding/json"
	"fmt"
	"github.com/teamnsrg/go-perspectiveapi/perspective"
)

func main() {
	perspectiveClient := perspective.NewAnalyzeCommentClient(PERSPECTIVE_API_KEY)
	analyzeComment := perspectiveClient.AnalyzeComment("toss my salad")
	resp, _ := json.Marshal(analyzeComment)
	fmt.Println(string(resp))
}
