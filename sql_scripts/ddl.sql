-- Tabela Usuário
DROP TABLE IF EXISTS Usuario;
CREATE TABLE Usuario (
    usuario_id INT PRIMARY KEY,
    nome VARCHAR(255),
    cargo VARCHAR(100)
);

-- Tabela Tarefa
DROP TABLE IF EXISTS Tarefa;
CREATE TABLE Tarefa (
    tarefa_id INT PRIMARY KEY,
    titulo VARCHAR(255),
    descricao TEXT,
    status VARCHAR(50),
    prioridade INT,
);

-- Tabela Equipe
DROP TABLE IF EXISTS Equipe;
CREATE TABLE Equipe (
    equipe_id INT PRIMARY KEY
);

-- Tabela Comentário
DROP TABLE IF EXISTS Comentario;
CREATE TABLE Comentario (
    comentario_id INT PRIMARY KEY,
    texto TEXT,
    tarefa_id INT,
    FOREIGN KEY (tarefa_id) REFERENCES Tarefa(tarefa_id)
);

-- Tabela Notificação
DROP TABLE IF EXISTS Notificacao;
CREATE TABLE Notificacao (
    notificacao_id INT PRIMARY KEY,
    tipo VARCHAR(50),
    conteudo TEXT,
    destinatario_id INT,
    FOREIGN KEY (destinatario_id) REFERENCES Usuario(usuario_id)
);

-- Tabela de associação entre Usuário e Tarefa (muitos para muitos)
DROP TABLE IF EXISTS Usuario_Tarefa;
CREATE TABLE Usuario_Tarefa (
    usuario_id INT,
    tarefa_id INT,
    PRIMARY KEY (usuario_id, tarefa_id),
    FOREIGN KEY (usuario_id) REFERENCES Usuario(usuario_id),
    FOREIGN KEY (tarefa_id) REFERENCES Tarefa(tarefa_id)
);
