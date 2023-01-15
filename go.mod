module github.com/RapidCodeLab/AuthService

go 1.19

replace github.com/RapidCodeLab/AuthService/internal/server => ./internal/server

replace github.com/RapidCodeLab/AuthService/pkg/jwt-tokener => ./pkg/jwt-tokener

require (
	github.com/cristalhq/jwt/v4 v4.0.2
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/ilyakaznacheev/cleanenv v1.4.2
	github.com/stretchr/testify v1.8.1
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
