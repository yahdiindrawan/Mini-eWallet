# Mini-eWallet
Backend e-Wallet using Go

<h3>Installations</h3>

1. Clone this project to your local machine<br>
`http://Github.com/yahdiindrawan/Mini-eWallet.git`

2. Install packages<br>
`$ go get github.com/gorilla/mux`<br>
`$ go get github.com/jinzhu/gorm`<br>
And import
`_ "github.com/jinzhu/gorm/dialects/mysql"`

3. Setup Database<br>
`$ DB_HOST = 127.0.0.1`<br>
`$ DB_NAME = "e-wallet"`<br>
Use `http://localhost:9999/` as base url for endpoints

4. Run Project<br>
`$ go run main.go` in the root directory<br>

<h3>Schema Database</h3>

![schema database e-wallet](https://user-images.githubusercontent.com/55028341/100875294-68c12200-34d8-11eb-8028-bb889410ed72.png)



<h3>Features</h3>

1. CRUD Rest API from the schema<br>
`http://localhost:9999/api/users` <br>
`http://localhost:9999/api/userbalances` <br>
`http://localhost:9999/api/userbalancehistories` <br>
`http://localhost:9999/api/blancebanks` <br>
`http://localhost:9999/api/blancebankhistories` <br>

2. Authentication<br>
`http://localhost:9999/login` <br>
`http://localhost:9999/register` <br>

3. TopUp Balance<br>
`http://localhost:9999/topup` <br>

4. Transfer Balance<br>
`http://localhost:9999/transfer` <br>

