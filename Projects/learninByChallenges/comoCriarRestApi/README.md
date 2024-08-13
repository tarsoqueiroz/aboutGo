# Como criar uma REST API completa do zero com GO

Golang tutorial - iniciante

> `https://www.youtube.com/watch?v=3p4mpId_ZU8`

## Prereq

- **Golang extension for VSCode**

Go (v0.42.0): Go Team at Google

Rich Go language support for Visual Studio Code

## Starting the project

```sh
go mod init go-api

mkdir cmd
touch cmd/main.go
```

- [main.go](./backup/main.v01.go)

```sh
go get github.com/gin-gonic/gin
```

- [main.go](./backup/main.v02.go)

```sh
go run cmd/main.go
```

- (1) [test.http](./backup/test.v01.http)

## Dockerizing Postgres

- [docker-compose.yaml](./backup/docker-compose.v01.yaml)

```sh
docker compose up -d go_db
```

## DBeaver

```sh
sudo snap install dbeaver-ce
```

Connect to Postgres and create a table:

```sql
create table product (
  id SERIAL primary key,
  product_name VARCHAR(50) not null,
  price NUMERIC(10, 2) not NULL
);
```

Try to access the table:

```sql
select * from product;
```

Inserting some values:

```sql
insert into product(product_name, price) values('Sushi', 100);
```

Try again to access the table:

```sql
select * from product;
```

## Creating "/products" route

```sh
mkdir model
touch model/product.go
```

- [product.go](./backup/product.v01.go)

```sh
mkdir controller
touch controller/product_controller.go
```

- [product_controller.go](./backup/product_controller.v01.go)

- [main.go](./backup/main.v03.go)

```sh
go run cmd/main.go
```

- (2) [test.http](./backup/test.v01.http)

...

> I'm in 17:37

...
