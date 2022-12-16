package dentistservice

import (
	"ck-2/internal/repository"
	"ck-2/internal/application/service"
	"ck-2/internal/domain"

	log "github.com/sirupsen/logrus"
)

type dentistService struct {
	repository repository.DentistRepository
}

func (d *dentistService) Post(dentist domain.CreateDentist) error {
	err := d.repository.Post(dentist)
	if err != nil {
		log.WithError(err).Info("error creating dentist")
		return err
	}
	return nil
}

func (d *dentistService) Get(id int) (domain.Dentist, error) {
	dentist, err := d.repository.Get(id)
	if err != nil {
		log.WithError(err).Info("error getting dentist")
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (d *dentistService) GetAll() ([]domain.Dentist, error) {
	dentists, err := d.repository.GetAll()
	if err != nil {
		log.WithError(err).Info("error getting all dentists")
		return nil, err
	}
	return dentists, nil
}

func (d *dentistService) Put(id int, dentist domain.UpdateDentist) error {
	err := d.repository.Put(id, dentist)
	if err != nil {
		log.WithError(err).Info("error putting the dentist")
		return err
	}
	return nil
}

func (d *dentistService) Patch(id int, dentist domain.PatchDentistName) error {
	err := d.repository.Patch(id, dentist)
	if err != nil {
		log.WithError(err).Info("error patching the dentist")
		return err
	}
	return nil
}

func (d *dentistService) Delete(id int) error {
	err := d.repository.Delete(id)
	if err != nil {
		log.WithError(err).Info("error deleting the dentist")
		return err
	}
	return nil
}

func NewDentistService(r repository.DentistRepository) service.Dentist {
	return &dentistService{r}
}
