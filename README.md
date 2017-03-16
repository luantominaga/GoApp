### Golang MySQL app (Agente Brasil)

#### Requisitos:

Instalar o GOLang no sistema:

1. Seguir os passos do site

https://golang.org/doc/install 

2. Requisitos de importação

import (
	"database/sql"
	"net/http"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


### Como executar 

1. Execute o terminal do linux.

1.1 Instale o mysql no sistema: sudo apt-get install mysql-server

1.2 Execute: mysql -h localhost -u root -p

1.3 Crie um database: CREATE DATABASE database_name;

1.4 Use a database: USE database_name;



2. Executar o aplicativo golang.

2.1 Abra o terminal e vá até a pasta que se encontra o aplicativo app.go (use comando cd para ir a pasta).

2.2 Execute o aplicativo com o comando: go run app.go


3. Página html.

3.1 Abra um navegador e digite: http://localhost:8080
(imagem1)

3.2 Crie uma tabela no menu "Cria/apagar tabela.

3.3 Digite no teminal onde está executando o mysql e digite: SELECT * FROM users;
(imagem2)

3.4 A tabela foi criada. Após, voltar a página web e entrer no menu "Entrar no formulário". Cadastre o e aperto o botão "adicionar"

3.5 Para verificar se inseriu os dados, vá ao terminal onde está executando o mysql e digite: SELECT * FROM users;

3.6 A tabela aparecerá com valores.

4. Finalizar o servidor

4.1 Para finalizar o aplicativo, digitar no terminal: killall app



Dentro do app.go **app.go** linha **130** colocar no <exemplo> as suas próprias credênciais.

```go
db, err = sql.Open("mysql", "<root>:<password>@/<dbname>")
// Coloque com 
db, err = sql.Open("mysql", "myUsername:myPassword@/myDatabase")
```
No meu caso, criei uma database, usei como root e uma database chamado webapp: 	db, err = sql.Open("mysql", "root: @/webapp")


Obrigado.
