class Node:
    def __init__(self, valor):
        self.esquerda = None
        self.direita = None
        self.valor = valor


class Arvore:
    def __init__(self):
        self.raiz = None

    def inserir(self, valor):
        self.raiz = self._inserir_recursivo(self.raiz, valor)

    def _inserir_recursivo(self, no_atual, valor):
        if no_atual is None:
            return Node(valor)
        
        if valor < no_atual.valor:
            no_atual.esquerda = self._inserir_recursivo(no_atual.esquerda, valor)
        elif valor > no_atual.valor:
            no_atual.direita = self._inserir_recursivo(no_atual.direita, valor)
        
        return no_atual
    
    def exibir_em_ordem(self):
        self._exibir_em_ordem_recursivo(self.raiz)

    def _exibir_em_ordem_recursivo(self, no_atual):
        if no_atual is not None:
            self._exibir_em_ordem_recursivo(no_atual.esquerda)
            print(no_atual.valor)
            self._exibir_em_ordem_recursivo(no_atual.direita)

    def buscar(self, valor):
        return self._buscar_recursivo(self.raiz, valor)

    def _buscar_recursivo(self, no_atual, valor):
        if no_atual is None or no_atual.valor == valor:
            return no_atual
        
        if valor < no_atual.valor:
            return self._buscar_recursivo(no_atual.esquerda, valor)
        else:
            return self._buscar_recursivo(no_atual.direita, valor)

    def remover(self, valor):
        self.raiz = self._remover_recursivo(self.raiz, valor)

    def _remover_recursivo(self, no_atual, valor):
        if no_atual is None:
            return no_atual
        
        if valor < no_atual.valor:
            no_atual.esquerda = self._remover_recursivo(no_atual.esquerda, valor)
        elif valor > no_atual.valor:
            no_atual.direita = self._remover_recursivo(no_atual.direita, valor)
        else:
            if no_atual.esquerda is None:
                return no_atual.direita
            elif no_atual.direita is None:
                return no_atual.esquerda
            
            no_atual.valor = self._encontrar_minimo(no_atual.direita)
            no_atual.direita = self._remover_recursivo(no_atual.direita, no_atual.valor)
        
        return no_atual
    
    def _encontrar_minimo(self, no_atual):
        while no_atual.esquerda is not None:
            no_atual = no_atual.esquerda
        return no_atual.valor
