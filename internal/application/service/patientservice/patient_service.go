package patientservice

import (
	"ck-2/internal/repository"
	"ck-2/internal/application/service"
	"ck-2/internal/domain"

	log "github.com/sirupsen/logrus"
)

type patientService struct {
	repository repository.PatientRepository
}

func (d *patientService) Post(patient domain.CreatePatient) error {
	err := d.repository.Post(patient)
	if err != nil {
		log.WithError(err).Info("error creating patient")
		return err
	}
	return nil
}

func (d *patientService) Get(id int) (domain.Patient, error) {
	patient, err := d.repository.Get(id)
	if err != nil {
		log.WithError(err).Info("error getting patient")
		return domain.Patient{}, err
	}
	return patient, nil
}

func (d *patientService) GetAll() ([]domain.Patient, error) {
	patients, err := d.repository.GetAll()
	if err != nil {
		log.WithError(err).Info("error getting all patient")
		return nil, err
	}
	return patients, nil
}

func (d *patientService) Put(id int, patient domain.UpdatePatient) error {
	err := d.repository.Put(id, patient)
	if err != nil {
		log.WithError(err).Info("error putting the patient")
		return err
	}
	return nil
}

func (d *patientService) Patch(id int, patient domain.PatchPatientName) error {
	err := d.repository.Patch(id, patient)
	if err != nil {
		log.WithError(err).Info("error patching the patient")
		return err
	}
	return nil
}

func (d *patientService) Delete(id int) error {
	err := d.repository.Delete(id)
	if err != nil {
		log.WithError(err).Info("error deleting the patient")
		return err
	}
	return nil
}

func NewPatientService(r repository.PatientRepository) service.Patient {
	return &patientService{r}
}
