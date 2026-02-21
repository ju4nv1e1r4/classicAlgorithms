package bst

import "fmt"

type Node struct {
	valor int
	esquerda *Node
	direita *Node
}

func (n *Node) Inserir(valor int) *Node {
	if n == nil {
		return &Node{valor: valor}
	}

	if valor < n.valor {
		n.esquerda = n.esquerda.Inserir(valor)
	} else if valor > n.valor {
		n.direita = n.direita.Inserir(valor)
	}

	return n
}

func (n *Node) ExibirEmOrdem() {
	if n == nil {
		return
	}

	n.esquerda.ExibirEmOrdem()
	fmt.Printf("%d ", n.valor)
	n.direita.ExibirEmOrdem()
}

func (n *Node) Buscar(valor int) bool {
	if n == nil {
		return false
	}
	if n.valor == valor {
		return true
	}

	if valor < n.valor {
		return n.esquerda.Buscar(valor)
	}
	return n.direita.Buscar(valor)
}

func (n *Node) Remover(valor int) *Node {
	if n == nil {
		return nil
	}

	if valor < n.valor {
		n.esquerda = n.esquerda.Remover(valor)
	} else if valor > n.valor {
		n.direita = n.direita.Remover(valor)
	} else {
		if n.esquerda == nil {
			return n.direita
		} else if n.direita == nil {
			return n.esquerda
		}

		substituto := n.direita.encontrarMinimo()
		n.valor = substituto.valor
		n.direita = n.direita.Remover(substituto.valor)
	}

	return n
}

func (n *Node) encontrarMinimo() *Node {
	atual := n
	for atual.esquerda != nil {
		atual = atual.esquerda
	}

	return atual
}
