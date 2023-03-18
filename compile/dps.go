package main


type Tree struct {
	left *Tree
	right *Tree
	val  int
}


func dps(root *Tree , value1 int,value2 int) {
	list1:= make([]*Tree, 0)
	list2:= make([]*Tree, 0)
	list1 = append(list1,root )
	list2 = append(list2,root )
	_,list1=dpsHelp(root,list1,value1)
	_,list2=dpsHelp(root,list2,value2)

}
func dpsHelp(root *Tree,list []*Tree,value int) (bool,[]*Tree){
	if root.val == value {
		return true,list
	}
	result := false
	if root.left != nil {
		list = append(list,root.left)
		result,list = dpsHelp(root.left,list,value)
		if result{
			return true,list
		}
		list = list[0:len(list)-1]
	}
	if root.right != nil{
		list = append(list,root.right)
		result ,list= dpsHelp(root.left,list,value)
		if result{
			return true,list
		}
		list = list[0:len(list)-1]
	}
	return false,list
}
