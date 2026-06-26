use std::error::Error;
use std::fs;
use std::path::Path;
use syn::{visit::Visit, ItemFn, ItemStruct, ItemEnum, ItemTrait, UseTree};

#[derive(Debug, Clone)]
pub struct Definition {
    pub name: String,
    pub kind: String, // "Function", "Struct", "Enum", "Trait"
}

#[derive(Debug, Clone)]
pub struct Dependency {
    pub path: String,
}

#[derive(Debug, Clone)]
pub struct FileSyntaxMap {
    pub file_path: String,
    pub definitions: Vec<Definition>,
    pub dependencies: Vec<Dependency>,
}

struct AstVisitor {
    definitions: Vec<Definition>,
    dependencies: Vec<Dependency>,
}

impl AstVisitor {
    fn new() -> Self {
        AstVisitor {
            definitions: Vec::new(),
            dependencies: Vec::new(),
        }
    }
}

// We implement the syn::visit::Visit trait to traverse the AST nodes
impl<'ast> Visit<'ast> for AstVisitor {
    fn visit_item_fn(&mut self, node: &'ast ItemFn) {
        self.definitions.push(Definition {
            name: node.sig.ident.to_string(),
            kind: "Function".to_string(),
        });
        syn::visit::visit_item_fn(self, node);
    }

    fn visit_item_struct(&mut self, node: &'ast ItemStruct) {
        self.definitions.push(Definition {
            name: node.ident.to_string(),
            kind: "Struct".to_string(),
        });
        syn::visit::visit_item_struct(self, node);
    }

    fn visit_item_enum(&mut self, node: &'ast ItemEnum) {
        self.definitions.push(Definition {
            name: node.ident.to_string(),
            kind: "Enum".to_string(),
        });
        syn::visit::visit_item_enum(self, node);
    }

    fn visit_item_trait(&mut self, node: &'ast ItemTrait) {
        self.definitions.push(Definition {
            name: node.ident.to_string(),
            kind: "Trait".to_string(),
        });
        syn::visit::visit_item_trait(self, node);
    }

    fn visit_use_tree(&mut self, node: &'ast UseTree) {
        // Very basic dependency extraction for demonstration.
        // In a full implementation, we'd recursively parse the UseTree to build the full path.
        if let UseTree::Path(use_path) = node {
            self.dependencies.push(Dependency {
                path: use_path.ident.to_string(),
            });
        }
        syn::visit::visit_use_tree(self, node);
    }
}

pub struct AstParser;

impl AstParser {
    pub fn parse_file(file_path: &str) -> Result<FileSyntaxMap, Box<dyn Error>> {
        let content = fs::read_to_string(Path::new(file_path))?;

        let syntax_tree = syn::parse_file(&content)?;

        let mut visitor = AstVisitor::new();
        visitor.visit_file(&syntax_tree);

        Ok(FileSyntaxMap {
            file_path: file_path.to_string(),
            definitions: visitor.definitions,
            dependencies: visitor.dependencies,
        })
    }

    pub fn extract_definitions(file_map: &FileSyntaxMap) -> Vec<Definition> {
        file_map.definitions.clone()
    }

    pub fn map_dependencies(file_map: &FileSyntaxMap) -> Vec<Dependency> {
        file_map.dependencies.clone()
    }
}
