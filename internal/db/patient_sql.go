package db

import (
	"ck-2/internal/repository"
	"ck-2/internal/domain"
	"database/sql"
	"errors"
)

type patientDatabase struct {
	db *sql.DB
}

func (d *patientDatabase) Post(patient domain.CreatePatient) error {
	_, err := d.db.Exec(
		`INSERT INTO patients
		(name, surname, rg, registry_date)
		VALUES
		(?, ?, ?, ?)`,
		patient.Name,
		patient.Surname,
		patient.RG,
		patient.RegistryDate,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *patientDatabase) Get(id int) (domain.Patient, error) {
	var patient domain.Patient

	rows, err := d.db.Query(`SELECT * from patients WHERE id=?`, id)
	if err != nil {
		return domain.Patient{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&patient.Id,
			&patient.Name,
			&patient.Surname,
			&patient.RG,
			&patient.RegistryDate,
		)
		if err != nil {
			return domain.Patient{}, err
		}
	}
	return patient, nil
}

func (d *patientDatabase) GetAll() ([]domain.Patient, error) {
	var patients []domain.Patient

	rows, err := d.db.Query("SELECT id, name, surname, rg, registry_date FROM patients")

	if err != nil {
		return patients, err
	}

	defer rows.Close()

	for rows.Next() {
		var patient domain.Patient

		err := rows.Scan(
			&patient.Id,
			&patient.Name,
			&patient.Surname,
			&patient.RG,
			&patient.RegistryDate,
		)
		if err != nil {
			return patients, err
		}
		patients = append(patients, patient)
	}
	return patients, nil
}

func (d *patientDatabase) Put(id int, patient domain.UpdatePatient) error {

	_, err := d.db.Exec("UPDATE patients SET name=?, surname=?, rg=?, registry_date=? WHERE id=?",
		patient.Name, patient.Surname, patient.RG, patient.RegistryDate, id)

	if err != nil {
		return errors.New("error to update")
	}

	return nil
}
func (d *patientDatabase) Patch(id int, patient domain.PatchPatientName) error {
	_, err := d.db.Exec("UPDATE patients SET name=? WHERE id=?",
		patient.Name, id)

	if err != nil {
		return errors.New("error to update")
	}

	return nil
}
func (d *patientDatabase) Delete(id int) error {
	stmt, err := d.db.Exec("DELETE FROM patients WHERE id=?", id)
	if err != nil {
		return err
	}

	rowsAffected, _ := stmt.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("error to delete")
	}
	return nil
}

func NewPatientDatabase(database *sql.DB) repository.PatientRepository {
	return &patientDatabase{database}
}
