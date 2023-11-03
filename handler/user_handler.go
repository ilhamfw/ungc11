package handler
import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/dgrijalva/jwt-go"
    "time"

    "AvengersCommerce/entity"
    "AvengersCommerce/db"
)

func RegisterUser(c echo.Context) error {
    // Handle user registration logic (simpan data ke database)
    userRequest := new(entity.User)
    if err := c.Bind(userRequest); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Gagal mengikat data pengguna"})
    }

    // Dapatkan koneksi database dari pool
    db, err := db.GetGormDB()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal terhubung ke database"})
    }
    

    // Simpan data pengguna ke database
    if err := db.Create(userRequest).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menyimpan data pengguna ke database"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Pengguna terdaftar"})
}



func LoginUser(c echo.Context) error {

    // Set klaim JWT
    claims := jwt.MapClaims{
        "sub":   "userID",  // Ganti ini dengan ID pengguna yang sesuai
        "exp":   time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 24 jam
    }

    // Buat token JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign token dengan secret key (gantilah dengan kunci yang aman)
    tokenString, err := token.SignedString([]byte("secret"))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuat token JWT"})
    }

    // Kirim token sebagai respons
    return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
