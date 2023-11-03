package entity

type User struct {
    ID            int     `json:"id"`
    Username      string  `json:"username"`
    Password      string  `json:"password"`
    DepositAmount float64 `json:"deposit_amount"`
}

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Stock int     `json:"stock"`
    Price float64 `json:"price"`
}

type Transaction struct {
    ID          int     `json:"id"`
    UserID      int     `json:"user_id"`
    ProductID   int     `json:"product_id"`
    StoreID     int     `json:"store_id"`
    Quantity    int     `json:"quantity"`
    TotalAmount float64 `json:"total_amount"`
}

type Store struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Address  string  `json:"address"`
    Longitude float64 `json:"longitude"`
    Latitude  float64 `json:"latitude"`
    Rating    float64 `json:"rating"`
}

type WeatherData struct {
    CloudPct      int     `json:"cloud_pct"`
    Temp          int     `json:"temp"`
    FeelsLike     int     `json:"feels_like"`
    Humidity      int     `json:"humidity"`
    MinTemp       int     `json:"min_temp"`
    MaxTemp       int     `json:"max_temp"`
    WindSpeed     float64 `json:"wind_speed"`
    WindDegrees   int     `json:"wind_degrees"`
    Sunrise       int     `json:"sunrise"`
    Sunset        int     `json:"sunset"`
}



//migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/avengers_ecommerce?sslmode=disable" up
//migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/avengers_ecommerce up
//migrate create -ext sql -dir ./migrations create_avengersCommerse_db
//migrate create -ext sql -dir ./migrations create_avengersCommerse_db

