package perspective

var DEFAULT_ATTRIBUTES = []string{
	"TOXICITY",
	"SEVERE_TOXICITY",
	"IDENTITY_ATTACK",
	"INSULT",
	"PROFANITY",
	"SEXUALLY_EXPLICIT",
	"THREAT",
	"ATTACK_ON_AUTHOR",
}

const ANALYZE_COMMENT_API_BASE = "https://commentanalyzer.googleapis.com/v1alpha1/comments:analyze"
