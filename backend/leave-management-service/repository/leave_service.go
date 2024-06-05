package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type LeaveService interface {
	GetLeaves(query Query) (DataJson, error)
	GetLeave(id int) (Leave, error)
	CreateLeave(leave Leave) (Leave, error)
	UpdateLeave(id int, leave Leave) (Leave, error)
	DeleteLeave(id int) error
	UpdateStatus(id int, leave LeaveStatus) (Leave, error)
	GetAllMe(query Query, eid string) (DataJson, error)
}

type LeaveServiceDB struct {
	repo LeaveRepository
}

func NewLeaveService(repo LeaveRepository) LeaveService {
	return &LeaveServiceDB{
		repo: repo,
	}
}

func (u *LeaveServiceDB) GetLeaves(query Query) (DataJson, error) {
	return u.repo.GetAll(query)
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

func (u *LeaveServiceDB) GetAllMe(query Query, eid string) (DataJson, error) {
	return u.repo.GetAllMe(query, eid)
}

type Attendance struct {
	ID         int       `db:"id" json:"id"`
	EmployeeID string    `db:"employee_id" json:"employee_id"`
	CheckIn    time.Time `db:"check_in" json:"check_in"`
	CheckOut   time.Time `db:"check_out" json:"check_out"`
	Date       time.Time `db:"date" json:"date"`
	LeaveID    int       `db:"leave_id" json:"leave_id"`
}

func PostAttendance(payload Attendance) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:8081/api/attendance", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (u *LeaveServiceDB) UpdateStatus(id int, leave LeaveStatus) (Leave, error) {
	existsLeave, err := u.repo.GetByID(id)
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
			err := PostAttendance(payload)
			if err != nil {
				return Leave{}, err
			}
		}
	}

	existsLeave.Status = leave.Status
	existsLeave.ManagerOpinion = leave.ManagerOpinion

	return u.repo.UpdateStatus(id, existsLeave)
}
