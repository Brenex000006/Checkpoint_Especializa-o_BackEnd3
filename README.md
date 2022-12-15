# Checkpoint_BackEnd3

Especialização em Back End III
Workshop de código: Desafio final
Objetivo
A seguir, apresenta-se um desafio integrador que nos permitirá avaliar todos os tópicos abordados no curso. 

Sistema de marcação de consultas:

Pretende-se implementar uma API que permita administrar a marcação de consultas para uma clínica odontológica, que deve atender aos requerimentos a seguir: 
Administração de dados de dentistas: listar, adicionar, alterar ou excluir dentistas. Registrar os seus sobrenome, nome e matrícula. Solicita-se o desenvolvimento de um CRU para a entidade Dentista.  

POST: adicionar dentista. 
GET: trazer dentista pelo seu ID.
PUT: atualizar dentista.
PATCH: atualizar um dentista através de algum dos seus campos. 
DELETE: excluir dentista.

Administração de dados dos pacientes: listar, adicionar, alterar ou excluir pacientes. Registrar os seus sobrenome, RG, nome e data de cadastro. Solicita-se o desenvolvimento de um CRU para a entidade Paciente.  

POST: adicionar dentista.

GET: trazer paciente pelo seu ID.
PUT: atualizar paciente.
PATCH: atualizar um paciente através de algum dos seus campos. 
DELETE: excluir paciente.


Marcação de consulta: deve ser possível atribuir a um paciente uma consulta com um dentista em uma determinada data e hora, e também adicionar uma descrição à consulta. Solicita-se o desenvolvimento de um CRUD para a entidade Consulta.  

POST: adicionar consulta
GET: trazer consulta pelo seu ID.
PUT: atualizar consulta.
PATCH: atualizar consulta por algum dos seus campos.
DELETE: excluir consulta.
POST: adicionar consulta pelo RG do paciente e matrícula do dentista.
GET: trazer consulta pelo RG do paciente. Deve conter o detalhamento da consulta (Data-Hora, descrição, Paciente e Dentista). 

Requerimentos técnicos:

A aplicação deve ser desenvolvida em design orientado a pacotes:
Camada/domínio de entidades de negócio.
Camada/domínio de acesso a dados (Repository).
Camada de acesso a dados (banco de dados): é o banco de dados do nosso sistema. Você poderá usar qualquer um dos bancos de dados relacionais modelado através de um modelo entidade-relação, como H2 ou MySQL, ou não relacional, como Mongo DB.
Camada/domínio service.
Camada/domínio handler.

Boa sorte!
