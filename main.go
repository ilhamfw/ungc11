package main

import (
    "log"
    "github.com/labstack/echo/v4"
    _ "github.com/lib/pq"
    "AvengersCommerce/db"
    "AvengersCommerce/handler" 
)

func main() {
    e := echo.New()
    _, err := db.GetGormDB()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    

    
    // Rute untuk registrasi pengguna
    e.POST("/users/register", handler.RegisterUser)

    // Rute untuk login pengguna
    e.POST("/users/login", handler.LoginUser)

    // Rute untuk mendapatkan daftar produk
    e.GET("/products", handler.GetProducts)

    // Rute untuk melakukan transaksi
    e.POST("/transactions", handler.PurchaseProduct)

    // Rute untuk mendapatkan daftar toko
    e.GET("/stores", handler.GetStores)

    // Rute untuk mendapatkan detail toko berdasarkan ID
    e.GET("/stores/:id", handler.GetStoreDetail)

    e.Start(":8080")
}
