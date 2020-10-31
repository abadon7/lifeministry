package main

import (
	"database/sql"
	"errors"
	//"time"

	_ "github.com/mattn/go-sqlite3"
)

type StorageSQLite struct {
	db *sql.DB
}

func NewStorageSQLite(location string) (*StorageSQLite, error) {
	var err error

	stg := new(StorageSQLite)

	if stg.db != nil {
		return stg, nil
	}

	stg.db, err = sql.Open("sqlite3", location)
	if err != nil {
		panic(err)
	}
	return stg, nil

}

func (s *StorageSQLite) MakeMigrations() error {
	q := `CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64) NOT NULL,
		gender VARCHAR(16) NOT NULL,
		cel INTEGER NULL,
		active BOOL NOT NULL,
		note VARCHAR(128) NULL,
		last TIMESTAMP DEFAULT DATETIME
		);`

	_, err := s.db.Exec(q)
	if err != nil {
		return err

	}
	return nil

}

func (s *StorageSQLite) SaveStudent(student Student) error {
	q := `INSERT INTO students (name,gender,cel,active,note,last)
				VALUES(?,?,?,?,?,?)`

	// y evitar código malicioso.
	stmt, err := s.db.Prepare(q)
	if err != nil {
		return err

	}
	defer stmt.Close()

	r, err := stmt.Exec(student.Name, student.Gender, student.Cel, student.Active, student.Note, student.Last)
	if err != nil {
		return err

	}
	// Confirmamos que una fila fuera afectada, debido a que insertamos un
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")

	}
	// Si llegamos a este punto consideramos que todo el proceso fue exitoso
	return nil
}

func (s *StorageSQLite) FindStudents() ([]Student, error) {
	var students []Student
	var tempUser Student

	q := "SELECT * from students"

	records, err := s.db.Query(q)
	if err != nil {
		return students, err
	}

	defer records.Close()

	for records.Next() {
		records.Scan(&tempUser.ID, &tempUser.Name, &tempUser.Gender, &tempUser.Cel, &tempUser.Active, &tempUser.Note, &tempUser.Last)
		students = append(students, tempUser)
	}

	return students, nil
}
