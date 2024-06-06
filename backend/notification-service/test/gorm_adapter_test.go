package test

import (
	"database/sql"
	"errors"
	"github.com/night-sornram/employee-management/notification-service/adapter"
	"github.com/night-sornram/employee-management/notification-service/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm database: %v", err)
	}

	return sqlDB, db, mock
}

func TestGetAll(t *testing.T) {
	t.Run("Valid-GetAll", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := adapter.NewGormAdapter(db)

		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		_, err := repo.GetAll()
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-GetAll", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := adapter.NewGormAdapter(db)

		mock.ExpectQuery(`SELECT`).
			WillReturnError(errors.New("invalid"))
		_, err := repo.GetAll()
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Valid-GetByID", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := adapter.NewGormAdapter(db)

		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		_, err := repo.GetByID(1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-GetByID", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := adapter.NewGormAdapter(db)

		mock.ExpectQuery(`SELECT`).
			WillReturnError(errors.New("invalid"))
		_, err := repo.GetByID(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestCreate(t *testing.T) {
	t.Run("Valid-Create", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := adapter.NewGormAdapter(db)

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.Create(repository.Notification{})
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-Create", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := adapter.NewGormAdapter(db)

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT`).
			WillReturnError(errors.New("invalid"))
		mock.ExpectRollback()

		_, err := repo.Create(repository.Notification{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Valid-Update", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := adapter.NewGormAdapter(db)

		notification := repository.Notification{
			EmployeeID: "E123",
			Message:    "This is a test message",
			Title:      "Test Title",
			Read:       true,
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE`).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.Update(1, notification)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-Update", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := adapter.NewGormAdapter(db)

		notification := repository.Notification{
			EmployeeID: "E123",
			Message:    "This is a test message",
			Title:      "Test Title",
			Read:       true,
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE`).
			WillReturnError(errors.New("invalid"))
		mock.ExpectRollback()

		_, err := repo.Update(1, notification)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	t.Run("Valid-Delete", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := adapter.NewGormAdapter(db)

		mock.ExpectBegin()
		mock.ExpectExec(`DELETE`).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Delete(1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-Delete", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := adapter.NewGormAdapter(db)

		mock.ExpectBegin()
		mock.ExpectExec(`DELETE`).
			WillReturnError(errors.New("invalid"))
		mock.ExpectRollback()

		err := repo.Delete(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
