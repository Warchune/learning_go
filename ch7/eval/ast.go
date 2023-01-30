package eval

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

// определяет переменную, например х
type Var string

// представляет собой числовую константу. например 3.141
type literal float64

// представляет собой выражение с унарным оператором, например -х.
type unary struct {
	op rune
	x  Expr
}

// представляет собой выражение с унарным оператором, например x+y.
type binary struct {
	op   rune
	x, y Expr
}

// представляет собой выражение с унарным оператором, например sin(x)
type call struct {
	fn   string // одно из "pow", "sin", "aqrt"
	args []Expr
}
