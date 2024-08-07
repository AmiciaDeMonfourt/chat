package profile

import (
	"flag"
	"log"
	"log/slog"
	"net"
	"pawpawchat/generated/proto/profilepb"
	"pawpawchat/internal/model/domain"
	"pawpawchat/pkg/profile/config"
	"pawpawchat/pkg/profile/database"
	"pawpawchat/pkg/profile/server"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Start() {
	env := flag.String("env", "dev", "Environment to use [dev/test]")
	flag.Parse()

	cfg := config.New(*env)
	db := sqldDbConn(cfg)

	gRPCServer := createServer(db)

	slog.Info("server is running", "addr", cfg.Addr)

	if err := gRPCServer.Serve(createListener(cfg)); err != nil {
		log.Fatal("profile -> server error:", err.Error())
	}
}

func sqldDbConn(cfg *config.ProfileConfig) *gorm.DB {
	var level logger.LogLevel
	if cfg.LogLevel == "info" {
		level = logger.Info
	} else {
		level = logger.Error
	}

	db, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})
	if err != nil {
		log.Fatal("failed to connect to db:", err.Error())
	}

	if err := db.AutoMigrate(domain.UserBiography{}); err != nil {
		log.Fatal("failed to run migrations:", err.Error())
	}
	return db
}

func createServer(db *gorm.DB) *grpc.Server {
	profileServer := server.New(database.NewPostgresDB(db))
	gRPCServer := grpc.NewServer()

	profilepb.RegisterProfileServer(gRPCServer, profileServer)
	return gRPCServer
}

func createListener(cfg *config.ProfileConfig) net.Listener {
	listener, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		log.Fatal("failed to create listener:", err.Error())
	}
	return listener
}
