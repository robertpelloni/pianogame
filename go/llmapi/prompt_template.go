package llmapi

import (
	"regexp"
	"strings"
)

// TemplateParameter defines an input parameter constraint.
type TemplateParameter struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Required     bool   `json:"required"`
	DefaultValue string `json:"defaultValue,omitempty"`
	Description  string `json:"description,omitempty"`
}

// TemplateMessage represents a single message in a multi-turn prompt.
type TemplateMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// PromptTemplateDefinition defines the entire schema for a template.
type PromptTemplateDefinition struct {
	Name        string              `json:"name"`
	Version     string              `json:"version"`
	Description string              `json:"description,omitempty"`
	Parameters  []TemplateParameter `json:"parameters"`
	Messages    []TemplateMessage   `json:"messages"`
}

// RenderedPromptResult holds the rendered messages and any missing parameters.
type RenderedPromptResult struct {
	Messages      []TemplateMessage `json:"messages"`
	MissingParams []string          `json:"missingParams"`
}

var varRegex = regexp.MustCompile(`\{\{([a-zA-Z0-9_]+)\}\}`)

// ExtractTemplateVariables finds all placeholder names in text.
func ExtractTemplateVariables(text string) []string {
	matches := varRegex.FindAllStringSubmatch(text, -1)
	seen := make(map[string]bool)
	var result []string
	for _, m := range matches {
		if len(m) > 1 && !seen[m[1]] {
			seen[m[1]] = true
			result = append(result, m[1])
		}
	}
	return result
}

// InterpolateTemplate replaces {{var}} with the provided values.
func InterpolateTemplate(text string, values map[string]string) string {
	return varRegex.ReplaceAllStringFunc(text, func(m string) string {
		varName := m[2 : len(m)-2]
		if val, ok := values[varName]; ok {
			return val
		}
		return m
	})
}

// RenderPromptTemplate validates and renders a prompt template.
func RenderPromptTemplate(def PromptTemplateDefinition, values map[string]string) RenderedPromptResult {
	resolved := make(map[string]string)
	for k, v := range values {
		resolved[k] = v
	}

	var missing []string
	for _, param := range def.Parameters {
		val, exists := resolved[param.Name]
		if !exists || strings.TrimSpace(val) == "" {
			if param.DefaultValue != "" {
				resolved[param.Name] = param.DefaultValue
			} else if param.Required {
				missing = append(missing, param.Name)
			}
		}
	}

	if len(missing) > 0 {
		return RenderedPromptResult{
			Messages:      []TemplateMessage{},
			MissingParams: missing,
		}
	}

	var rendered []TemplateMessage
	for _, msg := range def.Messages {
		rendered = append(rendered, TemplateMessage{
			Role:    msg.Role,
			Content: InterpolateTemplate(msg.Content, resolved),
		})
	}

	return RenderedPromptResult{
		Messages:      rendered,
		MissingParams: []string{},
	}
}
