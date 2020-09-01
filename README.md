# Tublessin

Cara Menjalankan Backend
1. Jalankan Microservice
- go run login_service.go
- go run montir_service.go
- go run user_service.go
2. Jalankan Api Gateway
- go run api_gateway.go

Login Service Api List

/account/login/montir
{
    "username": "kucing",
    "password": "admin"
}

/account/login/user
{
    "username": "kucing",
    "password": "admin"
}

/account/register/montir

{
    "username": "cicakakecil",
    "password": "admin",
    "profile": {
        "firstname": "cicak",
        "lastname": "Hitam",
        "gender": "L",
        "city": "Jakarta",
        "email": "kucing@yahoo.com",
        "phone_number": "08982227972"
    }
}
