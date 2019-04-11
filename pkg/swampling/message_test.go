package swampling

import "testing"

func TestMessageReaction(t *testing.T) {
	tables := []struct {
		message  Message
		reaction reaction
	}{
		{
			Message{From: "lpil", Text: "Hello"},
			noOpReaction{},
		},

		{
			Message{From: "lpil", Text: Nick + " Hello"},
			sayHelloReaction{to: "lpil"},
		},

		{
			Message{From: "lpil", Text: "@" + Nick + " Hello"},
			sayHelloReaction{to: "lpil"},
		},

		{
			Message{From: "jane", Text: Nick + " Hello"},
			sayHelloReaction{to: "jane"},
		},

		{
			Message{From: "jane", Text: Nick + "    Hello"},
			sayHelloReaction{to: "jane"},
		},

		{
			Message{From: "jane", Text: Nick + "    heLLo"},
			sayHelloReaction{to: "jane"},
		},

		{
			Message{From: "jane", Text: Nick + "    hellowekfjwf"},
			noOpReaction{},
		},

		{
			Message{From: "jane", Text: Nick + " Hi"},
			sayHelloReaction{to: "jane"},
		},

		{
			Message{From: "jane", Text: Nick + " Hi!"},
			sayHelloReaction{to: "jane"},
		},

		{
			Message{From: "jane", Text: Nick + " Hey"},
			sayHelloReaction{to: "jane"},
		},

		{
			Message{From: "lpil", Text: Nick + " huh?"},
			noOpReaction{},
		},
	}

	for _, table := range tables {
		reaction := table.message.Reaction()
		if reaction != table.reaction {
			t.Errorf("React(%+v)\n\ngot:  %T%+v,\nwant: %T%+v.",
				table.message,
				reaction,
				reaction,
				table.reaction,
				table.reaction)
		}
	}
}
