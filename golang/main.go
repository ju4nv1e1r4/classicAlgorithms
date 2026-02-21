package main

import (
	"classicAlgorithms/bst"
	"fmt"
)

func main() {
	fmt.Println("Iniciando a ávore")

	fmt.Println("--- Inserindo valores ---")
	// Vamos criar uma árvore com a raiz 50
	valoresParaInserir := [...]int{50, 30, 70, 20, 40, 60, 80, 35}

	var n *bst.Node

	for _, valor := range valoresParaInserir {
		n = n.Inserir(valor)
		fmt.Println("Inserido: ", valor)
	}

	fmt.Println("\n--- Árvore em Ordem Crescente ---")
	// Se tudo deu certo, veremos do 20 ao 80 em ordem!
	n.ExibirEmOrdem()

	fmt.Println("\n--- Testando a Busca ---")
	alvo1 := 40
	resultado1 := n.Buscar(alvo1)
	if resultado1 {
		fmt.Printf("O número %d foi encontrado.", alvo1)
	} else {
		fmt.Printf("O número %d não foi encontrado.", alvo1)
	}

	fmt.Println("\n--- Testando a Busca ---")
	alvo2 := 99
	resultado2 := n.Buscar(alvo2)
	if resultado2 {
		fmt.Printf("O número %d foi encontrado.", alvo2)
	} else {
		fmt.Printf("O número %d não foi encontrado.", alvo2)
	}

	fmt.Println("\n--- Testando a Remoção ---")
	fmt.Println("Removendo 20 (Nó folha)...")
	n = n.Remover(20)

	fmt.Println("\n--- Removendo 40 (Nó com 1 filho) ---")
	fmt.Println("Removendo 40 (Nó folha)...")
	n = n.Remover(40)

	fmt.Println("\n--- Removendo 50 (Nó com 2 filhos - A Raiz!) ---")
	fmt.Println("Removendo 50 (Nó folha)...")
	n = n.Remover(50)

	fmt.Println("\n--- Árvore Final em Ordem Crescente ---")
	n.ExibirEmOrdem()
}
