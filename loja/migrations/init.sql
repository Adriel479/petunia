CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(80) NOT NULL,
    descricao VARCHAR(150) NOT NULL,
    preco DECIMAL(10,2) NOT NULL,
    quantidade INTEGER NOT NULL
);

INSERT INTO produtos (nome, descricao, preco, quantidade)
    VALUES ('Camisata', 'Azul tamanho M', 25.00, 100),
           ('Tênis', 'Preto tamanho 40', 70, 50),
           ('Notebook', 'Notebook Geek 12', 3542.10, 100),
           ('Xbox Series S', 'Jogos Digitais', 100.15, 10);