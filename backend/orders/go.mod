module github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders

go 1.23.6

require (
	github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader v0.0.0-20241112163550-b21cd8c22af6
	github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto v0.0.0
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gofiber/fiber/v3 v3.0.0-beta.3
	github.com/golang-jwt/jwt v3.2.2+incompatible
)

replace github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto => ../../proto

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway v0.0.0-20241112163550-b21cd8c22af6 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/text v0.19.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/gofiber/utils/v2 v2.0.0-beta.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.55.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)
