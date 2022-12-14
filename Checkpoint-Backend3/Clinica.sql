CREATE SCHEMA IF NOT EXISTS clinica;
Use clicnica;
 Create	table dentistas (
	id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(250) NOT NULL,
    sobrenome VARCHAR(250) NOT NULL,
    matricula VARCHAR(250) NOT NULL unique,
    
	PRIMARY KEY (id)

 )ENGINE = INNODB;
    
Create	table pacientes (
	id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(250) NOT NULL,
    sobrenome VARCHAR(250) NOT NULL, 
    rg VARCHAR(100) NOT NULL unique,
    datadecadastro DATETIME NOT NULL,
   
	PRIMARY KEY (id)

)ENGINE = INNODB;
 
 Create	table consultas (
	id INT NOT NULL AUTO_INCREMENT,
    dataHora DATETIME NOT NULL,
    paciente VARCHAR(100) NOT NULL,
    descricao VARCHAR(250) NOT NULL,
    dentista VARCHAR(250) NOT NULL,
    
    PRIMARY KEY (id),

    CONSTRAINT fk_dentista
                          FOREIGN KEY (dentista)
                          REFERENCES dentistas(matricula)
                          ON DELETE CASCADE,
    CONSTRAINT fk_paciente
                          FOREIGN KEY (paciente)
                          REFERENCES pacientes(rg)
                          ON DELETE CASCADE
)ENGINE = INNODB;