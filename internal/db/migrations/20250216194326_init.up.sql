CREATE TABLE petroleira (
    id VARCHAR(36) PRIMARY KEY,
    organizacao VARCHAR(255) NOT NULL,
    nome VARCHAR(255) NOT NULL,
    descricao LONGTEXT,
    tags LONGTEXT NOT NULL,
    qtdRecursos INTEGER,
    qtdReusos INTEGER,
    qtdDownloads INTEGER,
    qtdSeguidores INTEGER
);
