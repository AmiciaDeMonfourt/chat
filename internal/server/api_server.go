package server

import "log"

func Start() {
	srv := newServer()

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("api_server -> start[9]: ", err)
	}
}
