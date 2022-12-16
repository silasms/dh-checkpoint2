package db

import (
	"ck-2/internal/repository"
	"ck-2/internal/domain"
	"database/sql"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type appointmentDatabase struct {
	db *sql.DB
}

func (a *appointmentDatabase) Post(appointment domain.CreateAppointment) error {
	_, err := a.db.Exec(
		`INSERT INTO appointments
		(patient_id, dentist_id, consult_date)
		VALUES
		(?, ?, ?)`,
		appointment.PatientId,
		appointment.DentistId,
		appointment.AppointmentDate,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *appointmentDatabase) Get(id int) (domain.Appointment, error) {
	var appointment domain.Appointment

	rows, err := a.db.Query(
		`SELECT a.id, a.consult_date, p.*, d.* FROM appointments as a
		INNER JOIN dentists as d ON a.dentist_id = d.id
		INNER JOIN patients as p ON a.patient_id = p.id
		WHERE a.id=?`,
		id)
	if err != nil {
		log.WithError(err).Info("error getting appointment")
		return domain.Appointment{}, err
	}

	for rows.Next() {
		err := rows.Scan(
			&appointment.Id,
			&appointment.AppointmentDate,
			&appointment.Patient.Id,
			&appointment.Patient.Name,
			&appointment.Patient.Surname,
			&appointment.Patient.RG,
			&appointment.Patient.RegistryDate,
			&appointment.Dentist.Id,
			&appointment.Dentist.Name,
			&appointment.Dentist.Surname,
			&appointment.Dentist.Registry,
		)
		if err != nil {
			return domain.Appointment{}, err
		}
	}
	return appointment, nil
}

func (a *appointmentDatabase) GetRG(rg string) (domain.Appointment, error) {
	var appointment domain.Appointment

	rows, err := a.db.Query(
		`SELECT a.id, a.consult_date, p.*, d.* FROM appointments as a
		INNER JOIN dentists as d ON a.dentist_id = d.id
		INNER JOIN patients as p ON a.patient_id = p.id
		WHERE p.rg=?`,
		rg)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		log.WithError(err).Info("error getting appointment")
		return domain.Appointment{}, err
	}

	for rows.Next() {
		err := rows.Scan(
			&appointment.Id,
			&appointment.AppointmentDate,
			&appointment.Patient.Id,
			&appointment.Patient.Name,
			&appointment.Patient.Surname,
			&appointment.Patient.RG,
			&appointment.Patient.RegistryDate,
			&appointment.Dentist.Id,
			&appointment.Dentist.Name,
			&appointment.Dentist.Surname,
			&appointment.Dentist.Registry,
		)
		if err != nil {
			return domain.Appointment{}, err
		}
	}
	return appointment, nil
}

func (a *appointmentDatabase) GetAll() ([]domain.Appointment, error) {
	appointments := []domain.Appointment{}
	rows, err := a.db.Query(
		`SELECT a.id, a.consult_date, p.*, d.* FROM appointments as a
		INNER JOIN dentists as d ON a.dentist_id = d.id
		INNER JOIN patients as p ON a.patient_id = p.id`)

	if err != nil {
		log.WithError(err).Info("error getting appointment")
		return appointments, err
	}

	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(
			&appointment.Id,
			&appointment.AppointmentDate,
			&appointment.Patient.Id,
			&appointment.Patient.Name,
			&appointment.Patient.Surname,
			&appointment.Patient.RG,
			&appointment.Patient.RegistryDate,
			&appointment.Dentist.Id,
			&appointment.Dentist.Name,
			&appointment.Dentist.Surname,
			&appointment.Dentist.Registry,
		)
		appointments = append(appointments, appointment)
		if err != nil {
			return appointments, err
		}
	}
	return appointments, nil
}
func (a *appointmentDatabase) Put(id int, appointment domain.UpdateAppointment) error {
	_, err := a.db.Exec("UPDATE appointments SET patient_id=?, dentist_id=?, consult_date=? WHERE id=?",
		appointment.PatientId, appointment.DentistId, appointment.AppointmentDate, id)

	if err != nil {
		return errors.New("error to update")
	}

	return nil
}
func (a *appointmentDatabase) Patch(id int, appointment domain.PatchAppointment) error {
	_, err := a.db.Exec("UPDATE appointments SET consult_date=? WHERE id=?",
		appointment.AppointmentDate, id)

	if err != nil {
		return errors.New("error to update")
	}

	return nil
}
func (a *appointmentDatabase) Delete(id int) error {
	stmt, err := a.db.Exec("DELETE FROM appointments WHERE id=?", id)
	if err != nil {
		return err
	}

	rowsAffected, _ := stmt.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("error to delete")
	}
	return nil
}

func NewAppointmentDatabase(database *sql.DB) repository.AppointmentRepository {
	return &appointmentDatabase{database}
}
