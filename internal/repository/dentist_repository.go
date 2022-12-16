package repository

import "ck-2/internal/domain"

type DentistRepository interface {
	Post(domain.CreateDentist) error
	Get(int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Put(int, domain.UpdateDentist) error
	Patch(int, domain.PatchDentistName) error
	Delete(int) error
}