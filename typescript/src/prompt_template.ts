/**
 * Prompt Template Engine
 * Standardized cross-engine variable interpolation, validation, and multi-turn message formatting.
 */

export interface TemplateParameter {
  name: string;
  type: string;
  required: boolean;
  defaultValue?: string;
  description?: string;
}

export interface TemplateMessage {
  role: "system" | "user" | "assistant";
  content: string;
}

export interface PromptTemplateDefinition {
  name: string;
  version: string;
  description?: string;
  parameters: TemplateParameter[];
  messages: TemplateMessage[];
}

export interface RenderedMessage {
  role: "system" | "user" | "assistant";
  content: string;
}

export class PromptTemplateEngine {
  /**
   * Extracts unique placeholder names from a template string (e.g. {{variable}}).
   */
  public static extractVariables(templateText: string): string[] {
    const regex = /\{\{([a-zA-Z0-9_]+)\}\}/g;
    const matches = new Set<string>();
    let match: RegExpExecArray | null;
    while ((match = regex.exec(templateText)) !== null) {
      matches.add(match[1]);
    }
    return Array.from(matches);
  }

  /**
   * Interpolates variables into a text template.
   */
  public static interpolate(text: string, values: Record<string, string>): string {
    return text.replace(/\{\{([a-zA-Z0-9_]+)\}\}/g, (match, varName) => {
      if (varName in values) {
        return values[varName];
      }
      return match;
    });
  }

  /**
   * Validates parameters against the template definition and renders multi-turn messages.
   */
  public static render(
    definition: PromptTemplateDefinition,
    inputValues: Record<string, string>
  ): { messages: RenderedMessage[]; missingParams: string[] } {
    const missing: string[] = [];
    const resolvedValues: Record<string, string> = { ...inputValues };

    for (const param of definition.parameters) {
      if (!(param.name in resolvedValues) || resolvedValues[param.name].trim() === "") {
        if (param.defaultValue !== undefined) {
          resolvedValues[param.name] = param.defaultValue;
        } else if (param.required) {
          missing.push(param.name);
        }
      }
    }

    if (missing.length > 0) {
      return { messages: [], missingParams: missing };
    }

    const messages: RenderedMessage[] = definition.messages.map((msg) => ({
      role: msg.role,
      content: this.interpolate(msg.content, resolvedValues),
    }));

    return { messages, missingParams: [] };
  }
}
