# auction-app

## Execução
### docker compose up -d
## Exemplos de uso:
#### Criar auction: POST na URL http://localhost:8080/auctions com o body seguinte:
{
    "product_name":"nome do prduto do leilão",
    "category":"categoria",
    "description":"descrição",
    "condition": 1
}
#### Consultar auctions: GET http://localhost:8080/auctions?status=0
#### Criar bid: POST na URL http://localhost:8080/bid com o body seguinte:
{
    "user_id":"user_id",
    "auction_id":"11c202ad-baec-4150-864b-d0d6a97ededa",
    "amount":990.0
}
#### Consultar bids: GET http://localhost:8080/bid/11c202ad-baec-4150-864b-d0d6a97ededa
#### Consultar bid vencedor: GET http://localhost:8080/auctions/winner/11c202ad-baec-4150-864b-d0d6a97ededa
