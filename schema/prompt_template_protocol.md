# Prompt Template Engine Protocol (v1.0.0)

## Overview
The Prompt Template Engine provides cross-engine standardization for defining, parameterizing, validating, and rendering structured multi-turn prompts for LLM agents.

## Core Schema

```json
{
  "name": "code_refactor_template",
  "version": "1.0.0",
  "description": "Standard prompt template for autonomous code refactoring",
  "parameters": [
    {
      "name": "language",
      "type": "string",
      "required": true,
      "description": "Target programming language"
    },
    {
      "name": "code_snippet",
      "type": "string",
      "required": true,
      "description": "Source code requiring refactoring"
    },
    {
      "name": "strictness",
      "type": "string",
      "required": false,
      "default": "high",
      "description": "Adherence level to style guidelines"
    }
  ],
  "messages": [
    {
      "role": "system",
      "content": "You are an expert {{language}} software architect adhering to {{strictness}} quality standards."
    },
    {
      "role": "user",
      "content": "Refactor the following code snippet cleanly:\n\n```{{language}}\n{{code_snippet}}\n```"
    }
  ]
}
```

## Template Directives & Features
1. **Variable Interpolation**: `{{var_name}}` replaced by parameter value from dictionary.
2. **Missing Parameter Enforcement**: Error raised if a required parameter is omitted and has no default.
3. **Parameter Extraction**: Ability to statically analyze raw templates to discover required variable placeholders.
4. **Multi-Turn Representation**: Renders array of role-annotated messages (`system`, `user`, `assistant`).
