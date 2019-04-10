package swampling

import "testing"

func TestReactionName(t *testing.T) {
	tables := []struct {
		r reaction
		n string
	}{
		{noOpReaction{}, "noOpReaction"},
		{sayHelloReaction{}, "sayHelloReaction"},
	}

	for _, table := range tables {
		name := ReactionName(table.r)
		if name != table.n {
			t.Errorf("ReactionName(%v) incorrect,\ngot:  %v,\nwant: %v.",
				table.r, name, table.n)
		}
	}
}
