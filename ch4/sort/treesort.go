package sort

type tree struct {
	value int
	left,right *tree
}

func TreeSort(values []int){
	var root *tree
	for _,v := range values {
		root = add(root,v)
	}
	appendValue(values[:0],root)
}

func appendValue(values []int,tree *tree) []int {
	if tree != nil {
		values = appendValue(values,tree.left)
		values = append(values,tree.value)
		values = appendValue(values,tree.right)
	}
	return values
}

func add(t *tree,value int) *tree {
	if t == nil {
		t = new(tree)
		t.value =value
		return t
	}
	if(value <t.value){
		t.left = add(t.left,value)
	}else{
		t.right = add(t.right,value)
	}
	return t
}
