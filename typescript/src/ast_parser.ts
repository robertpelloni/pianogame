import { parse, TSESTree } from '@typescript-eslint/typescript-estree';
import * as fs from 'fs';

export interface Definition {
    name: string;
    kind: string;
}

export interface Dependency {
    path: string;
}

export interface FileSyntaxMap {
    filePath: string;
    definitions: Definition[];
    dependencies: Dependency[];
}

export class AstParser {
    public parseFile(filePath: string): FileSyntaxMap {
        const content = fs.readFileSync(filePath, 'utf-8');
        const ast = parse(content, { loc: true, range: true });

        const map: FileSyntaxMap = {
            filePath,
            definitions: [],
            dependencies: [],
        };

        const traverse = (node: TSESTree.Node) => {
            switch (node.type) {
                case TSESTree.AST_NODE_TYPES.ImportDeclaration:
                    map.dependencies.push({
                        path: String(node.source.value)
                    });
                    break;
                case TSESTree.AST_NODE_TYPES.ClassDeclaration:
                    if (node.id) {
                        map.definitions.push({
                            name: node.id.name,
                            kind: 'Class'
                        });
                    }
                    break;
                case TSESTree.AST_NODE_TYPES.TSInterfaceDeclaration:
                    if (node.id) {
                        map.definitions.push({
                            name: node.id.name,
                            kind: 'Interface'
                        });
                    }
                    break;
                case TSESTree.AST_NODE_TYPES.FunctionDeclaration:
                case TSESTree.AST_NODE_TYPES.MethodDefinition:
                    let name = '';
                    if (node.type === TSESTree.AST_NODE_TYPES.FunctionDeclaration && node.id) {
                        name = node.id.name;
                    } else if (node.type === TSESTree.AST_NODE_TYPES.MethodDefinition && node.key.type === TSESTree.AST_NODE_TYPES.Identifier) {
                        name = node.key.name;
                    }
                    if (name) {
                        map.definitions.push({
                            name: name,
                            kind: 'Function/Method'
                        });
                    }
                    break;
            }

            // Simple recursive traversal
            for (const key of Object.keys(node)) {
                const child = (node as any)[key];
                if (child && typeof child === 'object') {
                    if (Array.isArray(child)) {
                        child.forEach(c => {
                            if (c && typeof c.type === 'string') traverse(c);
                        });
                    } else if (typeof child.type === 'string') {
                        traverse(child);
                    }
                }
            }
        };

        traverse(ast);

        return map;
    }

    public extractDefinitions(fileMap: FileSyntaxMap): Definition[] {
        return fileMap.definitions;
    }

    public mapDependencies(fileMap: FileSyntaxMap): Dependency[] {
        return fileMap.dependencies;
    }
}
