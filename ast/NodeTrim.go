package ast

import "bytes"

//Trim操作节点
type NodeTrim struct {
	childs []Node
	t      NodeType

	prefix          []byte
	suffix          []byte
	suffixOverrides []byte
	prefixOverrides []byte
}

func (it *NodeTrim) Type() NodeType {
	return NTrim
}

func (it *NodeTrim) Eval(env map[string]interface{}) ([]byte, error) {
	var sql, err = DoChildNodes(it.childs, env)
	if err != nil {
		return nil, err
	}
	if sql == nil {
		return nil, nil
	}
	for {
		if bytes.HasPrefix(sql, []byte(" ")) {
			sql = bytes.Trim(sql, " ")
		} else {
			break
		}
	}
	if it.prefixOverrides != nil {
		var prefixOverridesArray = bytes.Split(it.prefixOverrides, []byte("|"))
		if len(prefixOverridesArray) > 0 {
			for _, v := range prefixOverridesArray {
				sql = bytes.TrimPrefix(sql, []byte(v))
			}
		}
	}
	if it.suffixOverrides != nil {
		var suffixOverrideArray = bytes.Split(it.suffixOverrides, []byte("|"))
		if len(suffixOverrideArray) > 0 {
			for _, v := range suffixOverrideArray {
				sql = bytes.TrimSuffix(sql, []byte(v))
			}
		}
	}
	var newBuffer bytes.Buffer
	newBuffer.WriteString(` `)
	newBuffer.Write(it.prefix)
	newBuffer.WriteString(` `)
	newBuffer.Write(sql)
	newBuffer.WriteString(` `)
	newBuffer.Write(it.suffix)

	var newBufferBytes = newBuffer.Bytes()
	newBuffer.Reset()
	return newBufferBytes, nil
}
