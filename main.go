package main

func main() {
	printList()
	//var stack Stack
	//var queue Queue
	//
	//for {
	//	fmt.Println("=============== PILAS & COLAS ===============")
	//	fmt.Println("=============== MENU ===============\n" +
	//		"1.- Agregar item a Pila\n" +
	//		"2.- Agregar item a Cola\n" +
	//		"3.- Quitar item de Pila\n" +
	//		"4.- Quitar item de Cola\n" +
	//		"5.- Mostrar items de Pila\n" +
	//		"6.- Mostrar items de Cola\n" +
	//		"7.- Salir")
	//
	//	var option string
	//	_, err := fmt.Scan(&option)
	//	if err != nil {
	//		log.Printf("Error al escanear opcion del menu, err: %v\n", err)
	//		return
	//	}
	//
	//	switch option {
	//	case "1":
	//		fmt.Println("Ingrese item a agregar:")
	//		var item string
	//		_, err := fmt.Scan(&item)
	//		if err != nil {
	//			log.Printf("Error al escanear item, err: %v\n", err)
	//			return
	//		}
	//
	//		stack.Push(item)
	//		fmt.Printf("Elementos actuales: %v\n", stack.items)
	//	case "2":
	//		fmt.Println("Ingrese item a agregar:")
	//		var item string
	//		_, err := fmt.Scan(&item)
	//		if err != nil {
	//			log.Printf("Error al escanear item, err: %v\n", err)
	//			return
	//		}
	//
	//		queue.Enqueue(item)
	//		fmt.Printf("Elementos actuales: %v\n", queue.items)
	//	case "3":
	//		item, ok := stack.Pop()
	//		if !ok {
	//			fmt.Println("La pila no tiene items")
	//			continue
	//		}
	//
	//		fmt.Printf("Elmento saliente: %v\n", item)
	//		fmt.Printf("Elementos actuales: %v\n", stack.items)
	//	case "4":
	//		item, ok := queue.Dequeue()
	//		if !ok {
	//			fmt.Println("La cola no tiene items")
	//			continue
	//		}
	//
	//		fmt.Printf("Elmento saliente: %v\n", item)
	//		fmt.Printf("Elementos actuales: %v\n", queue.items)
	//	case "5":
	//		fmt.Printf("Elementos actuales: %v\n", stack.items)
	//	case "6":
	//		fmt.Printf("Elementos actuales: %v\n", queue.items)
	//	case "7":
	//		os.Exit(0)
	//	default:
	//		fmt.Println("Opción invalida.")
	//	}
	//}
}
