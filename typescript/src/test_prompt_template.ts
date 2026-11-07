import assert from "assert";
import { PromptTemplateEngine, PromptTemplateDefinition } from "./prompt_template.js";

const template: PromptTemplateDefinition = {
  name: "code_reviewer",
  version: "1.0.0",
  parameters: [
    { name: "language", type: "string", required: true },
    { name: "code", type: "string", required: true },
    { name: "style", type: "string", required: false, defaultValue: "standard" },
  ],
  messages: [
    { role: "system", content: "You are a {{language}} reviewer using {{style}} standards." },
    { role: "user", content: "Check this code:\n{{code}}" },
  ],
};

// 1. Variable extraction
const vars = PromptTemplateEngine.extractVariables(template.messages[0].content);
assert.deepStrictEqual(vars, ["language", "style"]);

// 2. Successful rendering with defaults
const result = PromptTemplateEngine.render(template, {
  language: "TypeScript",
  code: "console.log('hello');",
});
assert.strictEqual(result.missingParams.length, 0);
assert.strictEqual(result.messages[0].content, "You are a TypeScript reviewer using standard standards.");
assert.strictEqual(result.messages[1].content, "Check this code:\nconsole.log('hello');");

// 3. Missing required parameter detection
const failure = PromptTemplateEngine.render(template, {
  style: "strict",
});
assert.deepStrictEqual(failure.missingParams.sort(), ["code", "language"].sort());

console.log("TypeScript PromptTemplateEngine tests passed cleanly!");
