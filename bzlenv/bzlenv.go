// Package bzlenv provides function to create and update a static environment.
package bzlenv

import (
	"github.com/bazelbuild/buildtools/build"
)

// ValueKind describes how a binding was declared.
type ValueKind int

// List of ValueKind values.
const (
	Builtin   ValueKind = iota // language builtin
	Imported                   // declared with load()
	Global                     // declared with assignment on top-level
	Function                   // declared with a def
	Parameter                  // function parameter
	Local                      // local variable, defined with assignment or as a loop variable
)

func (k ValueKind) String() string {
	switch k {
	case Builtin:
		return "builtin"
	case Imported:
		return "imported"
	case Global:
		return "global"
	case Function:
		return "function"
	case Parameter:
		return "parameter"
	case Local:
		return "local"
	default:
		panic(k)
	}
}

// NameInfo represents information about a symbol name.
type NameInfo struct {
	ID         int    // unique identifier
	Name       string // name of the variable (not unique)
	Kind       ValueKind
	Definition build.Expr // node that defines the value
}

type block map[string]NameInfo

// Environment represents a static environment (e.g. information about all available symbols).
type Environment struct {
	Blocks   []block
	Function *build.DefStmt // enclosing function (or nil on top-level)
	nextID   int            // used to create unique identifiers
	Stack    []build.Expr   // parents of the current node
}

// NewEnvironment creates a new empty Environment.
func NewEnvironment() *Environment {
	sc := block{}
	return &Environment{[]block{sc}, nil, 0, []build.Expr{}}
}

func (e *Environment) enterBlock() {
	e.Blocks = append(e.Blocks, block{})
}

func (e *Environment) exitBlock() {
	if len(e.Blocks) < 1 {
		panic("no block to close")
	}
	e.Blocks = e.Blocks[:len(e.Blocks)-1]
}

func (e *Environment) currentBlock() block {
	return e.Blocks[len(e.Blocks)-1]
}

func (sc *block) declare(name string, kind ValueKind, definition build.Expr, id int) {
	(*sc)[name] = NameInfo{
		ID:         id,
		Name:       name,
		Definition: definition,
		Kind:       kind}
}

// Get resolves the name and resolves information about the binding (or nil if it's not defined).
func (e *Environment) Get(name string) *NameInfo {
	for i := len(e.Blocks) - 1; i >= 0; i-- {
		if ret, ok := e.Blocks[i][name]; ok {
			return &ret
		}
	}
	return nil
}

func (e *Environment) declare(name string, kind ValueKind, node build.Expr) {
	sc := e.currentBlock()
	sc.declare(name, kind, node, e.nextID)
	e.nextID++
}

func declareGlobals(stmts []build.Expr, env *Environment) {
	for _, node := range stmts {
		switch node := node.(type) {
		case *build.LoadStmt:
			for _, ident := range node.To {
				env.declare(ident.Name, Imported, ident)
			}
		case *build.BinaryExpr:
			if node.Op == "=" {
				kind := Local
				if env.Function == nil {
					kind = Global
				}
				for _, id := range collectLValues(node.X) {
					env.declare(id.Name, kind, node)
				}
			}
		case *build.DefStmt:
			env.declare(node.Name, Function, node)
		}
	}
}

// collectLValues returns the list of identifiers that are assigned (assuming that node is a valid
// LValue). For example, it returns `a`, `b` and `c` for the input `a, (b, c)`.
func collectLValues(node build.Expr) []*build.Ident {
	var result []*build.Ident
	// TODO: we collect a bit more, e.g.
	//     f(x)[y] = 2
	// should return only f
	build.Walk(node, func(e build.Expr, stk []build.Expr) {
		switch e := e.(type) {
		case *build.Ident:
			result = append(result, e)

		}
	})
	return result
}

func declareParams(fct *build.DefStmt, env *Environment) {
	for _, node := range fct.Params {
		switch node := node.(type) {
		case *build.Ident:
			env.declare(node.Name, Parameter, node)
		case *build.UnaryExpr:
			// either *args or **kwargs
			if ident, ok := node.X.(*build.Ident); ok {
				env.declare(ident.Name, Parameter, node)
			}
		case *build.BinaryExpr:
			// x = value
			if ident, ok := node.X.(*build.Ident); ok {
				env.declare(ident.Name, Parameter, node)
			}
		}
	}
}

func declareLocalVariables(stmts []build.Expr, env *Environment) {
	for _, stmt := range stmts {
		switch node := stmt.(type) {
		case *build.BinaryExpr:
			if node.Op == "=" {
				kind := Local
				if env.Function == nil {
					kind = Global
				}
				for _, id := range collectLValues(node.X) {
					env.declare(id.Name, kind, node)
				}
			}
		case *build.IfStmt:
			declareLocalVariables(node.True, env)
			declareLocalVariables(node.False, env)
		case *build.ForStmt:
			for _, id := range collectLValues(node.Vars) {
				env.declare(id.Name, Local, node)
			}
			declareLocalVariables(node.Body, env)
		}
	}
}

// WalkOnceWithEnvironment calls fct on every child of node, while maintaining the Environment of all available symbols.
func WalkOnceWithEnvironment(node build.Expr, env *Environment, fct func(e *build.Expr, env *Environment)) {
	env.Stack = append(env.Stack, node)
	switch node := node.(type) {
	case *build.File:
		declareGlobals(node.Stmt, env)
		build.WalkOnce(node, func(e *build.Expr) { fct(e, env) })
	case *build.DefStmt:
		env.enterBlock()
		env.Function = node
		declareParams(node, env)
		declareLocalVariables(node.Body, env)
		build.WalkOnce(node, func(e *build.Expr) { fct(e, env) })
		env.Function = nil
		env.exitBlock()
	case *build.Comprehension:
		env.enterBlock()
		for _, clause := range node.Clauses {
			switch clause := clause.(type) {
			case *build.ForClause:
				for _, id := range collectLValues(clause.Vars) {
					env.declare(id.Name, Local, node)
				}
			}
		}
		build.WalkOnce(node, func(e *build.Expr) { fct(e, env) })
		env.exitBlock()
	default:
		build.WalkOnce(node, func(e *build.Expr) { fct(e, env) })
	}
	env.Stack = env.Stack[:len(env.Stack)-1]
}
