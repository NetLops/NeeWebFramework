package nee

import "strings"

type prefixTreeNode struct {
	pattern  string            // 全匹配	/p/:lang/name
	part     string            //	:lang
	children []*prefixTreeNode // 儿子们
	isWild   bool              // 是否包含特殊符号	// example: * :
}

// matchChildrenByPart 匹配第一个成功的儿子
func (p *prefixTreeNode) matchChildrenByPart(part string) *prefixTreeNode {
	for _, child := range p.children {
		if child.part == part || p.isWild {
			return child
		}
	}
	return nil
}

// 找到所有匹配成功的节点
func (p *prefixTreeNode) matchAllChildrenByPart(part string) []*prefixTreeNode {
	matchPrefixTreeNodes := make([]*prefixTreeNode, 0)
	for _, child := range p.children {
		if child.part == part || child.isWild {
			matchPrefixTreeNodes = append(matchPrefixTreeNodes, child)
		}
	}
	return matchPrefixTreeNodes
}

// insert
// Example: /p/:lang/doc
// 最后一层pattern 存/p/:lang/doc
func (p *prefixTreeNode) insert(pattern string, parts []string, height int) {
	// 最后一层才存储匹配的完整字符串"/p/:lange/name"
	if len(parts) == height {
		p.pattern = pattern
		return
	}
	part := parts[height]
	child := p.matchChildrenByPart(part)
	if child == nil {
		child = &prefixTreeNode{
			part:     part,
			children: make([]*prefixTreeNode, 0),
			isWild:   part[0] == '*' || part[0] == ':',
			pattern:  "",
		}
		p.children = append(p.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (p *prefixTreeNode) search(parts []string, height int) *prefixTreeNode {
	// 是否匹配到最后一层，或者是全词匹配
	if len(parts) == height || strings.HasPrefix(p.part, "*") {
		// 比如原匹配是 /p/:lang/name  现在的路由过来的是 /p/:lang 所以应该不给予匹配
		if p.pattern == "" {
			return nil
		}

		return p
	}
	part := parts[height]
	// 找到它旗下的所有儿子
	childrens := p.matchAllChildrenByPart(part)
	for _, children := range childrens {
		// 继续往下搜
		resultNode := children.search(parts, height+1)
		// 找到了 会返回 prefixNode
		if resultNode != nil {
			return resultNode
		}
	}
	return nil
}
