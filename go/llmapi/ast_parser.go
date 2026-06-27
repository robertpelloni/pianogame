package llmapi

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type Definition struct {
	Name string `json:"name"`
	Kind string `json:"kind"` // "Function", "Struct", "Interface", etc.
}

type Dependency struct {
	Path string `json:"path"`
}

type FileSyntaxMap struct {
	FilePath     string       `json:"file_path"`
	Definitions  []Definition `json:"definitions"`
	Dependencies []Dependency `json:"dependencies"`
}

type AstParser struct{}

func NewAstParser() *AstParser {
	return &AstParser{}
}

// ParseFile synchronously parses a target source file and returns a structured map of its internal tokens.
func (p *AstParser) ParseFile(filePath string) (*FileSyntaxMap, error) {
	fset := token.NewFileSet()

	// Parse the file and create the AST
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file %s: %w", filePath, err)
	}

	var definitions []Definition
	var dependencies []Dependency

	// Traverse the AST
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			definitions = append(definitions, Definition{
				Name: x.Name.Name,
				Kind: "Function",
			})
		case *ast.TypeSpec:
			kind := "Type"
			switch x.Type.(type) {
			case *ast.StructType:
				kind = "Struct"
			case *ast.InterfaceType:
				kind = "Interface"
			}
			definitions = append(definitions, Definition{
				Name: x.Name.Name,
				Kind: kind,
			})
		case *ast.ImportSpec:
			if x.Path != nil {
				dependencies = append(dependencies, Dependency{
					Path: x.Path.Value, // the value includes the quotes
				})
			}
		}
		return true // continue traversal
	})

	return &FileSyntaxMap{
		FilePath:     filePath,
		Definitions:  definitions,
		Dependencies: dependencies,
	}, nil
}

// ExtractDefinitions extracts high-level structural definitions.
func (p *AstParser) ExtractDefinitions(fileMap *FileSyntaxMap) []Definition {
	if fileMap == nil {
		return nil
	}
	return fileMap.Definitions
}

// MapDependencies identifies local imports/includes to build a directed acyclic graph.
func (p *AstParser) MapDependencies(fileMap *FileSyntaxMap) []Dependency {
	if fileMap == nil {
		return nil
	}
	return fileMap.Dependencies
}
