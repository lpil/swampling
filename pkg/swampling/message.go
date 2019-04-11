package swampling

import (
	"regexp"
	"strings"
)

type Message struct {
	Text string
	From string
}

var (
	helloRegex = regexp.MustCompile("\\A(hello|hey|hi)\\b")
)

// Find an appropriate reaction to a message

func (m Message) Reaction() reaction {
	text := m.Text
	text = strings.TrimPrefix(text, Nick)
	text = strings.TrimPrefix(text, "@"+Nick)
	text = strings.TrimLeft(text, " ")

	if text == m.Text {
		return noOpReaction{}
	}

	lower := strings.ToLower(text)

	switch {
	case helloRegex.MatchString(lower):
		return sayHelloReaction{to: m.From}

	default:
		return noOpReaction{}
	}
}
