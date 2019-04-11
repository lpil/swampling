package swampling

// Reaction: actions that can be undertaken in response to messages

type reaction interface {
	exec(context)
}

// noOpReaction does nothing

type noOpReaction struct{}

func (r noOpReaction) exec(ctx context) {
	// Nothing here...
}

// noOpReaction says hello back

type sayHelloReaction struct {
	to string
}

func (r sayHelloReaction) exec(ctx context) {
	ctx.sendResponse("Hello " + r.to + "!")
}
