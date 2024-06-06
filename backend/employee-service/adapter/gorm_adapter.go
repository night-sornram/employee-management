package adapter

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"fmt"
	"strings"

	"github.com/night-sornram/employee-management/employee-service/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter(db *gorm.DB) repository.EmployeeRepository {
	return &GormAdapter{
		db: db,
	}
}

func (g *GormAdapter) GetAll() ([]repository.Employee, error) {
	var Employees []repository.Employee
	if err := g.db.Find(&Employees).Error; err != nil {
		return nil, err
	}
	return Employees, nil
}

func (g *GormAdapter) GetByID(eid string) (repository.Employee, error) {
	var Employee repository.Employee
	if err := g.db.Model(&Employee).Where("employee_id = ?", eid).First(&Employee).Error; err != nil {
		return Employee, err

	}
	return Employee, nil
}

func (g *GormAdapter) Create(Employee repository.Employee) (repository.Employee, error) {
	if err := g.db.Create(&Employee).Error; err != nil {
		return Employee, err
	}
	return Employee, nil
}

func (g *GormAdapter) Update(id string, Employee repository.Employee) (repository.Employee, error) {
	if err := g.db.Model(&Employee).Where("employee_id = ?", id).Updates(Employee).Error; err != nil {
		return Employee, err
	}
	return Employee, nil
}

func (g *GormAdapter) Delete(id int) error {
	if err := g.db.Delete(&repository.Employee{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormAdapter) Login(id string, password string) (repository.Employee, error) {
	var Employee repository.Employee
	if err := g.db.First(&Employee, "employee_id = ?", id).Error; err != nil {
		return repository.Employee{}, err
	}
	// -- bcrypt --
	// if err := bcrypt.CompareHashAndPassword([]byte(Employee.Password), []byte(password)); err != nil {
	// 	return repository.Employee{}, err
	// }

	// -- sha256 --
	hash := sha256.Sum256([]byte(password))
	hashPassword := base64.StdEncoding.EncodeToString(hash[:])
	if Employee.Password != hashPassword {
		return repository.Employee{}, errors.New("incorrect password")
	}

	return Employee, nil
}

func (g *GormAdapter) GetMe(id string) (repository.Employee, error) {
	var Employee repository.Employee
	if err := g.db.Where("employee_id = ?", id).First(&Employee).Error; err != nil {
		return repository.Employee{}, err
	}

	return Employee, nil
}

func (g *GormAdapter) ChangePassword(id string, password string, new_password string) (repository.Employee, error) {

	var Employee repository.Employee
	if err := g.db.Where("employee_id = ?", id).First(&Employee).Error; err != nil {
		return repository.Employee{}, err
	}
	hash := sha256.Sum256([]byte(password))
	hashPassword := base64.StdEncoding.EncodeToString(hash[:])
	if Employee.Password != hashPassword {
		return repository.Employee{}, errors.New("incorrect password")
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(new_password), 14)
	Employee.Password = string(hashedPassword)

	if err := g.db.Save(&Employee).Error; err != nil {
		return repository.Employee{}, err
	}
	return Employee, nil
}

func (g *GormAdapter) GetCSV(query string) ([]byte, error) {
	var results []repository.Employee
	if err := g.db.Raw(`SELECT * FROM employees;`).Scan(&results).Error; err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no data found")
	}
	var b strings.Builder
	w := csv.NewWriter(&b)

	header := []string{"ID", "EmployeeID", "TitleTH", "FirstNameTH", "LastNameTH", "TitleEN", "FirstNameEN", "LastNameEN", "DateOfBirth", "Gender", "Department", "Role", "Phone", "Email", "Password"}
	if err := w.Write(header); err != nil {
		return nil, err
	}

	// Write rows
	for _, emp := range results {
		record := []string{
			fmt.Sprintf("%d", emp.ID),
			emp.EmployeeID,
			emp.TitleTH,
			emp.FirstNameTH,
			emp.LastNameTH,
			emp.TitleEN,
			emp.FirstNameEN,
			emp.LastNameEN,
			emp.DateOfBirth,
			emp.Gender,
			emp.Department,
			emp.Role,
			emp.Phone,
			emp.Email,
			emp.Password,
		}
		if err := w.Write(record); err != nil {
			return nil, err
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		return nil, err
	}

	return []byte(b.String()), nil
}
