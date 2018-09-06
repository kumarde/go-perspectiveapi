package perspective

import (
	"github.com/dghubble/sling"
	log "github.com/sirupsen/logrus"
)



type Client struct {
	sling *sling.Sling
	apiKey string
}

func (c *Client) AnalyzeComment(commentText string) *AnalyzeCommentResponse {
	analyzeCommentRequest := NewAnalyzeCommentRequest(commentText)
	analyzeCommentResponse := AnalyzeCommentResponse{}
	_, err := c.sling.New().Post("").BodyJSON(analyzeCommentRequest).ReceiveSuccess(&analyzeCommentResponse)
	if err != nil {
		log.Error("there was an error : ", err)
	}
	return &analyzeCommentResponse
}

func NewAnalyzeCommentClient(apiKey string) *Client {
	analyzeCommentAPIPath := ANALYZE_COMMENT_API_BASE + "?key=" + apiKey
	base := sling.New().Client(nil).Base(analyzeCommentAPIPath);
	return &Client{
		sling: base,
		apiKey: apiKey,
	}
}