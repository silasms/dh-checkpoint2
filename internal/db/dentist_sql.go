package db

import (
	"ck-2/internal/repository"
	"ck-2/internal/domain"
	"database/sql"
	"errors"
)

type dentistDatabase struct {
	db *sql.DB
}

func (d *dentistDatabase) Post(dentist domain.CreateDentist) error {
	_, err := d.db.Exec(
		`INSERT INTO dentists (name, surname, registry) VALUES
		(?, ?, ?)`,
		dentist.Name,
		dentist.Surname,
		dentist.Registry,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *dentistDatabase) Get(id int) (domain.Dentist, error) {
	var dentist domain.Dentist

	rows, err := d.db.Query(`SELECT * from dentists WHERE id=?`, id)
	if err != nil {
		return domain.Dentist{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&dentist.Id,
			&dentist.Name,
			&dentist.Surname,
			&dentist.Registry,
		)
		if err != nil {
			return domain.Dentist{}, err
		}
	}
	return dentist, nil
}

func (d *dentistDatabase) GetAll() ([]domain.Dentist, error) {
	var dentists []domain.Dentist

	rows, err := d.db.Query("SELECT id, name, surname, registry FROM dentists")

	if err != nil {
		return dentists, err
	}

	defer rows.Close()

	for rows.Next() {
		var dentist domain.Dentist

		err := rows.Scan(
			&dentist.Id,
			&dentist.Name,
			&dentist.Surname,
			&dentist.Registry,
		)
		if err != nil {
			return dentists, err
		}
		dentists = append(dentists, dentist)
	}
	return dentists, nil
}
func (d *dentistDatabase) Put(id int, dentist domain.UpdateDentist) error {

	_, err := d.db.Exec("UPDATE dentists SET name=?, surname=?, registry=? WHERE id=?",
		dentist.Name, dentist.Surname, dentist.Registry, id)

	if err != nil {
		return errors.New("error to update")
	}

	return nil
}
func (d *dentistDatabase) Patch(id int, dentist domain.PatchDentistName) error {

	_, err := d.db.Exec("UPDATE dentists SET name=? WHERE id=?",
		dentist.Name, id)

	if err != nil {
		return errors.New("error to update")
	}

	return nil
}
func (d *dentistDatabase) Delete(id int) error {
	stmt, err := d.db.Exec("DELETE FROM dentists WHERE id=?", id)
	if err != nil {
		return err
	}

	rowsAffected, _ := stmt.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("error to delete")
	}
	return nil
}

func NewDentistDatabase(database *sql.DB) repository.DentistRepository {
	return &dentistDatabase{database}
}
