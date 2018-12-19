package cmd

import (
	"log"

	"github.com/chri5bot/reviews-pya/conf"
	migrator "github.com/golang-migrate/migrate"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"

	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func init() {
	rootCmd.AddCommand(migrateCmd, dropCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "db:migrate",
	Short: "Migrate reviews database structure.",
	Long:  "Migrate reviews database structure.",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

var dropCmd = &cobra.Command{
	Use:   "db:drop",
	Short: "Warning: Drop database.",
	Long:  "Warning: Drop database.",
	Run: func(cmd *cobra.Command, args []string) {
		drop()
	},
}

func migrate() {

	config, err := conf.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %+v", err)
		return

	}

	db, err := gorm.Open("postgres", config.DB.URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %+v", err)
	}
	defer db.Close()

	migrator, err := NewMigrator(db)
	if err != nil {
		log.Fatalf("Error opening database: %+v", err)
	}

	log.Println("Migrating...")

	err = migrator.Up()
	if err != nil {
		log.Fatalf("Error migrating database: %+v", err)
	}

	log.Println("Migrate finished.")

}

func drop() {
	config, err := conf.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %+v", err)
		return

	}

	db, err := gorm.Open("postgres", config.DB.URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %+v", err)
	}
	defer db.Close()

	migrator, err := NewMigrator(db)
	if err != nil {
		log.Fatalf("Error opening database: %+v", err)
	}

	log.Println("Dropping...")

	err = migrator.Down()
	if err != nil {
		log.Fatalf("Error dropping database: %+v", err)

	}

	log.Println("Drop finished.")
}

// NewMigrator create a new instance of database with our custom config
func NewMigrator(db *gorm.DB) (*migrator.Migrate, error) {
	driver, err := postgres.WithInstance(db.DB(), &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return migrator.NewWithDatabaseInstance("file://migrations", "postgres", driver)
}
