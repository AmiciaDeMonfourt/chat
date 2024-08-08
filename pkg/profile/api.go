package profile

import (
	"log"
	"log/slog"
	"net"
	"pawpawchat/config"
	"pawpawchat/generated/proto/profilepb"
	"pawpawchat/pkg/profile/controller"
	"pawpawchat/pkg/profile/handler"
	"pawpawchat/pkg/profile/repository/factory"

	"google.golang.org/grpc"
)

// Start ...
func Start() {
	config, envConfig := config.GetConfiguration("config.yaml")
	profileHandler := newProfileHandler(envConfig.ProfileEnvCfg())
	runGRPCServer(config.Profile, profileHandler)
}

// newProfileHandler ...
func newProfileHandler(envConfig config.EnvCfg) *handler.Profile {
	repoFactory := factory.NewPostgresRepositoryFactory()
	profileRepository := repoFactory.OpenProfile(envConfig.DBURL, "orm")
	return handler.NewProfile(controller.NewProfile(profileRepository))
}

// runGRPCServer ...
func runGRPCServer(cfg config.ServiceConfig, ph *handler.Profile) {
	gRPSServer := grpc.NewServer()
	profilepb.RegisterProfileServer(gRPSServer, newServer(ph))

	listener, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		log.Fatal(err)
	}

	slog.Debug("grpc server is running", "addr", cfg.Addr)
	if err := gRPSServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
