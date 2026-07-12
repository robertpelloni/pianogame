package com.harness;

import com.github.javaparser.JavaParser;
import com.github.javaparser.ParseResult;
import com.github.javaparser.ast.CompilationUnit;
import com.github.javaparser.ast.body.ClassOrInterfaceDeclaration;
import com.github.javaparser.ast.body.MethodDeclaration;
import com.github.javaparser.ast.ImportDeclaration;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;

public class AstParser {

    public static class Definition {
        public String name;
        public String kind;

        public Definition(String name, String kind) {
            this.name = name;
            this.kind = kind;
        }
    }

    public static class Dependency {
        public String path;

        public Dependency(String path) {
            this.path = path;
        }
    }

    public static class FileSyntaxMap {
        public String filePath;
        public List<Definition> definitions = new ArrayList<>();
        public List<Dependency> dependencies = new ArrayList<>();

        public FileSyntaxMap(String filePath) {
            this.filePath = filePath;
        }
    }

    public FileSyntaxMap parseFile(String filePath) throws FileNotFoundException {
        JavaParser parser = new JavaParser();
        ParseResult<CompilationUnit> result = parser.parse(new File(filePath));

        if (!result.isSuccessful() || result.getResult().isEmpty()) {
            throw new RuntimeException("Failed to parse file: " + filePath);
        }

        CompilationUnit cu = result.getResult().get();
        FileSyntaxMap map = new FileSyntaxMap(filePath);

        // Extract imports as dependencies
        for (ImportDeclaration id : cu.findAll(ImportDeclaration.class)) {
            map.dependencies.add(new Dependency(id.getNameAsString()));
        }

        // Extract classes/interfaces
        for (ClassOrInterfaceDeclaration cid : cu.findAll(ClassOrInterfaceDeclaration.class)) {
            String kind = cid.isInterface() ? "Interface" : "Class";
            map.definitions.add(new Definition(cid.getNameAsString(), kind));
        }

        // Extract methods
        for (MethodDeclaration md : cu.findAll(MethodDeclaration.class)) {
            map.definitions.add(new Definition(md.getNameAsString(), "Method"));
        }

        return map;
    }

    public List<Definition> extractDefinitions(FileSyntaxMap fileMap) {
        return fileMap != null ? fileMap.definitions : new ArrayList<>();
    }

    public List<Dependency> mapDependencies(FileSyntaxMap fileMap) {
        return fileMap != null ? fileMap.dependencies : new ArrayList<>();
    }
}
