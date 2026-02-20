# Estudos

## 1. A Base: Árvore Binária de Busca (BST)

### Como funciona:

Uma estrutura composta por "nós". Cada nó contém um valor e dois ponteiros (ou referências): um para o filho à esquerda e outro para o filho à direita. A regra de ouro é a organização: para qualquer nó dado, todos os valores na sua subárvore à esquerda devem ser menores que o valor dele, e todos os valores na sua subárvore à direita devem ser maiores.

* **Linguagem sugerida:** **Python** para entender a lógica; **Rust** para sofrer (e aprender) com o *Borrow Checker* ao manipular referências de nós.
* **Dificuldade:** Fácil / Média.

---

## 2. Percursos (Traversals): A Ordem do Caos

### Como funciona:

Existem quatro formas principais de "visitar" todos os nós de uma árvore binária:

1. **In-Order:** Visita esquerda, depois raiz, depois direita.
2. **Pre-Order:** Visita raiz, depois esquerda, depois direita.
3. **Post-Order:** Visita esquerda, depois direita, depois raiz.
4. **Level-Order (BFS):** Visita os nós nível por nível, de cima para baixo e da esquerda para a direita.

* **Linguagem sugerida:** **Go**. Use *channels* para implementar um percurso que envia valores conforme os encontra.
* **Dificuldade:** Fácil.

---

## 3. O Self-Balancing: Árvore AVL

### Como funciona:

É uma BST com uma regra adicional: a diferença de altura entre a subárvore esquerda e a direita de **qualquer** nó não pode ser maior que 1. Se uma inserção ou deleção quebra essa regra, você deve aplicar "rotações" (simples ou duplas) para reequilibrar a estrutura.

* **Linguagem sugerida:** **Rust**. Implementar rotações e manter a integridade dos ponteiros sem usar `unsafe` é um rito de passagem.
* **Dificuldade:** Difícil.

---

## 4. Binary Heap (Min-Heap / Max-Heap)

### Como funciona:

Uma árvore binária completa onde o valor de cada nó pai é sempre menor ou igual (no caso da Min-Heap) aos valores de seus filhos. Diferente da BST, não há ordem relativa entre os filhos da esquerda e direita. Geralmente, ela é implementada usando um **Array/Slice** em vez de nós conectados por ponteiros, onde a posição do filho é calculada matematicamente via índice.

* **Linguagem sugerida:** **Go**. A manipulação de fatias (slices) e a performance de memória do Go tornam essa implementação muito limpa.
* **Dificuldade:** Média.

---

## 5. Árvores de Busca Digital (Tries)

### Como funciona:

Diferente das árvores binárias, aqui cada nó pode ter múltiplos filhos (geralmente um para cada caractere de um alfabeto). Os nós não armazenam o valor completo; o valor é definido pelo caminho percorrido da raiz até um nó que marca o fim de uma sequência.

* **Linguagem sugerida:** **Python**. A flexibilidade com dicionários permite criar a lógica de ramificação rapidamente.
* **Dificuldade:** Média.

---

## 6. Algoritmos de Grafos: DFS e BFS

### Como funciona:

* **DFS (Depth-First Search):** Explora o máximo possível ao longo de cada ramo antes de retroceder (backtracking). Usa uma **Pilha** (ou recursão).
* **BFS (Breadth-First Search):** Explora todos os vizinhos de um nó antes de passar para os vizinhos do próximo nível. Usa uma **Fila**.
* **Linguagem sugerida:** **Rust** (para DFS recursivo) e **Go** (para BFS com concorrência).
* **Dificuldade:** Média.
