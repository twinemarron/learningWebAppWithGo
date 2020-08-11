package trace

import (
	"fmt"
	"io"
)

// Tracer はコード内でのできことクォ記録できるオブジェクトを表すインターフェースです。
type Tracer interface {
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

// 何も処理を行わない Trace メソッドを定義
func (t *nilTracer) Trace(a ...interface{}) {}

// Off は Trace メソッドの呼び出しを無視する Tracer を返します。
func Off() Tracer {
	return &nilTracer{}
}
