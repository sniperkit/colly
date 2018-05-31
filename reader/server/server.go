package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	log "github.com/sirupsen/logrus"
	"github.com/timjchin/logcounter"
)

type ServerConfig struct {
	Port string
}

type Server struct {
	config     *ServerConfig
	logcounter *logcounter.LogCounter
}

func NewServer(config *ServerConfig, l *logcounter.LogCounter) *Server {
	if config.Port == "" {
		config.Port = "6002"
	}
	return &Server{
		config:     config,
		logcounter: l,
	}
}

func (s *Server) Start() {
	http.HandleFunc("/ws", s.wsHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", s.config.Port), nil))
}

func (s *Server) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w, nil)
	if err != nil {
		log.WithField("err", err).Error("Failed to upgrade websocket connection")
		return
	}
	go func() {
		defer conn.Close()

		for {
			currState := s.logcounter.GetState()
			m, err := json.Marshal(currState)
			if err != nil {
				log.WithField("err", err).Fatal("Failed to marshal logcounter state.")
				return
			}
			err = wsutil.WriteServerMessage(conn, ws.OpText, m)
			if err != nil {
				log.WithField("err", err).Error("Failed to write to client.")
				return
			}
			time.Sleep(time.Second)
		}
	}()
}
