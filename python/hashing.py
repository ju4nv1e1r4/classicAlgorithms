# O Segredo das Colisões no Hashing Trick (neste caso específico)

# "O modelo não aprende o IP, ele aprende o padrão."

# Por que o modelo não quebra com as colisões?

# Em Machine Learning, nós trabalhamos com Sinal vs. Ruído.

# Imagine que o IP de um bot malicioso caia na Vaga 42. Esse bot faz 5.000 requisições por minuto. 
# O contador da Vaga 42 vai explodir. 
# Por coincidência (colisão), o IP de um usuário comum (que acessa 2 vezes ao dia) também cai na mesma Vaga 42.

# Quando o array for passado para a Rede Neural ou para um modelo de Árvore de Decisão, o algoritmo vai olhar para a Vaga 42 e ver 5002 acessos. 
# Ele vai aprender o padrão: "Quando a Vaga 42 tem valores altíssimos, a probabilidade de fraude é de 99%".

# Sim, o usuário comum que teve o azar de colidir naquela vaga pode sofrer um bloqueio injusto (um falso positivo). 
# Mas, do ponto de vista estatístico, o sinal (os 5000 acessos do bot) esmaga o ruído (os 2 acessos do usuário legítimo).

# Como o array tem 100.000 vagas, a probabilidade da colisão destruir o peso de características importantes é baixa o suficiente
# para o modelo abstrair o ruído e continuar altamente preciso, processando gigabytes de logs de texto usando apenas alguns kilobytes de RAM!

import numpy as np

class HashingTrick:
    def __init__(self):
        self.MEM_LIMIT = 100000
        self.counter = np.zeros(
            self.MEM_LIMIT,
            dtype=int,
        )

    def calcular_hash(self, tamanho:int, ip:str) -> int:
        hash_ip = hash(ip)
        return hash_ip % tamanho

    def registrar_ip(self, ip:str):
        hash_ip = self.calcular_hash(self.MEM_LIMIT, ip)
        self.counter[hash_ip] += 1
        return self.counter[hash_ip]
