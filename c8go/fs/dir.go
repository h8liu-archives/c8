package fs

type Node interface{}

type Dir struct {
	subs map[string]Node
}
