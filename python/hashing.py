# NOTE: apenas o esqueleto do algoritmo, por enquanto

import numpy as np

class Hashing:
    def __init__(self):
        self.MEM_LIMIT = 100000
        self.counter = np.zeros(
            dtype=int,
            ndmax=self.MEM_LIMIT
        )

    def calcular_hash(self, tamanho:int, ip:str) -> int:
        pass

    def registrar_ip(self, ip:str):
        pass
