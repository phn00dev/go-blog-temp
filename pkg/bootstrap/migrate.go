package bootstrap

import (
	"github.com/phn00dev/go-blog-temp/internal/database/migration"
	"github.com/phn00dev/go-blog-temp/pkg/config"
	"github.com/phn00dev/go-blog-temp/pkg/database"
)

func Migrate() {
	config.Set()
	// database connect
	database.Connect()
	migration.Migrate()
}
