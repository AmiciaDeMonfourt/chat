package profile

import (
	"log"
	"log/slog"
	"net"
	"os"
	"pawpawchat/generated/proto/profilepb"
	"pawpawchat/internal/model/domain"
	"pawpawchat/pkg/profile/database"
	"pawpawchat/pkg/profile/server"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(wd + "/.env"); err != nil {
		log.Fatal(err)
	}

	if os.Getenv("PROFILE_ADDR") == "" {
		log.Fatal("profile -> environmet PROFILE_ADDR is missing")

	} else if os.Getenv("PROFILE_DB_URL") == "" {
		log.Fatal("profile -> environment PROFILE_DB_URL is missing")
	}
}

func Start() {
	db := createDB()
	gRPCServer := createServer(db)

	slog.Info("profile -> server is running", "addr", os.Getenv("PROFILE_ADDR"))

	if err := gRPCServer.Serve(createListener()); err != nil {
		log.Fatal("profile -> server error:", err.Error())
	}
}

func createDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("PROFILE_DB_URL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("profile -> failed to connect to db:", err.Error())
	}

	if err := db.AutoMigrate(domain.UserPersonalInfo{}); err != nil {
		log.Fatal("profile -> failed to run migrations:", err.Error())
	}

	return db
}

func createServer(db *gorm.DB) *grpc.Server {
	profileServer := server.New(database.NewPostgresDB(db))

	gRPCServer := grpc.NewServer()

	profilepb.RegisterProfileServer(gRPCServer, profileServer)

	return gRPCServer
}

func createListener() net.Listener {
	listener, err := net.Listen("tcp", os.Getenv("PROFILE_ADDR"))
	if err != nil {
		log.Fatal("profile -> failed to create listener:", err.Error())
	}

	return listener
}
