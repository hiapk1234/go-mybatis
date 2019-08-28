package GoMybatis

import (
	"github.com/hiapk1234/go-mybatis/ast"
)

type GoMybatisSqlBuilder struct {
	sqlArgTypeConvert     ast.SqlArgTypeConvert
	expressionEngineProxy ExpressionEngineProxy
	enableLog             bool
	nodeParser ast.NodeParser
}

func (it *GoMybatisSqlBuilder) ExpressionEngineProxy() *ExpressionEngineProxy {
	return &it.expressionEngineProxy
}
func (it *GoMybatisSqlBuilder) SqlArgTypeConvert() ast.SqlArgTypeConvert {
	return it.sqlArgTypeConvert
}

func (it GoMybatisSqlBuilder) New(SqlArgTypeConvert ast.SqlArgTypeConvert, expressionEngine ExpressionEngineProxy, log Log, enableLog bool) GoMybatisSqlBuilder {
	it.sqlArgTypeConvert = SqlArgTypeConvert
	it.expressionEngineProxy = expressionEngine
	it.enableLog = enableLog
	it.nodeParser = ast.NodeParser{
		Holder: ast.NodeConfigHolder{
			Convert: SqlArgTypeConvert,
			Proxy:   &expressionEngine,
		},
	}
	return it
}

func (it *GoMybatisSqlBuilder) BuildSql(paramMap map[string]interface{}, nodes []ast.Node) (string, error) {
	//抽象语法树节点构建
	var sql, err = ast.DoChildNodes(nodes, paramMap)
	if err != nil {
		return "", err
	}
	var sqlStr = string(sql)
	return sqlStr, nil
}

func (it *GoMybatisSqlBuilder) SetEnableLog(enable bool) {
	it.enableLog = enable
}
func (it *GoMybatisSqlBuilder) EnableLog() bool {
	return it.enableLog
}


func (it *GoMybatisSqlBuilder) NodeParser() ast.NodeParser {
	return it.nodeParser
}
