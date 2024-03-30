package _3_generator_v1

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func (s *StubGenerator) parseSource(filename string) ([]*Interface, error) {
	fs := token.NewFileSet()

	parsedFile, err := parser.ParseFile(fs, filename, nil, 0)
	if err != nil {
		return nil, err
	}

	var outputs []*Interface

	packageName := getNameFromIdent(parsedFile.Name)

	// Loop over all top-level declarations
	for _, decl := range parsedFile.Decls {
		generalDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			// Skip anything that isn't a generic declaration.
			// Generic declarations include imports, constants, types or variable declarations.
			continue
		}

		// Loop over the specifications of the declaration
		for _, spec := range generalDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				// Skip anything that isn't a type declaration.
				continue
			}

			interfaceDefinition, ok := typeSpec.Type.(*ast.InterfaceType)
			if !ok {
				// Skip anything that isn't an interface declaration
				continue
			}

			// Extract the details
			outputInterface := &Interface{
				PackageName: packageName,
				Name:        typeSpec.Name.Name,
			}

			for _, method := range interfaceDefinition.Methods.List {
				if len(method.Names) == 0 {
					continue
				}

				funcType, ok := method.Type.(*ast.FuncType)
				if !ok {
					// Skip anything that isn't an interface method
					continue
				}

				outputMethod := &Method{
					Name: method.Names[0].Name,
				}

				outputMethod.Inputs = parseFieldList(funcType.Params)
				outputMethod.Outputs = parseFieldList(funcType.Results)

				outputInterface.Methods = append(outputInterface.Methods, outputMethod)
			}

			outputs = append(outputs, outputInterface)
		}
	}

	return outputs, nil
}

func parseFieldList(fieldList *ast.FieldList) []*Param {
	var out []*Param

	for _, param := range fieldList.List {
		paramType := extractParamType(param.Type)

		if len(param.Names) == 0 {
			out = append(out, &Param{
				Name: "",
				Type: paramType,
			})

			continue
		}

		for _, paramName := range param.Names {
			out = append(out, &Param{
				Name: paramName.Name,
				Type: paramType,
			})
		}
	}

	return out
}

func extractParamType(param ast.Expr) string {
	switch concreteType := param.(type) {
	case *ast.SelectorExpr:
		return getNameFromIdent(concreteType.X.(*ast.Ident)) + "." +
			getNameFromIdent(concreteType.Sel)

	case *ast.StarExpr:
		return "*" + extractParamType(concreteType.X)

	case *ast.Ident:
		return getNameFromIdent(concreteType)

	default:
		fmt.Printf("missing handler for type: %t", concreteType)
		return ""
	}
}

func getNameFromIdent(ident *ast.Ident) string {
	return ident.Name
}

type Interface struct {
	PackageName string
	Name        string
	Methods     []*Method
}

func (i *Interface) String() string {
	out := i.Name + ":\n"

	for _, method := range i.Methods {
		out += "\t" + method.String() + "\n"
	}

	return out
}

type Method struct {
	Name    string
	Inputs  []*Param
	Outputs []*Param
}

func (m *Method) String() string {
	out := m.Name + "("

	for index, param := range m.Inputs {
		if index > 0 {
			out += ", "
		}

		out += param.String()
	}

	out += ")"

	if len(m.Outputs) > 1 {
		out += " ("
	}

	for index, output := range m.Outputs {
		if index > 0 {
			out += ", "
		}

		out += output.String()
	}

	if len(m.Outputs) > 1 {
		out += ")"
	}

	return out
}

type Param struct {
	Name string
	Type string
}

func (p *Param) String() string {
	if p.Name == "" {
		return p.Type
	}

	return p.Name + " " + p.Type
}
