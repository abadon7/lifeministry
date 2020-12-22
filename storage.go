package main

type StorageType int

const (
	SQLiteLocation             = "./school.db"
	SQLite         StorageType = iota
)

type Storage interface {
	SaveStudent(...Student) error
	FindStudent(Student) ([]Student, error)
	UpdateStudent(Student) error
	RemoveStudent(Student) error
	FindStudents(string, string) ([]Student, error)

	SaveAssigment(...Assigment) error
	FindAssigments() ([]Assigment, error)
	FindAssigment(Assigment) ([]Assigment, error)

	MakeMigrations() error
}

func NewStorage(storageType StorageType) (Storage, error) {
	var stg Storage
	var err error

	switch storageType {
	case SQLite:
		stg, err = NewStorageSQLite(SQLiteLocation)
	}
	return stg, err
}
