package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

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
		last TIMESTAMP DEFAULT DATETIME,
		lastpartner INTEGER NOT NULL);`
	_, err := s.db.Exec(q)
	if err != nil {
		return err
	}

	qA := `CREATE TABLE IF NOT EXISTS assigments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64) NOT NULL,
		participants INTEGER NOT NULL,
		type VARCHAR(16));`
	if _, err := s.db.Exec(qA); err != nil {
		return err
	}

	qB := `CREATE TABLE IF NOT EXISTS assignedtolist (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		inchargeid INTEGER NOT NULL,
		helperid INTEGER NOT NULL,
		date TIMESTAMP DEFAULT DATETIME,
		assigmenttype INTEGER NOT NULL
	);`
	if _, err := s.db.Exec(qB); err != nil {
		return err
	}

	qC := `CREATE TABLE IF NOT EXISTS schedules(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		data TEXT,
		range TEXT
	);`
	if _, err := s.db.Exec(qC); err != nil {
		return err
	}

	qD := `CREATE TABLE IF NOT EXISTS weeks(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TIMESTAMP DEFAULT DATETIME,
		data TEXT
	);`
	if _, err := s.db.Exec(qD); err != nil {
		return err
	}

	return nil
}

func (s *StorageSQLite) SaveStudent(students ...Student) error {
	q := `INSERT INTO students (name,gender,cel,active,note,last,lastpartner)
				VALUES(?,?,?,?,?,?,?)`

	for _, student := range students {
		stmt, err := s.db.Prepare(q)
		if err != nil {
			return err

		}
		defer stmt.Close()

		r, err := stmt.Exec(student.Name, student.Gender, student.Cel, student.Active, student.Note, student.Last, student.LastPartner)
		if err != nil {
			return err

		}
		// Confirmamos que una fila fuera afectada, debido a que insertamos un
		if i, err := r.RowsAffected(); err != nil || i != 1 {
			return errors.New("ERROR: Se esperaba una fila afectada")

		}
	}
	// Si llegamos a este punto consideramos que todo el proceso fue exitoso
	return nil
}

func (s *StorageSQLite) FindStudents(active string, gender string) ([]Student, error) {
	var students []Student
	var tempUser Student

	if active == "all" {

		q := "SELECT * from students"

		if gender != "all" {
			q = "SELECT * FROM students WHERE gender = ?"
		}

		records, err := s.db.Query(q, gender)
		if err != nil {
			return students, err
		}

		defer records.Close()

		for records.Next() {
			records.Scan(&tempUser.ID, &tempUser.Name, &tempUser.Gender, &tempUser.Cel, &tempUser.Active, &tempUser.Note, &tempUser.Last, &tempUser.LastPartner)
			students = append(students, tempUser)
		}
		return students, nil
	}

	if active == "active" {
		q2 := "SELECT * from students WHERE active = 1"

		if gender != "all" {
			q2 = "SELECT * FROM students WHERE active = 1 and gender = ?"
		}

		records, err := s.db.Query(q2, gender)
		if err != nil {
			return students, err
		}

		defer records.Close()

		for records.Next() {
			records.Scan(&tempUser.ID, &tempUser.Name, &tempUser.Gender, &tempUser.Cel, &tempUser.Active, &tempUser.Note, &tempUser.Last, &tempUser.LastPartner)
			students = append(students, tempUser)
		}
		return students, nil
	}

	return students, nil
}

func (s *StorageSQLite) FindStudent(criteria Student) ([]Student, error) {
	var students []Student

	q := "SELECT * from students WHERE id=?"

	if criteria.ID != 0 {
		var student Student
		err := s.db.QueryRow(q, criteria.ID).Scan(&student.ID, &student.Name, &student.Gender, &student.Cel, &student.Active, &student.Note, &student.Last, &student.LastPartner)
		if err != nil {
			return students, err
		}

		students = append(students, student)
		return students, nil
	}
	return students, fmt.Errorf("No student id especified")
}

func (s *StorageSQLite) UpdateStudent(student Student) error {
	q := `UPDATE students set name=?, gender=?, cel=?, active=?, note=?, last=?, lastpartner=? WHERE id=?`

	// y evitar código malicioso.
	stmt, err := s.db.Prepare(q)
	if err != nil {
		return err

	}
	defer stmt.Close()

	r, err := stmt.Exec(student.Name, student.Gender, student.Cel, student.Active, student.Note, student.Last, student.LastPartner, student.ID)
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

func (s *StorageSQLite) RemoveStudent(criteria Student) error {
	q := "DELETE FROM students WHERE id=?"

	stmt, err := s.db.Prepare(q)
	if err != nil {
		return err

	}
	defer stmt.Close()

	r, err := stmt.Exec(criteria.ID)
	if err != nil {
		return err

	}
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")

	}
	return nil
}

func (s *StorageSQLite) SaveAssigment(assigments ...Assigment) error {
	q := `INSERT INTO assigments (name,participants,type) VALUES(?,?,?)`

	// y evitar código malicioso.
	for _, assigment := range assigments {
		stmt, err := s.db.Prepare(q)
		if err != nil {
			return err
		}

		defer stmt.Close()

		r, err := stmt.Exec(assigment.Name, assigment.Participants, assigment.Type)
		if err != nil {
			return err
		}
		// Confirmamos que una fila fuera afectada, debido a que insertamos un
		if i, err := r.RowsAffected(); err != nil || i != 1 {
			return errors.New("ERROR: Se esperaba una fila afectada")
		}
	}
	// Si llegamos a este punto consideramos que todo el proceso fue exitoso
	return nil
}

func (s *StorageSQLite) FindAssigments() ([]Assigment, error) {
	var assigments []Assigment
	var tempAssigment Assigment

	q := "SELECT * from assigments"

	records, err := s.db.Query(q)
	if err != nil {
		return assigments, err
	}

	defer records.Close()

	for records.Next() {
		records.Scan(&tempAssigment.ID, &tempAssigment.Name, &tempAssigment.Participants, &tempAssigment.Type)
		assigments = append(assigments, tempAssigment)
	}

	return assigments, nil
}

func (s *StorageSQLite) FindAssigment(criteria Assigment) ([]Assigment, error) {
	var assigments []Assigment

	q := "SELECT * from assigments WHERE id=?"

	if criteria.ID != 0 {
		var assigment Assigment
		err := s.db.QueryRow(q, criteria.ID).Scan(&assigment.ID, &assigment.Name, &assigment.Participants, &assigment.Type)
		if err != nil {
			return assigments, err
		}

		assigments = append(assigments, assigment)
		return assigments, nil
	}
	return assigments, fmt.Errorf("No assigment id especified")
}

func (s *StorageSQLite) SaveSchedule(schedule Schedule) (Schedule, error) {
	q := `INSERT INTO schedules (data,range) VALUES(?,?)`

	// y evitar código malicioso.
	//for _, schedule := range schedules {
	stmt, err := s.db.Prepare(q)
	if err != nil {
		return schedule, err
	}

	defer stmt.Close()

	r, err := stmt.Exec(schedule.Data, schedule.Range)
	if err != nil {
		return schedule, err
	}
	// Confirmamos que una fila fuera afectada, debido a que insertamos un
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return schedule, errors.New("ERROR: Se esperaba una fila afectada")
	}
	lid, err := r.LastInsertId()

	schedule.ID = lid

	//}
	// Si llegamos a este punto consideramos que todo el proceso fue exitoso

	fmt.Println("This is the las id saved " + strconv.FormatInt(lid, 10))
	return schedule, nil
}

func (s *StorageSQLite) FindSchedules() ([]Schedule, error) {
	var schedules []Schedule
	var tempSchedule Schedule

	q := "SELECT * from schedules"

	records, err := s.db.Query(q)
	if err != nil {
		return schedules, err
	}

	defer records.Close()

	for records.Next() {
		records.Scan(&tempSchedule.ID, &tempSchedule.Data, &tempSchedule.Range)
		schedules = append(schedules, tempSchedule)
	}

	return schedules, nil
}

func (s *StorageSQLite) FindSchedule(criteria int) (Schedule, error) {

	fmt.Println("Consulting schedule" + strconv.Itoa(criteria))
	var schedule Schedule
	q := "SELECT * from schedules WHERE id=?"

	if criteria != 0 {
		err := s.db.QueryRow(q, criteria).Scan(&schedule.ID, &schedule.Data, &schedule.Range)
		if err != nil {
			return schedule, err
		}

		return schedule, nil

	}
	return schedule, fmt.Errorf("No assigment id especified")
}
