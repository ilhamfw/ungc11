// handler/product_handler.go
package handler

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "AvengersCommerce/entity"
    "AvengersCommerce/db"
)

func GetProducts(c echo.Context) error {
    // Handle fetching products logic
    return c.String(http.StatusOK, "List of products")
}

func PurchaseProduct(c echo.Context) error {
    // Mendapatkan data transaksi dari permintaan
    transactionRequest := new(entity.Transaction)
    if err := c.Bind(transactionRequest); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Gagal mengikat data transaksi"})
    }

    // Dapatkan koneksi database dari pool (gunakan GORM)
    db, err := db.GetGormDB()
    
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal terhubung ke database"})
    }
    

    // Mulai transaksi database
    tx := db.Begin()

    // Periksa stok produk
    product := new(entity.Product)
    if err := tx.Where("id = ?", transactionRequest.ProductID).First(product).Error; err != nil {
        tx.Rollback() // Gulung kembali transaksi jika terjadi kesalahan
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Produk tidak ditemukan"})
    }

    if product.Stock < transactionRequest.Quantity {
        tx.Rollback() // Gulung kembali transaksi jika stok tidak mencukupi
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Stok produk tidak mencukupi"})
    }

    // Periksa saldo pengguna
    user := new(entity.User)
    if err := tx.Where("id = ?", transactionRequest.UserID).First(user).Error; err != nil {
        tx.Rollback() // Gulung kembali transaksi jika terjadi kesalahan
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Pengguna tidak ditemukan"})
    }

    if user.DepositAmount < product.Price*float64(transactionRequest.Quantity) {
        tx.Rollback() // Gulung kembali transaksi jika saldo tidak mencukupi
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Saldo pengguna tidak mencukupi"})
    }

    // Lakukan transaksi (insert data transaksi)
    newTransaction := &entity.Transaction{
        UserID:      transactionRequest.UserID,
        ProductID:   transactionRequest.ProductID,
        Quantity:    transactionRequest.Quantity,
        TotalAmount: product.Price * float64(transactionRequest.Quantity),
    }

    if err := tx.Create(newTransaction).Error; err != nil {
        tx.Rollback() // Gulung kembali transaksi jika terjadi kesalahan
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal melakukan transaksi"})
    }

    // Kurangi stok produk
    product.Stock -= transactionRequest.Quantity
    if err := tx.Save(product).Error; err != nil {
        tx.Rollback() // Gulung kembali transaksi jika terjadi kesalahan
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengurangi stok produk"})
    }

    // Kurangi saldo pengguna
    user.DepositAmount -= product.Price * float64(transactionRequest.Quantity)
    if err := tx.Save(user).Error; err != nil {
        tx.Rollback() // Gulung kembali transaksi jika terjadi kesalahan
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengurangi saldo pengguna"})
    }

    // Commit transaksi jika semuanya berhasil
    tx.Commit()

    return c.JSON(http.StatusOK, map[string]string{"message": "Transaksi berhasil"})
}
