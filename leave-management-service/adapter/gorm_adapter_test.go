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

func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock){
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

func TestGetAll(t *testing.T){
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	
	repo := NewGormAdapter(db)
	t.Run("Valid GetAll", func(t *testing.T){
		mock.ExpectQuery(`SELECT (.+) FROM "leaves"`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		_ , err := repo.GetAll()
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestGetByID(t *testing.T){
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid GetByID", func(t *testing.T){
		mock.ExpectQuery(`SELECT (.+) FROM "leaves" WHERE "leaves"."id" = \$1 ORDER BY "leaves"."id" LIMIT \$2`).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		_ , err := repo.GetByID(1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestCreate(t *testing.T){
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid Create", func(t *testing.T){
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "leaves" (.+) RETURNING "id"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		dateStart, _ := time.Parse(time.RFC3339, "2024-05-14T08:00:00Z")
		dateEnd, _ := time.Parse(time.RFC3339, "2024-05-16T08:00:00Z")

		_ , err := repo.Create(repository.Leave{
			EmployeeID: "E12779",
			DateStart: dateStart,
			DateEnd: dateEnd,
			Reason: "reason",
			Status: "pending",
		})
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdate(t *testing.T){
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid Update", func(t *testing.T){
		dateStart, _ := time.Parse(time.RFC3339, "2024-05-14T08:00:00Z")
		dateEnd, _ := time.Parse(time.RFC3339, "2024-05-16T08:00:00Z")
		leave := repository.Leave{
			EmployeeID: "E12779",
			DateStart: dateStart,
			DateEnd: dateEnd,
			Reason: "reason",
			Status: "pending",
		}
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "leaves" SET "employee_id"=\$1,"date_start"=\$2,"date_end"=\$3,"reason"=\$4,"status"=\$5 WHERE id = \$6`).
			WithArgs(leave.EmployeeID, leave.DateStart, leave.DateEnd, leave.Reason, leave.Status, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_ , err := repo.Update(1, leave)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T){
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid Delete", func(t *testing.T){
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "leaves" WHERE "leaves"."id" = \$1`).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Delete(1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdateStatus(t *testing.T){
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	repo := NewGormAdapter(db)
	t.Run("Valid UpdateStatus", func(t *testing.T){
		leave := repository.Leave{
			Status: "pending",
			ManagerOpinion: "Approve",
		}

		mock.ExpectQuery(`SELECT \* FROM "leaves" WHERE id = \$1 ORDER BY "leaves"."id" LIMIT \$2`).
            WithArgs(1, 1).
            WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "leaves" SET "employee_id"=\$1,"date_start"=\$2,"date_end"=\$3,"reason"=\$4,"manager_opinion"=\$5,"status"=\$6 WHERE "id" = \$7`).
            WithArgs("", time.Time{}, time.Time{}, "", leave.ManagerOpinion, leave.Status, 1).
            WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.UpdateStatus(1, leave)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}