CREATE TABLE alunos (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    nome VARCHAR(45),
    cpf VARCHAR(11),
    rg VARCHAR(9)
);

INSERT INTO alunos (nome, cpf, rg)
    VALUES ('Bob', '12312312345', '123456789'),
           ('Alice', '12312345674', '112233444'),
           ('Tom', '45621478595', '54545');