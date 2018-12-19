package cmd

import (
	"log"

	"github.com/chri5bot/reviews-pya/api"
	"github.com/chri5bot/reviews-pya/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a new reviews REST server",
	Long:  `Start a new reviews REST server`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	config, err := conf.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %+v", err)
	}

	db, err := gorm.Open("postgres", config.DB.URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %+v", err)
	}
	defer db.Close()

	api := api.NewAPI(db, config)

	api.ListenAndServe()

}
