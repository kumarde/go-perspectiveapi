package main

import (
	"encoding/json"
	"fmt"
	"github.com/teamnsrg/go-perspectiveapi/perspective"
)

func main() {
	perspectiveClient := perspective.NewAnalyzeCommentClient(PERSPECTIVE_API_KEY)
	analyzeComment := perspectiveClient.AnalyzeComment("I’m starting to feel pretty nervous about KHIII, and what will be happening in the fandom. Please prepare yourselves to not love every little part of it, be okay with that, and be civil to each other. pls.")
	resp, _ := json.Marshal(analyzeComment)
	fmt.Println(string(resp))
}
