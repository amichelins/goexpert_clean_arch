# goexpert_clean_arch

#Para rodar o Projeto

A partir da pasta github.com\amichelins\goexpert_clean_arch rodar :

-- Para subir o conteiner
docker-compose up -d

-- Para inicializar a tabela orders
migrate -path=sql/migration -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

------------
--- HTTP ---
------------
  - Porta: 8000
  - Teste:
        Create: /api/create_order.http
        List: /api/list_order.http

------------
--- gRPC ---
------------
evans -r repl

package pb

service OrderService

call ListOrders

---------------
--- Graphql ---
---------------
URL: http://localhost:8080/

//
query Orders {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}