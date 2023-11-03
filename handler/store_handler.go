package handler

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "io"
    "fmt"
	"encoding/json"

	"AvengersCommerce/entity"
	"AvengersCommerce/db"
)

func GetStores(c echo.Context) error {
    // Inisialisasi koneksi database
    gormDB, err := db.GetGormDB()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal terhubung ke database"})
    }

    // Deklarasi variabel untuk menyimpan daftar toko
    var stores []entity.Store

    // Mengambil daftar toko dari database
    if err := gormDB.Find(&stores).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengambil daftar toko"})
    }

    // Mengembalikan daftar toko dalam respons JSON
    return c.JSON(http.StatusOK, stores)
}

// GetStoreDetail mengambil detail dari toko berdasarkan ID toko yang diberikan dalam permintaan.
func GetStoreDetail(c echo.Context) error {
    storeID := c.Param("id") // Ambil storeID dari parameter URL

    // Ambil koneksi GORM DB
    db, err := db.GetGormDB()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal terhubung ke database"})
    }

    store := entity.Store{} // Deklarasi variabel untuk menyimpan data toko
    result := db.First(&store, storeID) // Mengambil data toko dari database berdasarkan storeID

    if result.Error != nil {
        return c.JSON(http.StatusOK, map[string]string{"message": "Toko tidak ditemukan"})
    }

    // Mengambil data cuaca	
    weatherURL := "https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?city=Seattle"
    weatherReq, _ := http.NewRequest("GET", weatherURL, nil)
    weatherReq.Header.Add("X-RapidAPI-Key", "196e675eb7msh21f740732fecadfp18fc8ajsn4901691e09c4") 
    weatherReq.Header.Add("X-RapidAPI-Host", "weather-by-api-ninjas.p.rapidapi.com")
    weatherRes, _ := http.DefaultClient.Do(weatherReq)
    
    weatherBody, _ := io.ReadAll(weatherRes.Body)
    fmt.Println(string(weatherBody))

    // Membaca data cuaca ke dalam struct WeatherData
    weatherData := entity.WeatherData{}
    if err := json.Unmarshal(weatherBody, &weatherData); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membaca data cuaca"})
    }

    // Mengembalikan detail toko beserta data cuaca
    storeWithWeather := map[string]interface{}{
        "store":  store,
        "weather": weatherData,
    }

    return c.JSON(http.StatusOK, storeWithWeather)
}
