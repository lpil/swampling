package swampling

import (
	"fmt"
	"reflect"
	"testing"
)

type mockContext struct {
	performed []string
}

func (ctx *mockContext) sendResponse(msg string) {
	fmt.Println("sending response")
	ctx.performed = append(ctx.performed, "sendResponse:"+msg)
}

func makeTestContext() mockContext {
	return mockContext{
		performed: make([]string, 0),
	}
}

func TestSayHelloReactionExec(t *testing.T) {
	tables := []struct {
		reaction reaction
		expected []string
	}{
		{
			reaction: sayHelloReaction{to: "Louis"},
			expected: []string{"sendResponse:Hello Louis!"},
		},

		{
			reaction: sayHelloReaction{to: "Emma"},
			expected: []string{"sendResponse:Hello Emma!"},
		},
	}

	for _, table := range tables {
		ctx := makeTestContext()
		table.reaction.exec(&ctx)

		if !reflect.DeepEqual(ctx.performed, table.expected) {
			t.Errorf("%T%+v.exec()\n\nperformed: %v,\nexpected:  %v.",
				table.reaction,
				table.reaction,
				ctx.performed,
				table.expected)
		}
	}
}
