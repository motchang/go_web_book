package trace

import (
	"fmt"
	"io"
)

// Tracer 型は Trace というメソッドを1つだけ含むインターフェイス
type Tracer interface {
	// Trace ...interface{} という引数の型は任意の型の引数を何個でも受け取れる
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

func Off() Tracer {
	return &nilTracer{}
}
