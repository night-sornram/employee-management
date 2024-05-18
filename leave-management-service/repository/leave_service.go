package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type LeaveService interface {
	GetLeaves() ([]Leave, error)
	GetLeave(id int) (Leave, error)
	CreateLeave(leave Leave) (Leave, error)
	UpdateLeave(id int, leave Leave) (Leave, error)
	DeleteLeave(id int) error
	UpdateStatus(id int, leave Leave) (Leave, error)
}

type LeaveServiceDB struct {
	repo LeaveRepository
}

func NewLeaveService(repo LeaveRepository) LeaveService {
	return &LeaveServiceDB{
		repo: repo,
	}
}

func (u *LeaveServiceDB) GetLeaves() ([]Leave, error) {
	return u.repo.GetAll()
}

func (u *LeaveServiceDB) GetLeave(id int) (Leave, error) {
	return u.repo.GetByID(id)
}

func (u *LeaveServiceDB) CreateLeave(leave Leave) (Leave, error) {
	return u.repo.Create(leave)
}

func (u *LeaveServiceDB) UpdateLeave(id int, leave Leave) (Leave, error) {
	return u.repo.Update(id, leave)
}

func (u *LeaveServiceDB) DeleteLeave(id int) error {
	return u.repo.Delete(id)
}

type Attendance struct {
	ID         int       `db:"id" json:"id"`
	EmployeeID string    `db:"employee_id" json:"employee_id"`
	CheckIn    time.Time `db:"check_in" json:"check_in"`
	CheckOut   time.Time `db:"check_out" json:"check_out"`
	Date       time.Time `db:"date" json:"date"`
	LeaveID    int       `db:"leave_id" json:"leave_id"`
}

func (u *LeaveServiceDB) UpdateStatus(id int, leave Leave) (Leave, error) {	
	existsLeave, err := u.repo.GetByID(id);
	if err != nil {
		return Leave{}, err
	}

	if leave.Status == "approve" {
		for d := existsLeave.DateStart; !d.After(existsLeave.DateEnd); d = d.AddDate(0, 0, 1) {
			payload := Attendance{
				EmployeeID: existsLeave.EmployeeID,
				CheckIn:    d,
				CheckOut:   d,
				Date:       d,
				LeaveID:    existsLeave.ID,
			}
			jsonData, err := json.Marshal(payload)
			if err != nil {
				return Leave{}, err
			}

			req, err := http.NewRequest("POST", "http://localhost:8081/attendance", bytes.NewBuffer(jsonData))
			if err != nil {
				return Leave{}, err
			}

			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return Leave{}, err
			}
			defer resp.Body.Close()

		}

	}
	return u.repo.UpdateStatus(id, leave)
}
