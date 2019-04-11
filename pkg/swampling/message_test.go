package swampling

import "testing"

func TestMessageReaction(t *testing.T) {
	tables := []struct {
		message  Message
		reaction reaction
	}{
		{
			Message{from: "lpil", text: "Hello"},
			noOpReaction{},
		},

		{
			Message{from: "lpil", text: Nick + " Hello"},
			sayHelloReaction{to: "lpil"},
		},

		{
			Message{from: "lpil", text: "@" + Nick + " Hello"},
			sayHelloReaction{to: "lpil"},
		},

		{
			Message{from: "jane", text: Nick + " Hello"},
			sayHelloReaction{to: "jane"},
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
