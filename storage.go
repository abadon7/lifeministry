package main

type StorageType int

const (
	SQLiteLocation             = "./school.db"
	SQLite         StorageType = iota
)

type Storage interface {
	SaveStudent(Student) error
	//FindStudent(Student) ([]Student, error)
	FindStudents() ([]Student, error)
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
