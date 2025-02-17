-- name: GetAllData :many
SELECT * FROM petroleira;

-- name: CreateRecord :execresult
INSERT INTO petroleira ( 
    id, 
    organizacao, 
    nome, 
    descricao, 
    tags, 
    qtdRecursos, 
    qtdReusos, 
    qtdDownloads, 
    qtdSeguidores )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);