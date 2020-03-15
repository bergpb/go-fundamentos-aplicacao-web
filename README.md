# go-fundamentos-aplicacao-web
Curso de fundamentos de uma aplicação web na linguagem Go.

Requerimentos: ```go```, ```gin```, ```docker```, ```docker-compose```.

Como executar:
1. Clone o projeto,
1. Entre na pasta do projeto,
1. Execute o comando ```docker-compose up -d```,
1. Acesse o postgres na url: [http://localhost:16543](http://localhost:16543),
    1. Credenciais: email: email@email.com, senha: 12345678
1. Crie uma nova base de dados com o nome: ```alura-loja```.
1. Execute os scripts de criação de tabela e inserção de dados na mesma,
    1. Scripts na pasta ```queries```.
1. Execute com o comando ```gin -i run main.go```,
1. Acesse [http://localhost:3000](http://localhost:3000)
