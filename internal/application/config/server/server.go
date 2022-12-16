package server

import (
	"ck-2/internal/application"
	"ck-2/internal/application/config/routes"
	"ck-2/internal/application/service/appointmentservice"
	"ck-2/internal/application/service/dentistservice"
	"ck-2/internal/application/service/patientservice"
	"ck-2/internal/db"
	"database/sql"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (addr = ":8080")

type server struct {
	app *gin.Engine
}

func (s *server) Run() error {
	err := s.app.Run(addr)
	if err != nil {
		log.WithFields(log.Fields{
			"port": addr,
		}).WithError(err).Info("error trying to running the server")
	}
	return nil
}

func NewServer() application.Application {
	return &server{
		app: gin.Default(),
	}
}

func (s *server) Setup() error {

	database, err := s.configureDatabase()
	if err != nil {
		log.WithError(err).Error("error when configure the database")
		return err
	}

	dentistDatabase := db.NewDentistDatabase(database)
	dentistService := dentistservice.NewDentistService(dentistDatabase)

	patientDatabase := db.NewPatientDatabase(database)
	patientService := patientservice.NewPatientService(patientDatabase)

	appointmentDatabase := db.NewAppointmentDatabase(database)
	appointmentService := appointmentservice.NewAppointmentService(
		appointmentDatabase,
		dentistDatabase,
		patientDatabase,
	)

	router := routes.NewRoutes(
		s.app,
		dentistService,
		patientService,
		appointmentService,
	)

	router.SetupRoutes()

	return nil
}


func (s *server) configureDatabase() (*sql.DB, error) {
	database, err := db.Connection()
	if err != nil {
		log.WithError(err).Info("database error")
		return nil, err
	}
	return database, nil
}