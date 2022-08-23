package nee

import "strings"

type node struct {
	pattern  string  // 特匹配路由，例如：/p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如[doc, tutorial, intro]
	isWild   bool    // 是否精准匹配，part含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert
// Example: /p/:lang/doc
// 只有在第三层节点，即doc节点，pattern才会设置
// p和:lang节点的pattern属性皆为空
// 当匹配结束时，可以使用n.pattern == ""来判断路由规则是否匹配成功
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// search
// /p/python虽能成功匹配到:lang，但:lang的pattern值为空，因此匹配失败
// 递归查询每一层的节点，退出规则是，匹配到了*，匹配失败，或者匹配到了第len(parts)层节点
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" { // 当匹配结束时，可以使用n.pattern == ""来判断路由规则是否匹配成功
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
