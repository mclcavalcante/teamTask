package main

import (
	"database/sql"
	"os"

	// "os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mclcavalcante/teamTask/config"
	"github.com/mclcavalcante/teamTask/controller"
	"github.com/mclcavalcante/teamTask/router"
	"github.com/mclcavalcante/teamTask/services"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	config.InitLog()
}

func main() {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		zap.NewAtomicLevel(),
	)

	logger := zap.New(core, zap.AddCaller())
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	db := ConnectDB(logger)

	repo := NewRepository(db, logger)
	svc := service.NewService(repo, logger)
	controller := controller.ControllerInit(svc, logger)

	app := config.NewInitialization(repo, svc, controller)

	router := router.Init(app)

	// router.Static("/", "./ui")

	router.Run(":8000")

}

func ConnectDB(logger *zap.Logger) (db *sql.DB) {
	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "root:M@roca2002@(127.0.0.1:3306)/teamtask?parseTime=true")
	if err != nil {
		logger.Error("Failed to connect to mysql", zap.Error(err))
	}

	pingErr := db.Ping()
	if pingErr != nil {
		logger.Error("Failed to ping to mysql", zap.Error(err))
	}
	logger.Info("DB Connected!")

	return db
}
