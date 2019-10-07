package wax

import (
	"fmt"
	"strings"
)

type disasmTree []*disasmLineComponents

type disasmLineComponents struct {
	binary   []byte
	mnemonic string
	indent   int
	nest     disasmTree
}

func (dt disasmTree) Join(delim string) string {
	indent := []string{}
	binaryLines := []string{}
	mnemonicLines := []string{}
	maxHexLen := 0

	flat := flattenDisasmTree(dt, 0)
	for _, node := range flat {
		h := hexString(node.binary)
		if len(h)+(node.indent*2) > maxHexLen {
			maxHexLen = len(h) + (node.indent * 2)
		}
		indent = append(indent, nspace(node.indent*2))
		binaryLines = append(binaryLines, h)
		mnemonicLines = append(mnemonicLines, node.mnemonic)
	}

	f := fmt.Sprintf("%%s%%-%ds %%s%%s", maxHexLen)
	result := []string{}
	for i := range binaryLines {
		result = append(result, fmt.Sprintf(f, indent[i], binaryLines[i], indent[i], mnemonicLines[i]))
	}

	return strings.Join(result, delim)
}

func nspace(n int) string {
	f := fmt.Sprintf("%%%ds", n)
	return fmt.Sprintf(f, "")
}

func flattenDisasmTree(dt disasmTree, indent int) []*disasmLineComponents {
	result := []*disasmLineComponents{}
	for _, node := range dt {
		node.indent = indent
		result = append(result, node)
		if node.nest != nil {
			children := flattenDisasmTree(node.nest, indent+1)
			result = append(result, children...)
		}
	}
	return result
}

func hexString(bs []byte) string {
	result := []string{}
	for _, b := range bs {
		result = append(result, fmt.Sprintf("%02x", b))
	}
	return strings.Join(result, " ")
}
