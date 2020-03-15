# go-fundamentos-aplicacao-web
Curso de fundamentos de uma aplicação web na linguagem Go.

Requerimentos: ```go```, ```docker```, ```docker-compose```.

Como executar:
1. Clone o projeto,
2. Entre na pasta do projeto,
3. Instale as dependências com o comando: ```go mod download```,
4. Execute o comando ```docker-compose up -d```,
5. Acesse o postgres na url: [http://localhost:16543](http://localhost:16543),
    * Credenciais acesso: email: email@email.com, senha: 12345678,
    * Credenciais postgres: usuário: postgres, senha: 12345678,
6. Crie uma nova base de dados com o nome: ```alura-loja```.
7. Execute os scripts de criação de tabela e inserção de dados de teste,
    * Scripts na pasta ```queries```.
8. Execute com o comando ```gin -i run main.go```,
9. Acesse [http://localhost:3000](http://localhost:3000)
