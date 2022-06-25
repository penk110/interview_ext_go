#!/bin/zsh


#curl --location --request GET '127.0.0.1:8081/api/v1/cache/product_book_2' \
#--header 'Content-Type: application/json' \
#--data-raw '{
#    "key": "product_book_2",
#    "value": "{\"bookName\": \"三体\",\"author\": \"刘慈欣\",\"price\": 99.99}"
#}'


# GET
ab -n 5000 -c 500 http://127.0.0.1:8081/api/v1/cache/product_book_1
