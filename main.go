package main

import (
	"github.com/teamnsrg/go-perspectiveapi/perspective"
)

func main() {
	perspectiveClient := perspective.NewAnalyzeCommentClient(PERSPECTIVE_API_KEY)
	perspectiveClient.AnalyzeComment("I hate you.")
}