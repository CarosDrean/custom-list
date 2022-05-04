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
	head   *Node
	tail   *Node
	length int
}

func (l LinkedList) ToString() string {
	var result []string
	node := l.head
	for l.length != 0 {
		result = append(result, fmt.Sprintf("%v", node.Data))
		node = node.next
		l.length--
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

	fmt.Println("Valor Inicial")
	myList.PutHead(node1)
	fmt.Println(myList.ToString())

	fmt.Println("Agregando datos al Inicio")
	myList.PutHead(node2)
	myList.PutHead(node3)
	myList.PutHead(node6)
	fmt.Println(myList.ToString())

	fmt.Println("Agregando datos al Final")
	myList.PutTail(node4)
	myList.PutTail(node5)
	fmt.Println(myList.ToString())

	fmt.Println("Eliminando Head")
	myList.DeleteHead()
	fmt.Println(myList.ToString())

	fmt.Println("Eliminando Tail")
	myList.DeleteTail()
	fmt.Println(myList.ToString())

	fmt.Println("Eliminando por valor: 18")
	myList.DeleteWithValue(18)
	fmt.Println(myList.ToString())

	fmt.Println("Obteniendo Head")
	head, ok := myList.GetHead()
	if ok {
		fmt.Println(head.Data)
	}

	fmt.Println("Obteniendo Tail")
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

	fmt.Println("Agregando Head")
	myList.PutHead(node7)
	fmt.Println(myList.ToString())

	fmt.Println("Eliminando Indice 3")
	if err := myList.DeleteIndex(3); err != nil {
		log.Println(err)
		return
	}
	fmt.Println(myList.ToString())
}
