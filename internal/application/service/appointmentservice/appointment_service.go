package appointmentservice

import (
	"ck-2/internal/repository"
	"ck-2/internal/application/service"
	"ck-2/internal/domain"

	log "github.com/sirupsen/logrus"
)

type appointmentService struct {
	repository        repository.AppointmentRepository
	dentistRepository repository.DentistRepository
	patientRepository repository.PatientRepository
}

func (d *appointmentService) Post(appointment domain.CreateAppointment) error {
	_, err := d.dentistRepository.Get(appointment.DentistId)
	if err != nil {
		log.WithError(err).Info("error getting dentist")
		return err
	}

	_, err = d.patientRepository.Get(appointment.PatientId)
	if err != nil {
		log.WithError(err).Info("error getting dentist")
		return err
	}

	err = d.repository.Post(appointment)
	if err != nil {
		log.WithError(err).Info("error creating appointment")
		return err
	}
	return nil
}

func (d *appointmentService) Get(id int) (domain.Appointment, error) {
	appointment, err := d.repository.Get(id)
	if err != nil {
		log.WithError(err).Info("error getting appointment")
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (d *appointmentService) GetRG(rg string) (domain.Appointment, error) {
	appointment, err := d.repository.GetRG(rg)
	if err != nil {
		log.WithError(err).Info("error getting appointment")
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (d *appointmentService) GetAll() ([]domain.Appointment, error) {
	appointments, err := d.repository.GetAll()
	if err != nil {
		log.WithError(err).Info("error getting all appointments")
		return nil, err
	}
	return appointments, nil
}

func (d *appointmentService) Put(id int, appointment domain.UpdateAppointment) error {
	err := d.repository.Put(id, appointment)
	if err != nil {
		log.WithError(err).Info("error putting the appointment")
		return err
	}
	return nil
}

func (d *appointmentService) Patch(id int, appointment domain.PatchAppointment) error {
	err := d.repository.Patch(id, appointment)
	if err != nil {
		log.WithError(err).Info("error patching the appointment")
		return err
	}
	return nil
}

func (d *appointmentService) Delete(id int) error {
	err := d.repository.Delete(id)
	if err != nil {
		log.WithError(err).Info("error deleting the appointment")
		return err
	}
	return nil
}

func NewAppointmentService(
	r repository.AppointmentRepository,
	dr repository.DentistRepository,
	pr repository.PatientRepository,
) service.Appointment {
	return &appointmentService{r, dr, pr}
}
