package swampling

import "strings"

type Message struct {
	text string
	from string
}

// Find an appropriate reaction to a message

func (m Message) Reaction() reaction {
	switch {
	case strings.HasPrefix(m.text, Nick):
		return sayHelloReaction{to: m.from}

	case strings.HasPrefix(m.text, "@"+Nick):
		return sayHelloReaction{to: m.from}

	default:
		return noOpReaction{}
	}
}
