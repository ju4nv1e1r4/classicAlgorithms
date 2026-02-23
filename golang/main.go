package main

import (
	"classicAlgorithms/hashing"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func getMemAlloc() uint64 {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func main() {
	numRequisicoes := 1_000_000
	limiteMemoria := int32(100_000)

	fmt.Printf("Gerando %d requisições de IPs aleatórios...\n", numRequisicoes)
	ips := make([]string, numRequisicoes)
	for i := 0; i < numRequisicoes; i++ {
		ips[i] = fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(20), rand.Intn(50))
	}
	fmt.Println("Dados gerados! Iniciando...")

	memAntesMapa := getMemAlloc()
	inicioMapa := time.Now()

	mapaTradicional := make(map[string]int32)
	for _, ip := range ips {
		mapaTradicional[ip]++
	}

	tempoMapa := time.Since(inicioMapa)
	memDepoisMapa := getMemAlloc()
	memoriaMapaMB := float64(memDepoisMapa-memAntesMapa) / 1024 / 1024

	runtime.KeepAlive(mapaTradicional)

	mapaTradicional = nil
	runtime.GC()

	inicioHT := time.Now()

	ht := hashing.NovoHashingTrick(limiteMemoria)
	for _, ip := range ips {
		hashing.RegistrarIP(ht, ip)
	}

	tempoHT := time.Since(inicioHT)

	tamanhoTeoricoMB := float64(limiteMemoria*4) / 1024 / 1024

	fmt.Println(stringsRepeat("-", 50))
	fmt.Printf("RESULTADOS PARA %d REQUISIÇÕES\n", numRequisicoes)
	fmt.Println(stringsRepeat("-", 50))

	fmt.Println("Map Tradicional:")
	fmt.Printf(" -> Memória Consumida: ~%.2f MB\n", memoriaMapaMB)
	fmt.Printf(" -> Tempo de execução: %s\n\n", tempoMapa)

	fmt.Printf("Hashing Trick (Array de %d posições):\n", limiteMemoria)
	fmt.Printf(" -> Memória Consumida: ~%.2f MB (Exatos e imutáveis!)\n", tamanhoTeoricoMB)
	fmt.Printf(" -> Tempo de execução: %s\n\n", tempoHT)

	economia := memoriaMapaMB / tamanhoTeoricoMB
	fmt.Printf("CONCLUSÃO: O Hashing Trick usou %.1fx MENOS memória!\n", economia)
	fmt.Println(stringsRepeat("-", 50))
}

// auxiliar para repetir strings
func stringsRepeat(s string, count int) string {
	res := ""
	for i := 0; i < count; i++ {
		res += s
	}
	return res
}

// import (
// 	"classicAlgorithms/bst"
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Iniciando a ávore")

// 	fmt.Println("--- Inserindo valores ---")
// 	// Vamos criar uma árvore com a raiz 50
// 	valoresParaInserir := [...]int{50, 30, 70, 20, 40, 60, 80, 35}

// 	var n *bst.Node

// 	for _, valor := range valoresParaInserir {
// 		n = n.Inserir(valor)
// 		fmt.Println("Inserido: ", valor)
// 	}

// 	fmt.Println("\n--- Árvore em Ordem Crescente ---")
// 	// Se tudo deu certo, veremos do 20 ao 80 em ordem!
// 	n.ExibirEmOrdem()

// 	fmt.Println("\n--- Testando a Busca ---")
// 	alvo1 := 40
// 	resultado1 := n.Buscar(alvo1)
// 	if resultado1 {
// 		fmt.Printf("O número %d foi encontrado.", alvo1)
// 	} else {
// 		fmt.Printf("O número %d não foi encontrado.", alvo1)
// 	}

// 	fmt.Println("\n--- Testando a Busca ---")
// 	alvo2 := 99
// 	resultado2 := n.Buscar(alvo2)
// 	if resultado2 {
// 		fmt.Printf("O número %d foi encontrado.", alvo2)
// 	} else {
// 		fmt.Printf("O número %d não foi encontrado.", alvo2)
// 	}

// 	fmt.Println("\n--- Testando a Remoção ---")
// 	fmt.Println("Removendo 20 (Nó folha)...")
// 	n = n.Remover(20)

// 	fmt.Println("\n--- Removendo 40 (Nó com 1 filho) ---")
// 	fmt.Println("Removendo 40 (Nó folha)...")
// 	n = n.Remover(40)

// 	fmt.Println("\n--- Removendo 50 (Nó com 2 filhos - A Raiz!) ---")
// 	fmt.Println("Removendo 50 (Nó folha)...")
// 	n = n.Remover(50)

// 	fmt.Println("\n--- Árvore Final em Ordem Crescente ---")
// 	n.ExibirEmOrdem()
// }
