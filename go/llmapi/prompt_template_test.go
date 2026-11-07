package llmapi

import (
	"reflect"
	"sort"
	"testing"
)

func TestExtractTemplateVariables(t *testing.T) {
	text := "Hello {{name}}, welcome to {{language}}! Also hello again {{name}}."
	vars := ExtractTemplateVariables(text)
	expected := []string{"name", "language"}
	if !reflect.DeepEqual(vars, expected) {
		t.Fatalf("expected %v, got %v", expected, vars)
	}
}

func TestRenderPromptTemplate(t *testing.T) {
	def := PromptTemplateDefinition{
		Name:    "unit_test_template",
		Version: "1.0.0",
		Parameters: []TemplateParameter{
			{Name: "framework", Type: "string", Required: true},
			{Name: "target", Type: "string", Required: true},
			{Name: "verbosity", Type: "string", Required: false, DefaultValue: "normal"},
		},
		Messages: []TemplateMessage{
			{Role: "system", Content: "Framework: {{framework}}, Verbosity: {{verbosity}}"},
			{Role: "user", Content: "Generate tests for {{target}}"},
		},
	}

	// 1. Valid render with default value fallback
	res := RenderPromptTemplate(def, map[string]string{
		"framework": "Go testing",
		"target":    "PromptTemplate",
	})
	if len(res.MissingParams) != 0 {
		t.Fatalf("expected no missing params, got %v", res.MissingParams)
	}
	if res.Messages[0].Content != "Framework: Go testing, Verbosity: normal" {
		t.Fatalf("unexpected rendered content: %s", res.Messages[0].Content)
	}
	if res.Messages[1].Content != "Generate tests for PromptTemplate" {
		t.Fatalf("unexpected rendered content: %s", res.Messages[1].Content)
	}

	// 2. Missing required parameters
	failed := RenderPromptTemplate(def, map[string]string{
		"verbosity": "verbose",
	})
	sort.Strings(failed.MissingParams)
	expectedMissing := []string{"framework", "target"}
	if !reflect.DeepEqual(failed.MissingParams, expectedMissing) {
		t.Fatalf("expected %v missing params, got %v", expectedMissing, failed.MissingParams)
	}
}
