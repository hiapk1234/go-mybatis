package GoMybatis

import "github.com/hiapk1234/go-mybatis/ast"

//sql文本构建
type SqlBuilder interface {
	BuildSql(paramMap map[string]interface{}, nodes []ast.Node) (string, error)
	ExpressionEngineProxy() *ExpressionEngineProxy
	SqlArgTypeConvert() ast.SqlArgTypeConvert
	SetEnableLog(enable bool)
	EnableLog() bool
	NodeParser() ast.NodeParser
}
