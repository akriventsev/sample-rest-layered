# Sample app

# Start

make build_sample_local

make create_sample_local


POST http://localhost:8081/signup

```
request
{
    "first_name": "username",
    "last_name": "boom",
    "married": true,
    "login": "user",
    "password": "mypass",
    "birthday": "2003-12-07T00:00:00Z"
}

response

{
    "id": "c29912dc-e3c3-4544-a2e8-186ca3ed9afc",
    "first_name": "username",
    "last_name": "boom",
    "married": true,
    "login": "user",
    "birthday": "2003-12-07T00:00:00Z"
}
```

POST http://localhost:8081/login

form-data username and password

```
response

{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidXNlcm5hbWUgYm9vbSIsImlkIjoiYzI5OTEyZGMtZTNjMy00NTQ0LWEyZTgtMTg2Y2EzZWQ5YWZjIiwiZXhwIjoxNzAxOTMyODUzfQ.YhbInoQZb21vTbGU60gAYjG3omysD-1cYl_Q6NPEYeI"
}
```


POST http://localhost:8081/order 

```
request
{
    "items":[
        {
            "product_id":"69dad391-2968-43ce-a751-bc52b4f5d3d7",
            "quantity": 1
        },
        {
            "product_id":"58c4271a-6c40-499f-bc2e-44524f37fa0d",
            "quantity": 1
        }
    ]
}

response

{
    "id": "9500b83b-2240-4d43-8f97-cf3203793258",
    "items": [
        {
            "product_id": "69dad391-2968-43ce-a751-bc52b4f5d3d7",
            "quantity": 1,
            "price": 12000
        },
        {
            "product_id": "58c4271a-6c40-499f-bc2e-44524f37fa0d",
            "quantity": 1,
            "price": 15000
        }
    ]
}

```