package linked_list
type Node struct {
	value int
	next *Node
}

func insertNode(head,node *Node){
	if head == nil{
		head.next = node
	}
	temp := head
	for{
		if temp.next == nil{
			break
		}
		temp = temp.next
	}

	temp.next = node



}

func deleteNode(head *Node,target int) *Node{
	tempP := head
	tempQ := head.next
	//头结点是要删除的节点
	if head.value == target {
		head = head.next
		return head
	}
	for{
		if tempQ == nil {
			return nil
		}
		if tempQ.value == target{
			tempP.next = tempQ.next
			break
		}

		tempP = tempQ
		tempQ = tempQ.next


	}
	return head
}