package routes

import (
	"ck-2/internal/application/service"
	"ck-2/internal/handler/appointmenthandler"
	"ck-2/internal/handler/dentisthandler"
	"ck-2/internal/handler/patienthandler"

	"github.com/gin-gonic/gin"
)

type routes struct {
	server             *gin.Engine
	dentistService     service.Dentist
	patientService     service.Patient
	appointmentService service.Appointment
}

func NewRoutes(server *gin.Engine, dentistService service.Dentist, patientService service.Patient, appointmentService service.Appointment) routes {
	return routes{
		server,
		dentistService,
		patientService,
		appointmentService,
	}
}

func (r *routes) SetupRoutes() {

	dentistGroup := r.server.Group("/dentistas")
	patientGroup := r.server.Group("/pacientes")
	appointmentGroup := r.server.Group("/consultas")

	routesDentist := dentisthandler.NewDentistHandler(dentistGroup, r.dentistService)
	routesDentist.ConfigureDentistRouter()

	routesPatient := patienthandler.NewPatientHandler(patientGroup, r.patientService)
	routesPatient.ConfigurePatientRouter()

	routesAppointment := appointmenthandler.NewAppointmentHandler(appointmentGroup, r.appointmentService)
	routesAppointment.ConfigureAppointmentRouter()

}

