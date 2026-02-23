import numpy as np
import sys
import random
import time

from hashing import HashingTrick

def medir_memoria_dicionario(dicionario):
    """
    Função auxiliar para medir o tamanho real de um dicionário e seu conteúdo
    """
    tamanho_bytes = sys.getsizeof(dicionario)
    for chave, valor in dicionario.items():
        tamanho_bytes += sys.getsizeof(chave)
        tamanho_bytes += sys.getsizeof(valor)
    return tamanho_bytes

def main():
    num_requisicoes = 1_000_000
    limite_hashing_trick = 100_000

    print(f"Gerando {num_requisicoes:,} requisições de IPs aleatórios...")

    ips_simulados = [
        f"{random.randint(1, 255)}.{random.randint(0, 255)}.{random.randint(0, 20)}.{random.randint(0, 50)}" 
        for _ in range(num_requisicoes)
    ]
    
    print("Dados gerados! Iniciando...\n")

    inicio_tempo = time.time()
    dicionario_tradicional = {}
    
    for ip in ips_simulados:
        if ip in dicionario_tradicional:
            dicionario_tradicional[ip] += 1
        else:
            dicionario_tradicional[ip] = 1
            
    tempo_dict = time.time() - inicio_tempo
    memoria_dict_bytes = medir_memoria_dicionario(dicionario_tradicional)
    memoria_dict_mb = memoria_dict_bytes / (1024 * 1024)

    inicio_tempo = time.time()
    hashing_trick = HashingTrick(limite_memoria=limite_hashing_trick)
    
    for ip in ips_simulados:
        hashing_trick.registrar_ip(ip)
        
    tempo_ht = time.time() - inicio_tempo
    # nbytes do numpy dá o peso exato do array na memória
    memoria_ht_bytes = hashing_trick.counter.nbytes
    memoria_ht_mb = memoria_ht_bytes / (1024 * 1024)

    print("-" * 50)
    print(f"RESULTADOS PARA {num_requisicoes:,} REQUISIÇÕES")
    print("-" * 50)
    print(f"Dicionário Tradicional:")
    print(f" -> IPs Únicos armazenados: {len(dicionario_tradicional):,}")
    print(f" -> Memória Consumida:      {memoria_dict_mb:.2f} MB")
    print(f" -> Tempo de execução:      {tempo_dict:.3f} segundos\n")

    print(f"Hashing Trick (Array de {limite_hashing_trick:,} posições):")
    print(f" -> Memória Consumida:      {memoria_ht_mb:.2f} MB")
    print(f" -> Tempo de execução:      {tempo_ht:.3f} segundos\n")
    
    economia = memoria_dict_mb / memoria_ht_mb
    print(f"CONCLUSÃO: O Hashing Trick usou {economia:.1f}x MENOS memória!")
    print("-" * 50)

if __name__ == "__main__":
    main()
