package linkedlist

import (
	"fmt"
	"log"
	"strings"
)

type Node struct {
	Data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l LinkedList) ToString() string {
	var result []string
	node := l.head
	for l.size != 0 {
		result = append(result, fmt.Sprintf("%v", node.Data))
		node = node.next
		l.size--
	}

	return strings.Join(result, " ")
}

func TestList() {
	fmt.Println("=============== DEMO LINKEDLIST ===============")

	myList := LinkedList{}
	node1 := &Node{Data: 48}
	node2 := &Node{Data: 18}
	node3 := &Node{Data: 16}
	node4 := &Node{Data: 20}
	node5 := &Node{Data: 22}
	node6 := &Node{Data: 23}
	node7 := &Node{Data: 33}

	fmt.Println("Valor Inicial: 48")
	myList.PutHead(node1)
	fmt.Println(myList.ToString())

	fmt.Println("Agregando datos al Inicio: 18 <- 16 <- 23")
	myList.PutHead(node2)
	myList.PutHead(node3)
	myList.PutHead(node6)
	fmt.Println(myList.ToString())

	fmt.Println("Agregando datos al Final: 20 -> 22")
	myList.PutTail(node4)
	myList.PutTail(node5)
	fmt.Println(myList.ToString())

	fmt.Println("Eliminando Head: 23")
	myList.DeleteHead()
	fmt.Println(myList.ToString())

	fmt.Println("Eliminando Tail: 22")
	myList.DeleteTail()
	fmt.Println(myList.ToString())

	fmt.Println("Eliminando por valor: 18")
	myList.DeleteWithValue(18)
	fmt.Println(myList.ToString())

	fmt.Println("Obteniendo Head: 16")
	head, ok := myList.GetHead()
	if ok {
		fmt.Println(head.Data)
	}

	fmt.Println("Obteniendo Tail: 20")
	tail, ok := myList.GetTail()
	if ok {
		fmt.Println(tail.Data)
	}

	fmt.Println("Obteniendo Indice 1")
	node, err := myList.GetIndex(1)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(node.Data)
	}

	fmt.Println("Agregando Head: 33")
	myList.PutHead(node7)
	fmt.Println(myList.ToString())

	fmt.Println("Eliminando Indice: 3")
	if err := myList.DeleteIndex(3); err != nil {
		log.Println(err)
		return
	}
	fmt.Println(myList.ToString())
}
