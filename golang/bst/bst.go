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

	if n.valor < valor {
		n.direita = n.direita.Inserir(valor)
	} else {
		n.esquerda = n.esquerda.Inserir(valor)
	}

	return n
}

func (n *Node) ExibirEmOrdem() {
	if n != nil {
		n.esquerda.ExibirEmOrdem()
		fmt.Println(n.valor)
		n.direita.ExibirEmOrdem()
	}
}

func (n *Node) Buscar(valor int) bool {
	if n == nil || n.valor == valor {
		return n != nil
	}

	if valor < n.valor {
		return n.esquerda.Buscar(valor)
	} else {
		return n.direita.Buscar(valor)
	}
}

func (n *Node) Remover(valor int) *Node {
	if n == nil {
		return n
	}

	if valor < n.valor {
		return n.esquerda.Remover(valor)
	} else if valor > n.valor {
		return n.direita.Remover(valor)
	} else {
		if n.esquerda == nil {
			return n.direita
		} else if n.direita == nil {
			return n.esquerda
		}

		n.encontrarMinimo()
		n.direita.Remover(valor)
	}

	return n
}

func (n *Node) encontrarMinimo() *Node {
	for n.esquerda != nil {
		n = n.esquerda
	}

	return n

}