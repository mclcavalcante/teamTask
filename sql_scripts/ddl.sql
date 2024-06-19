DROP TABLE IF EXISTS Equipe;
DROP TABLE IF EXISTS Comentario;
DROP TABLE IF EXISTS Notificacao;
DROP TABLE IF EXISTS Task_user_associations;
DROP TABLE IF EXISTS User;
DROP TABLE IF EXISTS Tasks;

-- Tabela Usuário
CREATE TABLE User (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(100),
    password VARCHAR(100)
);

-- Tabela Tarefa
CREATE TABLE Tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    status VARCHAR(50),
    priority VARCHAR(50)
);

-- Tabela Equipe
CREATE TABLE Equipe (
    equipe_id INT AUTO_INCREMENT PRIMARY KEY
);

-- Tabela Comentário
CREATE TABLE Comentario (
    comentario_id INT AUTO_INCREMENT PRIMARY KEY,
    texto TEXT,
    tarefa_id INT,
    FOREIGN KEY (tarefa_id) REFERENCES Tasks(id)
);

-- Tabela Notificação
CREATE TABLE Notificacao (
    notificacao_id INT AUTO_INCREMENT PRIMARY KEY,
    tipo VARCHAR(50),
    conteudo TEXT,
    destinatario_id INT,
    FOREIGN KEY (destinatario_id) REFERENCES User(id)
);

-- Tabela de associação entre Usuário e Tarefa (muitos para muitos)
CREATE TABLE Task_user_associations (
    user_id INT,
    task_id INT,
    PRIMARY KEY (user_id, task_id),
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (task_id) REFERENCES Tasks(id)
);
