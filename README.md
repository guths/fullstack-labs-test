# FSL Golang Code Challenge

## Makefile

| target | description |
| ----------------- | ------------------------------------------- |
| install           | Install all dependencies |
| prepare           | Install all dependencies and run migrations |
| migrate-up        | Run migrations up |
| migrate-down      | Run migrations down |
| lint              | Run linters |
| start             | Start the application Battle of Monsters |
| test              | Run all application tests |
| migrate-test-up   | Run migrations up in the test database |
| migrate-test-down | Run migrations down in the test database |
| challenge-1       | Run all tests of challenge 1 |
| challenge-2       | Run all tests of challenge 2 |
| challenge-3       | Run all tests of challenge 3 |

## Technologies

- [Viper](https://pkg.go.dev/github.com/spf13/viper#section-readme)
- [GORM](https://pkg.go.dev/gorm.io/gorm)
- [Migrate](https://pkg.go.dev/github.com/golang-migrate/migrate/v4)
- [Cobra](https://pkg.go.dev/github.com/spf13/cobra)
- [GIN](https://pkg.go.dev/github.com/gin-gonic/gin)
- [GinkGo](https://pkg.go.dev/github.com/onsi/ginkgo/v2)
- [Gomega](https://pkg.go.dev/github.com/onsi/gomega)
- [Validator](https://pkg.go.dev/github.com/go-playground/validator)