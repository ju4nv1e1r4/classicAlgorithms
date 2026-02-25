mod bst;
use bst::Node;

fn main() {
    println!("Iniciando a árvore binária de busca.\n");

    let mut raiz = Some(Box::new(Node::new(50)));
    let raiz_ref = raiz.as_mut().unwrap();

    println!("Inserindo valores.\n");

    raiz_ref.inserir(30);
    raiz_ref.inserir(20);
    raiz_ref.inserir(40);
    raiz_ref.inserir(70);
    raiz_ref.inserir(60);
    raiz_ref.inserir(80);

     print!("Árvore em Ordem: [ ");
    raiz.as_ref().unwrap().exibir_em_ordem();
    println!("]\n");

    let busca_40 = raiz.as_ref().unwrap().buscar(40);
    let busca_99 = raiz.as_ref().unwrap().buscar(99);
    println!("Busca pelo 40 (Existe?): {}", busca_40);
    println!("Busca pelo 99 (Existe?): {}\n", busca_99);

    println!("Removendo o nó 50 (raiz com dois filhos).\n");

     raiz = Node::remover(raiz.take().unwrap(), 50);

    print!("Árvore após remoção: [ ");
    raiz.as_ref().unwrap().exibir_em_ordem();
    println!("]\n");
}
