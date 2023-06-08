package main

import "fmt"

type Node struct {
	info string
	next *Node
}

type Stack struct {
	top *Node
}

func stackUp(value string, stack *Stack) {
	fmt.Printf("\n Top %v", stack.top)
	var node *Node
	node = new(Node)
	node.info = value
	node.next = stack.top
	stack.top = node
}

func unstack(stack *Stack) string {
	var value string

	var aux *Node
	aux = stack.top

	if stack.top == nil {
		return "Pilha vazia!"
	}

	value = aux.info
	stack.top = aux.next
	return value
}

func showStack(stack *Stack) {
	aux := stack.top

	if stack.top == nil {
		fmt.Println("\n Pilha Vazia")
	}
	fmt.Println("\n Dados existentes na pilha: ")
	for {
		if aux == nil {
			break
		}
		fmt.Printf("\n Dado: %s", aux.info)
		aux = aux.next
	}

}

func initStack() *Stack {
	var newStack *Stack
	newStack = new(Stack)

	return newStack
}

func main() {
	var option string
	var value string
	var stackInstance *Stack

	stackInstance = initStack()

	for {
		fmt.Println("\n (E) Empilhar (D) Desempilha (F) Finalizar: ")
		fmt.Scanf("%s", &option)

		if option == "E" || option == "e" {
			fmt.Println("\n Qual palavra a ser adicionada a pilha: ")
			fmt.Scanf("%s", &value)
			stackUp(value, stackInstance)
			showStack(stackInstance)
			continue
		}

		if option == "D" || option == "d" {
			response := unstack(stackInstance)
			fmt.Printf("Dado desempilha: %v", response)
			continue
		}

		if option == "F" || option == "f" {
			fmt.Println("\n Saindo do programa!")
			break
		}
	}
}
