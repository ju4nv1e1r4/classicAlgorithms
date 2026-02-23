package hashing

import "hash/fnv"

type HashingTrick struct {
	Size int32
	Counter []int32
}

func NovoHashingTrick(tamanho int32) *HashingTrick {
	return &HashingTrick{
		Size: tamanho,
		Counter: make([]int32, tamanho),
	}
}

func calcularHash(tamanho int32, ip string) uint32 {
	hashIP := fnv.New32a()
	hashIP.Write([]byte(ip))
	return hashIP.Sum32() % uint32(tamanho)
}

func RegistrarIP(hashingTrick *HashingTrick, ip string) int32 {
	hashIP := calcularHash(hashingTrick.Size, ip)
	indice := int(hashIP)
	hashingTrick.Counter[indice]++
	return hashingTrick.Counter[indice]
}
