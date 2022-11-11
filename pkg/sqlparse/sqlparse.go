package sqlparse

import (
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"strings"
)

var (
	p *parser.Parser
)

type colX struct {
	colAsNames []string
}

func (v *colX) Enter(in ast.Node) (ast.Node, bool) {
	fieldList := in.(*ast.SelectStmt).Fields.Fields
	for _, field := range fieldList {
		if field.AsName.String() != "" {
			v.colAsNames = append(v.colAsNames, field.AsName.String())
		} else {
			name := field.Text()
			if strings.Contains(name, ".") {
				name = strings.Split(name, ".")[1]
			}
			v.colAsNames = append(v.colAsNames, name)
		}
	}
	return in, true
}

func (v *colX) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func New() {
	p = parser.New()
}

func GetColAsName(sql string) (asName string, err error) {
	stmt, err := p.ParseOneStmt(sql, "", "")
	if err != nil {
		return "", err
	}
	v := &colX{}
	stmt.Accept(v)
	return strings.Join(v.colAsNames, ", "), err
}
