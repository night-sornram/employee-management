package adapter

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/night-sornram/employee-management/repository"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
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

		repo := NewGormAdapter(db)

		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		_, err := repo.GetAll()
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-GetAll", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

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

		repo := NewGormAdapter(db)

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

		repo := NewGormAdapter(db)

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

		repo := NewGormAdapter(db)

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "employees"`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.Create(repository.Employee{})
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-Create", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "employees"`).
			WillReturnError(errors.New("invalid"))
		mock.ExpectRollback()

		_, err := repo.Create(repository.Employee{})
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Valid-Update", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		employee := repository.Employee{
			EmployeeID:  "ADMIN",
			TitleTH:     "นาย",
			FirstNameTH: "สมชาย",
			LastNameTH:  "ใจดี",
			TitleEN:     "Mr.",
			FirstNameEN: "Somchai",
			LastNameEN:  "Jaidee",
			DateOfBirth: "1990-01-01",
			Gender:      "Male",
			Department:  "IT",
			Role:        "admin",
			Phone:       "080-123-4567",
			Email:       "admin@example.com",
			Password:    "123456",
		}
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "employees" SET "employee_id"=\$1,"title_th"=\$2,"first_name_th"=\$3,"last_name_th"=\$4,"title_en"=\$5,"first_name_en"=\$6,"last_name_en"=\$7,"date_of_birth"=\$8,"gender"=\$9,"department"=\$10,"role"=\$11,"phone"=\$12,"email"=\$13,"password"=\$14 WHERE employee_id = \$15`).
			WithArgs(employee.EmployeeID, employee.TitleTH, employee.FirstNameTH, employee.LastNameTH, employee.TitleEN, employee.FirstNameEN, employee.LastNameEN, employee.DateOfBirth, employee.Gender, employee.Department, employee.Role, employee.Phone, employee.Email, employee.Password, "1").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.Update("1", employee)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-Update", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		employee := repository.Employee{
			EmployeeID:  "ADMIN",
			TitleTH:     "นาย",
			FirstNameTH: "สมชาย",
			LastNameTH:  "ใจดี",
			TitleEN:     "Mr.",
			FirstNameEN: "Somchai",
			LastNameEN:  "Jaidee",
			DateOfBirth: "1990-01-01",
			Gender:      "Male",
			Department:  "IT",
			Role:        "admin",
			Phone:       "080-123-4567",
			Email:       "admin@example.com",
			Password:    "123456",
		}
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "employees" SET "employee_id"=\$1,"title_th"=\$2,"first_name_th"=\$3,"last_name_th"=\$4,"title_en"=\$5,"first_name_en"=\$6,"last_name_en"=\$7,"date_of_birth"=\$8,"gender"=\$9,"department"=\$10,"role"=\$11,"phone"=\$12,"email"=\$13,"password"=\$14 WHERE employee_id = \$15`).
			WithArgs(employee.EmployeeID, employee.TitleTH, employee.FirstNameTH, employee.LastNameTH, employee.TitleEN, employee.FirstNameEN, employee.LastNameEN, employee.DateOfBirth, employee.Gender, employee.Department, employee.Role, employee.Phone, employee.Email, employee.Password, "1").
			WillReturnError(errors.New("invalid"))
		mock.ExpectRollback()

		_, err := repo.Update("1", employee)
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	t.Run("Valid-Delete", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

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

		repo := NewGormAdapter(db)

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

func TestLogin(t *testing.T) {
	t.Run("Valid-Login", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 14)

		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id", "password"}).AddRow(1, hashedPassword))

		_, err := repo.Login("id", "password")
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-Login", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		mock.ExpectQuery(`SELECT`).
			WillReturnError(errors.New("invalid"))

		_, err := repo.Login("id", "password")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetMe(t *testing.T) {
	t.Run("Valid-GetMe", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		_, err := repo.GetMe("id")
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-GetMe", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		mock.ExpectQuery(`SELECT`).
			WillReturnError(errors.New("invalid"))

		_, err := repo.GetMe("id")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestChangePassword(t *testing.T) {
	t.Run("Valid-ChangePassword", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 14)

		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id", "password"}).AddRow(1, hashedPassword))

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "employees" SET "employee_id"=\$1,"title_th"=\$2,"first_name_th"=\$3,"last_name_th"=\$4,"title_en"=\$5,"first_name_en"=\$6,"last_name_en"=\$7,"date_of_birth"=\$8,"gender"=\$9,"department"=\$10,"role"=\$11,"phone"=\$12,"email"=\$13,"password"=\$14 WHERE "id" = \$15`).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.ChangePassword("id", "password", "newPassword")
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-ID-ChangePassword", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		mock.ExpectQuery(`SELECT`).
			WillReturnError(errors.New("invalid"))

		_, err := repo.ChangePassword("id", "password", "newPassword")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Invalid-ChangePassword", func(t *testing.T) {
		sqlDB, db, mock := DbMock(t)
		defer sqlDB.Close()

		repo := NewGormAdapter(db)

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 14)

		mock.ExpectQuery(`SELECT`).
			WillReturnRows(sqlmock.NewRows([]string{"id", "password"}).AddRow(1, hashedPassword))

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "employees" SET "employee_id"=\$1,"title_th"=\$2,"first_name_th"=\$3,"last_name_th"=\$4,"title_en"=\$5,"first_name_en"=\$6,"last_name_en"=\$7,"date_of_birth"=\$8,"gender"=\$9,"department"=\$10,"role"=\$11,"phone"=\$12,"email"=\$13,"password"=\$14 WHERE "id" = \$15`).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
			WillReturnError(errors.New("invalid"))
		mock.ExpectRollback()

		_, err := repo.ChangePassword("id", "password", "newPassword")
		assert.Error(t, err)
		assert.Equal(t, "invalid", err.Error())
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
