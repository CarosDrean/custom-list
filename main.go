package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CarosDrean/custom-list/linkedlist"
    "github.com/CarosDrean/custom-list/queue"
    "github.com/CarosDrean/custom-list/stack"
)

type IStack interface {
	Push(item interface{})
	Pop() (interface{}, bool)
	ToString() string
}

type IQueue interface {
	Enqueue(item interface{})
	Dequeue() (interface{}, bool)
	ToString() string
}

func main() {
    var queueSimple queue.Queue
    var stackSimple stack.Stack
    var queueLinkedList queue.QueueLinkedList
    var stackLinkedList stack.StackLinkedList
    
	for {
		fmt.Println("=============== MENU ===============\n" +
			"1.- Ejecutar test de LinkedList personalizado\n" +
			"2.- Menu de Pilas y Colas simple\n" +
			"3.- Menu de Pilas y Colas con LinkedList\n" +
			"4.- Salir")

		fmt.Print("Ingrese Opción: ")

		var option string
		_, err := fmt.Scan(&option)
		if err != nil {
			log.Printf("Error al escanear opcion del menu, err: %v\n", err)
			return
		}

		switch option {
		case "1":
			linkedlist.TestList()
		case "2":
			menuQueueStack(&stackSimple, &queueSimple)
		case "3":
			menuQueueStack(&stackLinkedList, &queueLinkedList)
		case "4":
			os.Exit(0)
		default:
			fmt.Println("Opción invalida.")
		}
	}
}

func menuQueueStack(stack IStack, queue IQueue) {
	for {
		fmt.Println("=============== PILAS & COLAS ===============")
		fmt.Println("=============== MENU ===============\n" +
			"1.- Agregar item a Pila\n" +
			"2.- Agregar item a Cola\n" +
			"3.- Quitar item de Pila\n" +
			"4.- Quitar item de Cola\n" +
			"5.- Mostrar items de Pila\n" +
			"6.- Mostrar items de Cola\n" +
			"7.- Ir al menu anterior\n" +
			"8.- Salir")

		fmt.Print("Ingrese Opción: ")

		var option string
		_, err := fmt.Scan(&option)
		if err != nil {
			log.Printf("Error al escanear opcion del menu, err: %v\n", err)
			return
		}

		switch option {
		case "1":
			fmt.Print("Ingrese item a agregar: ")
			var item string
			_, err := fmt.Scan(&item)
			if err != nil {
				log.Printf("Error al escanear item, err: %v\n", err)
				return
			}

			stack.Push(item)
			fmt.Printf("Elementos actuales: %v\n", stack.ToString())
		case "2":
			fmt.Print("Ingrese item a agregar: ")
			var item string
			_, err := fmt.Scan(&item)
			if err != nil {
				log.Printf("Error al escanear item, err: %v\n", err)
				return
			}

			queue.Enqueue(item)
			fmt.Printf("Elementos actuales: %v\n", queue.ToString())
		case "3":
			item, ok := stack.Pop()
			if !ok {
				fmt.Println("La pila no tiene items")
				continue
			}

			fmt.Printf("Elmento saliente: %v\n", item)
			fmt.Printf("Elementos actuales: %v\n", stack.ToString())
		case "4":
			item, ok := queue.Dequeue()
			if !ok {
				fmt.Println("La cola no tiene items")
				continue
			}

			fmt.Printf("Elmento saliente: %v\n", item)
			fmt.Printf("Elementos actuales: %v\n", queue.ToString())
		case "5":
			fmt.Printf("Elementos actuales: %v\n", stack.ToString())
		case "6":
			fmt.Printf("Elementos actuales: %v\n", queue.ToString())
		case "7":
			return
		case "8":
			os.Exit(0)
		default:
			fmt.Println("Opción invalida.")
		}
	}
}
