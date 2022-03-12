module account_service

replace config => ../../pkg/config

replace log => ../../pkg/log
replace error_code => ../../error_code

go 1.12

require (
	config v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/rhinoceros100/trainingCamp v0.0.0-20220312021848-60fe8653d358
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	google.golang.org/grpc v1.45.0 // indirect
)
