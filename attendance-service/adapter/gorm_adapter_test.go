package adapter

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/night-sornram/employee-management/repository"
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

func TestGetAll(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid-GetAll", func(t *testing.T) {
		mock.ExpectQuery(`select`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		_, err := repo.GetAll()
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-GetAll", func(t *testing.T) {
		mock.ExpectQuery(`select`).
			WillReturnError(errors.New("invalid"))
		_, err := repo.GetAll()
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
func TestGetByID(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid-GetByID", func(t *testing.T) {
		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		_, err := repo.GetByID(1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-GetByID", func(t *testing.T) {
		mock.ExpectQuery(`SELECT`).
			WillReturnError(errors.New("invalid"))
		_, err := repo.GetByID(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCreate(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid-Create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.Create(repository.Attendance{})
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-Create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT`).
			WillReturnError(errors.New("invalid"))
		mock.ExpectRollback()

		_, err := repo.Create(repository.Attendance{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Valid-Update", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		checkIn, _ := time.Parse(time.RFC3339, "2024-05-14T08:00:00Z")
		checkOut, _ := time.Parse(time.RFC3339, "2024-05-16T08:00:00Z")
		attendance := repository.Attendance{
			EmployeeID: "E12779",
			CheckIn:    checkIn,
			CheckOut:   checkOut,
			Date:       "today",
			LeaveID:    1,
		}

		mock.ExpectBegin()
		//?IDK why I can't just use only `UPDATE` like others
		mock.ExpectExec(`UPDATE "attendances" SET "employee_id"=\$1,"check_in"=\$2,"check_out"=\$3,"date"=\$4,"leave_id"=\$5 WHERE id = \$6`).
			WithArgs(attendance.EmployeeID, attendance.CheckIn, attendance.CheckOut, attendance.Date, attendance.LeaveID, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.Update(1, attendance)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Valid-Update", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		checkIn, _ := time.Parse(time.RFC3339, "2024-05-14T08:00:00Z")
		checkOut, _ := time.Parse(time.RFC3339, "2024-05-16T08:00:00Z")
		attendance := repository.Attendance{
			EmployeeID: "E12779",
			CheckIn:    checkIn,
			CheckOut:   checkOut,
			Date:       "today",
			LeaveID:    1,
		}

		mock.ExpectBegin()
		//?IDK why I can't just use only `UPDATE` like others
		mock.ExpectExec(`UPDATE "attendances" SET "employee_id"=\$1,"check_in"=\$2,"check_out"=\$3,"date"=\$4,"leave_id"=\$5 WHERE id = \$6`).
			WithArgs(attendance.EmployeeID, attendance.CheckIn, attendance.CheckOut, attendance.Date, attendance.LeaveID, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.Update(1, attendance)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-Update", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		checkIn, _ := time.Parse(time.RFC3339, "2024-05-14T08:00:00Z")
		checkOut, _ := time.Parse(time.RFC3339, "2024-05-16T08:00:00Z")
		attendance := repository.Attendance{
			EmployeeID: "E12779",
			CheckIn:    checkIn,
			CheckOut:   checkOut,
			Date:       "today",
			LeaveID:    1,
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "attendances" SET "employee_id"=\$1,"check_in"=\$2,"check_out"=\$3,"date"=\$4,"leave_id"=\$5 WHERE id = \$6`).
			WithArgs(attendance.EmployeeID, attendance.CheckIn, attendance.CheckOut, attendance.Date, attendance.LeaveID, 1).
			WillReturnError(errors.New("invalid"))
		mock.ExpectRollback()

		_, err := repo.Update(1, attendance)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestDelete(t *testing.T) {
	t.Run("Valid-Delete", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE`).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Delete(1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-Delete", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE`).
			WillReturnError(errors.New("invalid"))
		mock.ExpectRollback()

		err := repo.Delete(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCheckIn(t *testing.T) {
	t.Run("Valid-CheckIn", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.CheckIn("E12777")
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-CheckIn", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT`).
			WillReturnError(errors.New("invalid"))
		mock.ExpectRollback()

		_, err := repo.CheckIn("E12777")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCheckOut(t *testing.T) {
	t.Run("Valid-CheckOut", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)
		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE`).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.CheckOut(1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-ID-CheckOut", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)
		mock.ExpectQuery(`SELECT`).
			WillReturnError(errors.New("invalid"))

		_, err := repo.CheckOut(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
	t.Run("Invalid-Update-CheckOut", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)
		mock.ExpectQuery(`SELECT`).
			WillReturnError(errors.New("invalid"))
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE`).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.CheckOut(1)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestGetAllMe(t *testing.T) {
	t.Run("Valid-GetAllMe", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)
		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		_, err := repo.GetAllMe("E12777")
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-GetAllMe", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)
		mock.ExpectQuery(`SELECT`).
			WillReturnError(errors.New("invalid"))

		_, err := repo.GetAllMe("E12777")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}

func TestCheckToDay(t *testing.T) {
	t.Run("Valid-CheckToDay", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)
		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		_, err := repo.CheckToday("E12777")
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-CheckToDay", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)
		mock.ExpectQuery(`SELECT`).
			WillReturnError(errors.New("invalid"))

		_, err := repo.CheckToday("E12777")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
	})
}
