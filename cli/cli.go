package cli

import (
	"blom/analyzer"
	"blom/compiler"
	"blom/interpreter"
	"blom/lexer"
	"blom/parser"
	"blom/tokens"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gookit/goutil/dump"
)

func Run(args []string) {
	var emitTokens bool
	var emitAst bool
	var emitSse bool
	var inputFile string

	for _, arg := range args {
		switch arg {
		case "-tokens", "--emit-tokens":
			emitTokens = true
		case "-ast", "--emit-ast":
			emitAst = true
		case "-sse", "--emit-sse":
			emitSse = true
		default:
			if strings.HasSuffix(arg, ".blom") {
				inputFile = arg
			}
		}
	}

	if inputFile == "" {
		panic("No input file specified")
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	lex := lexer.New(inputFile, string(content))
	tkns := make([]tokens.Token, 0)

	current := lex.Next()
	for current.Kind != tokens.Eof {
		tkns = append(tkns, *current)
		current = lex.Next()
	}

	if emitTokens {
		dump.Println(tkns)
	}

	parser := parser.New(inputFile)
	ast := parser.AST(inputFile, string(content))

	analyzer := analyzer.New(inputFile, ast)
	analyzer.Analyze()

	if emitAst {
		dump.Println(ast)
	}

	//os.Exit(1)

	_ = interpreter.New(inputFile)

	fmt.Printf("Interpreting %s\n", inputFile)

	startTime := time.Now()
	//inp.Interpret(ast, int64(len(os.Args)-1))
	endTime := time.Since(startTime)

	fmt.Printf("Interpreted %s ran for %s\n", inputFile, endTime)

	fmt.Println()

	fmt.Printf("Compiling %s\n", inputFile)
	compiler := compiler.New(compiler.QBE)
	sse := compiler.Compile(ast)

	if emitSse {
		fmt.Println(sse)
	}

	err = os.WriteFile("out.sse", []byte(sse), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Running %s\n", inputFile)
	startTime = time.Now()
	cmd := exec.Command("sh", "-c", "qbe -o out.s out.sse && cc -O3 out.s -o a.out && ./a.out")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd.Wait()
	endTime = time.Since(startTime)

	fmt.Printf("Compiled %s ran for %s\n", inputFile, endTime)

	os.Exit(cmd.ProcessState.ExitCode())
}
