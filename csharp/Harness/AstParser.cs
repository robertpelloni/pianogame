using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Microsoft.CodeAnalysis;
using Microsoft.CodeAnalysis.CSharp;
using Microsoft.CodeAnalysis.CSharp.Syntax;

namespace Harness
{
    public class Definition
    {
        public string Name { get; set; } = string.Empty;
        public string Kind { get; set; } = string.Empty;
    }

    public class Dependency
    {
        public string Path { get; set; } = string.Empty;
    }

    public class FileSyntaxMap
    {
        public string FilePath { get; set; } = string.Empty;
        public List<Definition> Definitions { get; set; } = new List<Definition>();
        public List<Dependency> Dependencies { get; set; } = new List<Dependency>();
    }

    public class AstParser
    {
        public FileSyntaxMap ParseFile(string filePath)
        {
            var content = File.ReadAllText(filePath);
            var tree = CSharpSyntaxTree.ParseText(content);
            var root = tree.GetCompilationUnitRoot();

            var fileMap = new FileSyntaxMap { FilePath = filePath };

            // Extract using directives as dependencies
            var usings = root.DescendantNodes().OfType<UsingDirectiveSyntax>();
            foreach (var u in usings)
            {
                fileMap.Dependencies.Add(new Dependency { Path = u.Name.ToString() });
            }

            // Extract classes
            var classes = root.DescendantNodes().OfType<ClassDeclarationSyntax>();
            foreach (var c in classes)
            {
                fileMap.Definitions.Add(new Definition { Name = c.Identifier.Text, Kind = "Class" });
            }

            // Extract interfaces
            var interfaces = root.DescendantNodes().OfType<InterfaceDeclarationSyntax>();
            foreach (var i in interfaces)
            {
                fileMap.Definitions.Add(new Definition { Name = i.Identifier.Text, Kind = "Interface" });
            }

            // Extract methods
            var methods = root.DescendantNodes().OfType<MethodDeclarationSyntax>();
            foreach (var m in methods)
            {
                fileMap.Definitions.Add(new Definition { Name = m.Identifier.Text, Kind = "Method" });
            }

            return fileMap;
        }

        public List<Definition> ExtractDefinitions(FileSyntaxMap fileMap)
        {
            return fileMap?.Definitions ?? new List<Definition>();
        }

        public List<Dependency> MapDependencies(FileSyntaxMap fileMap)
        {
            return fileMap?.Dependencies ?? new List<Dependency>();
        }
    }
}
