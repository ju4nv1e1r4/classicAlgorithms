pub struct Node {
    pub valor: i32,
    pub esquerda: Option<Box<Node>>,
    pub direita: Option<Box<Node>>,
}

impl Node {
    pub fn new(valor: i32) -> Self {
        Self { valor, esquerda: None, direita: None }
    }

    pub fn inserir(&mut self, valor: i32) {
        if valor < self.valor {
            if let Some(filho) = &mut self.esquerda {
                filho.inserir(valor);   
            } else {
                self.esquerda = Some(Box::new(Node::new(valor)));
            }
        } else if valor > self.valor {  
            if let Some(filho) = &mut self.direita {
                filho.inserir(valor);
            } else {
                self.direita = Some(Box::new(Node::new(valor)));
            }
        }
    }

    pub fn exibir_em_ordem(&self) {
        if let Some(filho) = &self.esquerda {
            filho.exibir_em_ordem();
        }

        println!("{}", self.valor);

        if let Some(filho) = &self.direita {
            filho.exibir_em_ordem();
        }
    }

    pub fn buscar(&self, valor: i32) -> bool {
        if self.valor == valor {
            return true;
        }

        if valor < self.valor {
            if let Some(filho) = &self.esquerda {
                return filho.buscar(valor);
            }
        } else if valor > self.valor {
            if let Some(filho) = &self.direita {
                return filho.buscar(valor);
            }
        }

        false
    }

    pub fn remover(mut no_atual: Box<Self>, valor: i32) -> Option<Box<Self>>{
        if valor < no_atual.valor {
            if let Some(filho_esquerda) = no_atual.esquerda.take() {
                no_atual.esquerda = Node::remover(filho_esquerda, valor);
            }
            return Some(no_atual);

        } else if valor > no_atual.valor {
            if let Some(filho_direita) = no_atual.direita.take() {
                no_atual.direita = Node::remover(filho_direita, valor);
            }

            return Some(no_atual);
        } else {
            if no_atual.esquerda.is_none() {
                return no_atual.direita.take();
            }
            if no_atual.direita.is_none() {
                return no_atual.esquerda.take();
            }
            
            let minimo = no_atual.direita.as_ref().unwrap().encontrar_minimo();
            no_atual.valor = minimo;
    
            let filho_direita = no_atual.direita.take().unwrap();
            no_atual.direita = Node::remover(filho_direita, minimo);
    
            return Some(no_atual);
        }

        }
    pub fn encontrar_minimo(&self) -> i32 {
        if let Some(filho) = &self.esquerda {
            filho.encontrar_minimo()
        } else {
            self.valor
        }
    }
}
