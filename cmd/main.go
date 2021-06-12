package main

import (
	"github.com/carrot-systems/carrot-automation/src/adapters/gateway/jwt"
	"github.com/carrot-systems/carrot-automation/src/adapters/gateway/services"
	"github.com/carrot-systems/carrot-automation/src/adapters/persistence/postgres"
	"github.com/carrot-systems/carrot-automation/src/adapters/rest"
	"github.com/carrot-systems/carrot-automation/src/config"
	"github.com/carrot-systems/carrot-automation/src/core/usecases"
	"github.com/heroku/rollrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func setupLogs(logConfig config.LogConfig) {
	rollrus.SetupLogging(logConfig.RollbarToken, "development")
}

func main() {
	config.LoadEnv()

	log.Info("starting carrot-automation...")
	setupLogs(config.LoadLogConfiguration())

	ginConfiguration := config.LoadGinConfiguration()
	dbConfig := config.LoadGormConfiguration()
	jwtConfig := config.LoadJwtConfiguration()

	jwtInstance := jwt.LoadJwt(jwtConfig.Secret)

	services := services.CreateServiceManager()
	var workflowRepo usecases.WorkflowRepo

	var db *gorm.DB
	if dbConfig.Engine == "POSTGRES" {
		db = postgres.StartGormDatabase(dbConfig)
		workflowRepo = postgres.NewWorkflowRepo(db)

		db.AutoMigrate(&postgres.Workflow{})
		db.AutoMigrate(&postgres.RunningHistory{})
		db.AutoMigrate(&postgres.Action{})
		db.AutoMigrate(&postgres.ActionVariable{})
		/*err := postgres.Migrate(db, "./migrations", "carrot_automation_migration")
		if err != nil {
			log.Fatalln(err.Error())
		}*/
		_ = db
	}

	usecasesHandler := usecases.NewInteractor(jwtInstance, &services, workflowRepo)

	restServer := rest.NewServer(ginConfiguration)
	routesHandler := rest.NewRouter(usecasesHandler)

	rest.SetRoutes(restServer.Router, routesHandler)
	restServer.Start()
}
