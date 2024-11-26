package bootstrap

import (
	"github.com/phn00dev/go-blog-temp/internal/database/seeder"
	"github.com/phn00dev/go-blog-temp/pkg/config"
	"github.com/phn00dev/go-blog-temp/pkg/database"
)

func Seeder() {
	config.Set()
	database.Connect()
	seeder.Seed()
}
