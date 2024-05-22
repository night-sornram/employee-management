package adapter

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/night-sornram/employee-management/repository"
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
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid-GetAll", func(t *testing.T) {
		mock.ExpectQuery(`SELECT (.+) FROM "attendances"`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		_, err := repo.GetAll()
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestGetByID(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid-GetByID", func(t *testing.T) {
		mock.ExpectQuery(`SELECT (.+) FROM "attendances" WHERE "attendances"."id" = \$1 ORDER BY "attendances"."id" LIMIT \$2`).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		_, err := repo.GetByID(1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestCreate(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid-Create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "attendances" (.+) RETURNING "id"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		checkIn, _ := time.Parse(time.RFC3339, "2024-05-14T08:00:00Z")
		checkOut, _ := time.Parse(time.RFC3339, "2024-05-16T08:00:00Z")

		_, err := repo.Create(repository.Attendance{
			EmployeeID: "E12779",
			CheckIn:    checkIn,
			CheckOut:   checkOut,
			Date:       "today",
			LeaveID:    1,
		})
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdate(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid-Update", func(t *testing.T) {
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
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.Update(1, attendance)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid-Delete", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "attendances" WHERE "attendances"."id" = \$1`).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Delete(1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
