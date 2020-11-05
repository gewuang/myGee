package gee

type node struct {
	path   string  // url路径
	isWild bool    // 是否是模糊查询
	part   string  // 当前匹配字段
	child  *[]node // 子节点集合
}

// 将parts中的内容插入到树中
func (*node) insert(path string, parts []string, height int) {

}

// 从树中查询到结果并返回
func (*node) search(parts []string, height int) *node {

	return nil
}
