package initialization

import (
	"log"
	"net"
	"net/http"
	"time"
	"time-logger/internal/pkg/config"
	"time-logger/internal/pkg/database-access"
	handler "time-logger/shared/http-wrappers"
)

func StartupEnv() *handler.Env {
	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}

	cr := config.NewReader(client)

	config, err := cr.GetConfig()

	if err != nil {
		log.Fatal(err)
	}

	session, err := database_access.NewSession(config.Server)
	if err != nil {
		log.Fatal(err)
	}

	db := session.DB(config.Database)

	env := &handler.Env{
		DB:         db,
		HttpClient: client,
		// We might also have a custom log.Logger, our
		// template instance, and a config struct as fields
		// in our Env struct.
	}

	return env
}
