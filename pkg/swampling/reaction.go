package swampling

// Reaction: actions that can be undertaken in response to messages

type reactionVisitor interface {
	visitNoOp(noOpReaction)
	visitSayHello(sayHelloReaction)
}

type reaction interface {
	visit(v reactionVisitor)
}

// Instances of reaction

type noOpReaction struct{}

func (c noOpReaction) visit(v reactionVisitor) { v.visitNoOp(c) }

type sayHelloReaction struct{}

func (c sayHelloReaction) visit(v reactionVisitor) { v.visitSayHello(c) }

// Visitors

type reactionNameVisitor struct{ name string }

func (v *reactionNameVisitor) visitNoOp(r noOpReaction) {
	v.name = "noOpReaction"
}

func (v *reactionNameVisitor) visitSayHello(r sayHelloReaction) {
	v.name = "sayHelloReaction"
}

func ReactionName(r reaction) string {
	visitor := reactionNameVisitor{}
	r.visit(&visitor)
	return visitor.name
}
