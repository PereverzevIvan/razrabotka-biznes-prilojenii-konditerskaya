module github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway

go 1.22.7

replace github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader => ./pkg/config_loader

require github.com/gofiber/fiber/v3 v3.0.0-beta.3

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/text v0.19.0 // indirect
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/pkg/config_loader v0.0.0-20241029204316-6ba74e7ea086 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gofiber/utils/v2 v2.0.0-beta.4 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang-migrate/migrate/v4 v4.18.1
	github.com/google/uuid v1.6.0 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.55.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.28.0
	golang.org/x/sys v0.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
