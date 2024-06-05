package test

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/night-sornram/employee-management/leave-management-service/adapter"
	"github.com/night-sornram/employee-management/leave-management-service/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
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

func getLeaveData() repository.Leave {
	dateStart, _ := time.Parse(time.RFC3339, "2024-05-14T08:00:00Z")
	dateEnd, _ := time.Parse(time.RFC3339, "2024-05-16T08:00:00Z")
	leave := repository.Leave{
		EmployeeID: "E12779",
		DateStart:  dateStart,
		DateEnd:    dateEnd,
		Reason:     "reason",
		Status:     "pending",
	}
	return leave
}

func TestGetAll(t *testing.T) {
	t.Run("Valid-GetAll", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		repo := adapter.NewGormAdapter(db)

		mock.ExpectQuery(`select`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		_, err := repo.GetAll()
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-GetAll", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		repo := adapter.NewGormAdapter(db)

		mock.ExpectQuery(`select`).
			WillReturnError(errors.New("invalid"))
		_, err := repo.GetAll()
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Valid-GetByID", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		repo := adapter.NewGormAdapter(db)
		mock.ExpectQuery(`SELECT`).
			WithArgs(1, 1).
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

		_, err := repo.Create(repository.Leave{})
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

		_, err := repo.Create(repository.Leave{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Valid-Update", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		repo := adapter.NewGormAdapter(db)

		leave := getLeaveData()
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE`).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.Update(1, leave)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-Update", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		repo := adapter.NewGormAdapter(db)

		leave := getLeaveData()
		mock.ExpectBegin()
		// IDK why I can't use mock.ExpectExec(`UPDATE`) without .WithArgs like others methods
		mock.ExpectExec(`UPDATE`).
			WillReturnError(errors.New("invalid"))
		mock.ExpectRollback()

		_, err := repo.Update(1, leave)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestDelete(t *testing.T) {
	t.Run("Valid Delete", func(t *testing.T) {
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
	t.Run("Valid Delete", func(t *testing.T) {
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
	})
}

func TestUpdateStatus(t *testing.T) {
	t.Run("Valid-UpdateStatus", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		repo := adapter.NewGormAdapter(db)
		defer sqlDB.Close()
		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE`).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.UpdateStatus(1, repository.Leave{})
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-Select-UpdateStatus", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		repo := adapter.NewGormAdapter(db)
		defer sqlDB.Close()
		mock.ExpectQuery(`SELECT`).
			WillReturnError(errors.New("invalid"))
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE`).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.UpdateStatus(1, repository.Leave{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
	t.Run("Invalid-Update-UpdateStatus", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		repo := adapter.NewGormAdapter(db)
		defer sqlDB.Close()
		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE`).
			WillReturnError(errors.New("invalid"))
		mock.ExpectRollback()

		_, err := repo.UpdateStatus(1, repository.Leave{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetAllMe(t *testing.T) {
	t.Run("Valid-GetAllMe", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		repo := adapter.NewGormAdapter(db)

		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		_, err := repo.GetAllMe("E12777")
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-GetAllMe", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()
		repo := adapter.NewGormAdapter(db)

		mock.ExpectQuery(`SELECT`).
			WillReturnError(errors.New("invalid"))
		_, err := repo.GetAllMe("E12777")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
