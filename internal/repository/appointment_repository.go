package repository

import "ck-2/internal/domain"

type AppointmentRepository interface {
	Post(domain.CreateAppointment) error
	Get(int) (domain.Appointment, error)
	GetRG(string) (domain.Appointment, error)
	GetAll() ([]domain.Appointment, error)
	Put(int, domain.UpdateAppointment) error
	Patch(int, domain.PatchAppointment) error
	Delete(int) error
}