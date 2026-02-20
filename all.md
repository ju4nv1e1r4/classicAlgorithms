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

---

## Trilha de Algoritmos e Estruturas de Dados para IA e Machine Learning

Este roteiro foca na transição de algoritmos clássicos para aplicações práticas em Inteligência Artificial e Ciência de Dados.

### 1. Fundamentos de Acesso e Organização

* **Tabelas Hash (Dicionários):** Nosso próximo passo. Essencial para entender como o Python gerencia nomes de variáveis e como sistemas de "Feature Stores" buscam dados de usuários em tempo real.
* **Heaps / Filas de Prioridade:** Fundamentais para algoritmos de busca como o A* (usado em robótica e pathfinding) e para encontrar os "Top K" elementos em grandes conjuntos de dados (como os resultados de uma busca no Google).

### 2. Estruturas para Dados em Alta Dimensão

* **Tries (Árvores de Prefixos):** Como o corretor ortográfico ou o "Autocomplete" de uma IA funciona tão rápido.
* **KD-Trees & Ball Trees:** Versões avançadas da sua Árvore Binária usadas para encontrar os vizinhos mais próximos (K-Nearest Neighbors) em espaços multidimensionais. Sem isso, sistemas de recomendação seriam impossíveis.

### 3. Algoritmos de Grafos (O Coração das Redes Neurais)

* **DAGs (Grafos Acíclicos Dirigidos):** Representação de como o fluxo de cálculo passa por uma rede neural (Backpropagation).
* **Algoritmos de Menor Caminho (Dijkstra/A*):** Essenciais para IA em jogos e logística.
* **Grafos de Conhecimento (Knowledge Graphs):** Como as IAs modernas conectam fatos (ex: "Paris" -> "Capital" -> "França").

### 4. Otimização e Programação Dinâmica

* **Programação Dinâmica (DP):** Usada em algoritmos de processamento de linguagem natural (NLP) antigos e em Aprendizado por Reforço (Reinforcement Learning).
* **Gradiente Descendente (Versão Algorítmica):** Entender a matemática por trás de como um modelo "aprende" ajustando pesos.

### 5. Algoritmos de Amostragem e Probabilidade

* **Reservoir Sampling:** Como escolher uma amostra aleatória de uma base de dados tão grande que não cabe na memória (Big Data).
* **Filtros de Bloom:** Como saber se um dado NÃO existe em um banco de dados gigante sem precisar consultar o disco (economiza bilhões de dólares em infraestrutura).

### Por que isso importa para um AI Scientist?

* **Eficiência de Memória:** Treinar modelos grandes exige saber exatamente como os dados são alocados.
* **Performance de Inferência:** Colocar um modelo em produção exige que ele responda rápido. Um algoritmo $O(n^2)$ pode matar o seu produto, enquanto um $O(n \log n)$ o salva.
* **Customização:** Muitas vezes, as bibliotecas prontas (PyTorch, TensorFlow) não resolvem um problema específico, e você precisará codar a lógica de otimização na mão.