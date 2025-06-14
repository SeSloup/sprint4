module github.com/Yandex-Practicum/tracker

go 1.24.3

require github.com/stretchr/testify v1.10.0

require (
	github.com/Yandex-Practicum/tracker/internal/daysteps v0.0.0-00010101000000-000000000000 // indirect
	github.com/Yandex-Practicum/tracker/internal/spentcalories v0.0.0-00010101000000-000000000000 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect

)

replace github.com/Yandex-Practicum/tracker/internal/daysteps => ./internal/daysteps

replace github.com/Yandex-Practicum/tracker/internal/spentcalories => ./internal/spentcalories
