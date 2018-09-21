package perspective

import (
	"fmt"
	"encoding/json"
)

type AnalyzeCommentRequest struct {
	CommentText *TextEntry `json:"comment"`
	Languages []string `json:"languages"`
	RequestedAttributes RequestedAttributes `json:"requestedAttributes"`
}

type TextEntry struct {
	Text string `json:"text"`
	TextType string `json:"type,omitempty"`
}

type RequestedAttributes map[string]*Attribute

type Attribute struct {
	ScoreType string `json:"scoreType,omitempty"`
	ScoreThreshold string `json:"scoreThreshold,omitempty"`
}

type AnalyzeCommentResponse struct {
	AttributeScores AttributeScores `json:"attributeScores"`
}

type AttributeScores map[string]*Scores

type Scores struct {
	SpanScores *SpanScores `json:"spanScores"`
	SummaryScore *Score `json:"summaryScore"`
}

type SpanScores []SpanScore

type SpanScore struct {
	Begin int64 `json:"begin"`
	End int64 `json:"end"`
	Score Score `json:"score"`
}

type Score struct {
	Value float32 `json:"value"`
	ScoreType string `json:"type"`
}

func NewAnalyzeCommentRequest(commentText string) *AnalyzeCommentRequest {
	analyzeCommentRequest := AnalyzeCommentRequest{
		CommentText: buildTextEntry(commentText),
		Languages: getDefaultLanguages(),
		RequestedAttributes: getDefaultRequestedAttributes(),
	}
	return &analyzeCommentRequest
}

func printAnalyzeCommentRequestJSON(acr *AnalyzeCommentRequest) {
	marshalledJSON, _ := json.Marshal(acr);
	fmt.Println(string(marshalledJSON))
}

func getDefaultRequestedAttributes() RequestedAttributes {
	requestedAttributes := RequestedAttributes{}
	for _, attributeName := range(DEFAULT_ATTRIBUTES) {
		newAttribute := Attribute{}
		requestedAttributes[attributeName] = &newAttribute
	}
	return requestedAttributes
}

func getDefaultLanguages() []string {
	return []string{
		"en",
	}
}

func buildTextEntry(commentText string) *TextEntry {
	return &TextEntry{
		Text: commentText,
	}
}