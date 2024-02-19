package replacer

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strconv"
	"strings"

	"github.com/ignite/cli/v28/ignite/pkg/errors"
)

type (
	// functionOpts represent the options for functions.
	functionOpts struct {
		newParams  []param
		body       string
		newLines   []line
		insideCall []call
		appendCode []string
		returnVars []string
	}

	// FunctionOptions configures code generation.
	FunctionOptions func(*functionOpts)

	call struct {
		name  string
		code  string
		index int
	}
	param struct {
		name    string
		varType string
		newLine bool
	}
	line struct {
		code   string
		number uint64
	}
)

// AppendParams add a new param value.
func AppendParams(name, varType string, newLine bool) FunctionOptions {
	return func(c *functionOpts) {
		c.newParams = append(c.newParams, param{
			name:    name,
			varType: varType,
			newLine: newLine,
		})
	}
}

// ReplaceBody replace all body of the function, the method will replace first and apply the other options after.
func ReplaceBody(body string) FunctionOptions {
	return func(c *functionOpts) {
		c.body = body
	}
}

// AppendCode append code before the end or the return, if exists, of a function in Go source code content.
func AppendCode(code string) FunctionOptions {
	return func(c *functionOpts) {
		c.appendCode = append(c.appendCode, code)
	}
}

// AppendAtLine append a new code at line.
func AppendAtLine(code string, lineNumber uint64) FunctionOptions {
	return func(c *functionOpts) {
		c.newLines = append(c.newLines, line{
			code:   code,
			number: lineNumber,
		})
	}
}

// InsideCall add code inside another function call. For instances, the method have a parameter a
// call 'New(param1, param2)' and we want to add the param3 the result will be 'New(param1, param2, param3)'.
// Or if we have a struct call Params{Param1: param1} and we want to add the param2 the result will
// be Params{Param1: param1, Param2: param2}.
func InsideCall(callName, code string) FunctionOptions {
	return func(c *functionOpts) {
		c.insideCall = append(c.insideCall, call{
			name: callName,
			code: code,
		})
	}
}

// NewReturn replaces return statements in a Go function with a new return statement.
func NewReturn(returnVars ...string) FunctionOptions {
	return func(c *functionOpts) {
		c.returnVars = append(c.returnVars, returnVars...)
	}
}

func newFunctionOptions() functionOpts {
	return functionOpts{
		newParams:  make([]param, 0),
		body:       "",
		newLines:   make([]line, 0),
		appendCode: make([]string, 0),
		returnVars: make([]string, 0),
	}
}

func ModifyFunction(fileContent, functionName string, functions ...FunctionOptions) (modifiedContent string, err error) {
	// Apply function options.
	opts := newFunctionOptions()
	for _, o := range functions {
		o(&opts)
	}

	fileSet := token.NewFileSet()

	// Parse the Go source code content.
	f, err := parser.ParseFile(fileSet, "", fileContent, parser.ParseComments)
	if err != nil {
		return "", err
	}

	// Parse the content of the new function into an ast.
	var newFunctionBody *ast.BlockStmt
	if opts.body != "" {
		newFuncContent := fmt.Sprintf("package p; func _() { %s }", strings.TrimSpace(opts.body))
		newContent, err := parser.ParseFile(fileSet, "", newFuncContent, parser.ParseComments)
		if err != nil {
			return "", err
		}
		newFunctionBody = newContent.Decls[0].(*ast.FuncDecl).Body
	}

	// Parse the content of the append code an ast.
	appendCode := make([]ast.Stmt, 0)
	for _, codeToInsert := range opts.appendCode {
		insertionExpr, err := parser.ParseExpr(codeToInsert)
		if err != nil {
			return "", err
		}
		appendCode = append(appendCode, &ast.ExprStmt{X: insertionExpr})
	}

	// Parse the content of the return vars into an ast.
	returnStmts := make([]ast.Expr, 0)
	for _, returnVar := range opts.returnVars {
		// Parse the new return var to expression.
		newRetExpr, err := parser.ParseExpr(returnVar)
		if err != nil {
			return "", err
		}
		returnStmts = append(returnStmts, newRetExpr)
	}

	callMap := make(map[string]call)
	for _, call := range opts.insideCall {
		callMap[call.name] = call
	}

	// Parse the Go code to insert.
	var (
		found      bool
		errInspect error
	)
	ast.Inspect(f, func(n ast.Node) bool {
		funcDecl, ok := n.(*ast.FuncDecl)
		if !ok || funcDecl.Name.Name != functionName {
			return true
		}

		for _, param := range opts.newParams {
			funcDecl.Type.Params.List = append(funcDecl.Type.Params.List, &ast.Field{
				Names: []*ast.Ident{ast.NewIdent(param.name)},
				Type:  &ast.Ident{Name: param.varType},
			})
		}

		// Check if the function has the code you want to replace.
		if newFunctionBody != nil {
			funcDecl.Body = newFunctionBody
		}

		ast.Inspect(funcDecl, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			// Check if the call expression matches the function call name
			var c call
			ident, ok := callExpr.Fun.(*ast.Ident)
			if ok {
				c = callMap[ident.Name]
			} else {
				selector, ok := callExpr.Fun.(*ast.SelectorExpr)
				if ok {
					c = callMap[selector.Sel.Name]
				} else {
					return true
				}
			}

			// Construct the new argument to be added
			newArg := &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(c.code),
			}
			switch {
			case c.index == -1:
				// Append the new argument to the end
				callExpr.Args = append(callExpr.Args, newArg)
				found = true
			case c.index >= 0 && c.index <= len(callExpr.Args):
				// Insert the new argument at the specified index
				callExpr.Args = append(callExpr.Args[:c.index], append([]ast.Expr{newArg}, callExpr.Args[c.index:]...)...)
				found = true
			default:
				errInspect = fmt.Errorf("index out of range")
				return false // Stop the inspection, an error occurred
			}
			return true // Continue the inspection for duplicated calls
		})

		// Add the new code at line.
		for _, newLine := range opts.newLines {
			// Check if the function body has enough lines.
			if newLine.number <= uint64(len(funcDecl.Body.List)) {
				// Parse the Go code to insert.
				insertionExpr, err := parser.ParseExpr(newLine.code)
				if err != nil {
					errInspect = err
					return false
				}
				// Insert code at the specified line number.
				funcDecl.Body.List = append(funcDecl.Body.List[:newLine.number-1], append([]ast.Stmt{&ast.ExprStmt{X: insertionExpr}}, funcDecl.Body.List[newLine.number-1:]...)...)
			}
		}

		// Check if there is a return statement in the function.
		if len(funcDecl.Body.List) > 0 {
			// Replace the return statements.
			for _, stmt := range funcDecl.Body.List {
				if retStmt, ok := stmt.(*ast.ReturnStmt); ok && len(returnStmts) > 0 {
					// Remove existing return statements.
					retStmt.Results = nil
					// Add the new return statement.
					retStmt.Results = append(retStmt.Results, returnStmts...)
				}
			}

			lastStmt := funcDecl.Body.List[len(funcDecl.Body.List)-1]
			switch lastStmt.(type) {
			case *ast.ReturnStmt:
				// If there is a return, insert before it.
				appendCode = append(appendCode, lastStmt)
				funcDecl.Body.List = append(funcDecl.Body.List[:len(funcDecl.Body.List)-1], appendCode...)
			default:
				// If there is no return, insert at the end of the function body.
				funcDecl.Body.List = append(funcDecl.Body.List, appendCode...)
			}
		} else {
			// If there are no statements in the function body, insert at the end of the function body.
			funcDecl.Body.List = append(funcDecl.Body.List, appendCode...)
		}

		found = true
		return false
	})
	if errInspect != nil {
		return "", errInspect
	}
	if !found {
		return "", errors.Errorf("function %s not found in file content", functionName)
	}

	// Format the modified AST.
	var buf bytes.Buffer
	if err := format.Node(&buf, fileSet, f); err != nil {
		return "", err
	}

	// Return the modified content.
	return buf.String(), nil
}