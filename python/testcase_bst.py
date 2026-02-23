from bst import Arvore

print("Inicializando a árvore")
minha_arvore = Arvore()

print("--- Inserindo valores ---")
# Vamos criar uma árvore com a raiz 50
valores_para_inserir = [50, 30, 70, 20, 40, 60, 80, 35]
for valor in valores_para_inserir:
    minha_arvore.inserir(valor)
    print(f"Inserido: {valor}")

print("\n--- Árvore em Ordem Crescente ---")
# Se tudo deu certo, veremos do 20 ao 80 em ordem!
minha_arvore.exibir_em_ordem()

print("\n--- Testando a Busca ---")
alvo_1 = 40
resultado_1 = minha_arvore.buscar(alvo_1)
print(f"Buscando {alvo_1}: {'Encontrado' if resultado_1 else 'Não encontrado'}")

alvo_2 = 99
resultado_2 = minha_arvore.buscar(alvo_2)
print(f"Buscando {alvo_2}: {'Encontrado' if resultado_2 else 'Não encontrado'}")

print("\n--- Testando a Remoção ---")
# Cenário 1: Nó folha (ex: 20)
print("Removendo 20 (Nó folha)...")
minha_arvore.remover(20)

# Cenário 2: Nó com um filho (ex: 40 tem o 35)
print("Removendo 40 (Nó com 1 filho)...")
minha_arvore.remover(40)

# Cenário 3: Nó com dois filhos (ex: 50, a própria raiz!)
print("Removendo 50 (Nó com 2 filhos - A Raiz!)...")
minha_arvore.remover(50)

# Para provar que a árvore não quebrou
print("\n--- Árvore Final em Ordem Crescente ---")
minha_arvore.exibir_em_ordem()
