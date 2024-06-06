package adapter

import (
	"encoding/csv"
	"fmt"
	"strings"

	"math"
	"time"

	"github.com/night-sornram/employee-management/attendance-service/repository"

	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter(db *gorm.DB) repository.AttendanceRepository {
	return &GormAdapter{
		db: db,
	}
}

func (g *GormAdapter) GetAll(query repository.Query) (repository.DataJson, error) {
	var attendances []repository.Attendance

	sql := `SELECT * FROM attendances a JOIN dblink('dbname=employee', 'select employee_id, first_name_en, last_name_en from employees') 
	AS employees(employee_id text, employee_name text, employee_lastname text) on a.employee_id = employees.employee_id`

	if query.Option == "All" || query.Option == "" {
		sql = fmt.Sprintf("%s WHERE 1=1", sql)
	} else if query.Option == "Month" {
		sql = fmt.Sprintf("%s WHERE substring(CAST(a.date AS TEXT),1,7) = '%s'  ", sql, time.Now().Format("2006-01-02")[0:7])
	} else {
		sql = fmt.Sprintf("%s WHERE substring(CAST(a.date AS TEXT),1,4) = '%s'  ", sql, time.Now().Format("2006-01-02")[0:4])
	}

	if query.LeaveID == -1 {
		sql = fmt.Sprintf("%s AND (a.leave_id = -1) ", sql)
	}

	fmt.Println(sql)
	if query.Date == "" {
		if query.Name == "" {
			err := g.db.Raw(sql).Scan(&attendances).Error
			if err != nil {
				return repository.DataJson{}, err
			}
			total := len(attendances)

			sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, query.PerPage, (query.Page-1)*query.PerPage)

			err = g.db.Raw(sql).Scan(&attendances).Error

			if err != nil {
				return repository.DataJson{}, err
			}

			dataJson := repository.DataJson{
				Data:     attendances,
				Total:    total,
				Page:     query.Page,
				LastPage: int(math.Ceil(float64(total) / float64(query.PerPage))),
			}
			return dataJson, nil
		} else {

			sql = fmt.Sprintf("%s AND (employees.employee_name LIKE '%%%s%%' OR employees.employee_lastname LIKE '%%%s%%')", sql, query.Name, query.Name)
			err := g.db.Raw(sql).Scan(&attendances).Error
			if err != nil {
				return repository.DataJson{}, err
			}
			total := len(attendances)

			sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, query.PerPage, (query.Page-1)*query.PerPage)

			err = g.db.Raw(sql).Scan(&attendances).Error

			if err != nil {
				return repository.DataJson{}, err
			}

			dataJson := repository.DataJson{
				Data:     attendances,
				Total:    total,
				Page:     query.Page,
				LastPage: int(math.Ceil(float64(total) / float64(query.PerPage))),
			}
			return dataJson, nil
		}
	} else {
		if query.Name == "" {
			sql = fmt.Sprintf("%s AND a.date = '%s'", sql, query.Date)
			err := g.db.Raw(sql).Scan(&attendances).Error
			if err != nil {
				return repository.DataJson{}, err
			}
			total := len(attendances)

			sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, query.PerPage, (query.Page-1)*query.PerPage)

			err = g.db.Raw(sql).Scan(&attendances).Error

			if err != nil {
				return repository.DataJson{}, err
			}

			dataJson := repository.DataJson{
				Data:     attendances,
				Total:    total,
				Page:     query.Page,
				LastPage: int(math.Ceil(float64(total) / float64(query.PerPage))),
			}
			return dataJson, nil
		} else {
			sql = fmt.Sprintf("%s AND a.date = '%s' AND (employees.employee_name LIKE '%%%s%%' OR employees.employee_lastname LIKE '%%%s%%')", sql, query.Date, query.Name, query.Name)
			err := g.db.Raw(sql).Scan(&attendances).Error
			if err != nil {
				return repository.DataJson{}, err
			}
			total := len(attendances)

			sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, query.PerPage, (query.Page-1)*query.PerPage)

			err = g.db.Raw(sql).Scan(&attendances).Error

			if err != nil {
				return repository.DataJson{}, err
			}

			dataJson := repository.DataJson{
				Data:     attendances,
				Total:    total,
				Page:     query.Page,
				LastPage: int(math.Ceil(float64(total) / float64(query.PerPage))),
			}
			return dataJson, nil
		}
	}
}

func (g *GormAdapter) GetByID(id int) (repository.Attendance, error) {
	var attendance repository.Attendance
	err := g.db.First(&attendance, id).Error
	if err != nil {
		return attendance, err
	}
	return attendance, nil
}

func (g *GormAdapter) Create(attendance repository.Attendance) (repository.Attendance, error) {
	err := g.db.Create(&attendance).Error
	if err != nil {
		return attendance, err
	}
	return attendance, nil
}

func (g *GormAdapter) Update(id int, attendance repository.Attendance) (repository.Attendance, error) {
	err := g.db.Model(&attendance).Where("id = ?", id).Updates(attendance).Error
	if err != nil {
		return attendance, err
	}
	return attendance, nil
}

func (g *GormAdapter) Delete(id int) error {
	err := g.db.Delete(&repository.Attendance{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (g *GormAdapter) CheckIn(eid string) (repository.Attendance, error) {

	newAttendance := repository.Attendance{
		EmployeeID: eid,
		CheckIn:    time.Now(),
		CheckOut:   time.Time{},
		Date:       time.Now().Format("2006-01-02"),
		LeaveID:    -1,
	}
	err := g.db.Create(&newAttendance).Error
	if err != nil {
		return newAttendance, err
	}

	return newAttendance, nil
}

func (g *GormAdapter) CheckOut(id int) (repository.Attendance, error) {

	var attendance repository.Attendance
	err := g.db.First(&attendance, id).Error
	if err != nil {
		return attendance, err
	}
	attendance.CheckOut = time.Now()
	err = g.db.Save(&attendance).Error
	if err != nil {
		return attendance, err
	}
	return attendance, nil

}

func (g *GormAdapter) GetAllMe(query repository.Query, eid string) (repository.DataJson, error) {
	var attendances []repository.Attendance

	sql := `SELECT * FROM attendances a `

	if query.Option == "All" || query.Option == "" {
		sql = fmt.Sprintf("%s WHERE 1=1", sql)
	} else if query.Option == "Month" {
		sql = fmt.Sprintf("%s WHERE substring(CAST(a.date AS TEXT),1,7) = '%s'  ", sql, time.Now().Format("2006-01-02")[0:7])
	} else {
		sql = fmt.Sprintf("%s WHERE substring(CAST(a.date AS TEXT),1,4) = '%s'  ", sql, time.Now().Format("2006-01-02")[0:4])
	}

	if query.Date == "" {
		sql = fmt.Sprintf("%s AND a.employee_id = '%s'", sql, eid)
		err := g.db.Raw(sql).Scan(&attendances).Error
		if err != nil {
			return repository.DataJson{}, err
		}
		total := len(attendances)
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, query.PerPage, (query.Page-1)*query.PerPage)
		err = g.db.Raw(sql).Scan(&attendances).Error
		if err != nil {
			return repository.DataJson{}, err
		}

		dataJson := repository.DataJson{
			Data:     attendances,
			Total:    total,
			Page:     query.Page,
			LastPage: int(math.Ceil(float64(total) / float64(query.PerPage))),
		}
		return dataJson, nil
	} else {
		sql = fmt.Sprintf("%s AND a.employee_id = '%s' AND date = '%s' ", sql, eid, query.Date)
		err := g.db.Raw(sql).Scan(&attendances).Error
		if err != nil {
			return repository.DataJson{}, err
		}
		total := len(attendances)
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, query.PerPage, (query.Page-1)*query.PerPage)
		err = g.db.Raw(sql).Scan(&attendances).Error
		if err != nil {
			return repository.DataJson{}, err
		}

		dataJson := repository.DataJson{
			Data:     attendances,
			Total:    total,
			Page:     query.Page,
			LastPage: int(math.Ceil(float64(total) / float64(query.PerPage))),
		}
		return dataJson, nil
	}

}

func (g *GormAdapter) CheckToday(eid string) (repository.Attendance, error) {
	var attendance repository.Attendance
	err := g.db.Where("employee_id = ? AND date = ?", eid, time.Now().Format(time.DateOnly)).First(&attendance).Error
	if err != nil {
		return repository.Attendance{}, nil
	}
	return attendance, nil
}

func (g *GormAdapter) GetDayLate() ([]repository.Attendance, error) {
	var attendances []repository.Attendance
	query := `SELECT * FROM attendances a JOIN dblink('dbname=employee', 'select employee_id, first_name_en, last_name_en from employees') 
				AS employees(employee_id text, employee_name text, employee_lastname text) on a.employee_id = employees.employee_id 
         		WHERE check_in > '0001-01-01 10:00:00' AND EXTRACT(DAY FROM check_in)= EXTRACT(DAY FROM CURRENT_DATE);`
	err := g.db.Raw(query).Scan(&attendances).Error
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

func (g *GormAdapter) GetMonthLate(month int, year int) ([]repository.Attendance, error) {
	var attendances []repository.Attendance
	query := fmt.Sprintf(`SELECT * FROM attendances a JOIN dblink('dbname=employee', 'select employee_id, first_name_en, last_name_en from employees') 
				AS employees(employee_id text, employee_name text, employee_lastname text) on a.employee_id = employees.employee_id 
				WHERE check_in > '0001-01-01 10:00:00' AND EXTRACT(MONTH FROM check_in)= %d AND EXTRACT(YEAR FROM check_in)= %d;`, month, year)
	err := g.db.Raw(query).Scan(&attendances).Error
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

func (g *GormAdapter) GetYearLate(year int) ([]repository.Attendance, error) {
	var attendances []repository.Attendance
	query := fmt.Sprintf(`SELECT * FROM attendances a JOIN dblink('dbname=employee', 'select employee_id, first_name_en, last_name_en from employees') 
				AS employees(employee_id text, employee_name text, employee_lastname text) on a.employee_id = employees.employee_id 
         		WHERE check_in > '0001-01-01 10:00:00' AND EXTRACT(YEAR FROM check_in)= %d;`, year)
	err := g.db.Raw(query).Scan(&attendances).Error
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

func (g *GormAdapter) GetAllLate() ([]repository.Attendance, error) {
	var attendances []repository.Attendance
	query := `SELECT * FROM attendances a JOIN dblink('dbname=employee', 'select employee_id, first_name_en, last_name_en from employees') 
	AS employees(employee_id text, employee_name text, employee_lastname text) on a.employee_id = employees.employee_id WHERE check_in > '0001-01-01 10:00:00';`
	err := g.db.Raw(query).Scan(&attendances).Error
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

func (g *GormAdapter) GetCSV(query string) ([]byte, error) {
	var results []repository.Attendance
	if err := g.db.Raw(`SELECT * FROM attendances a JOIN dblink('dbname=employee', 'select employee_id, first_name_en, last_name_en from employees') 
		AS employees(employee_id text, employee_name text, employee_lastname text) on a.employee_id = employees.employee_id 
		WHERE EXTRACT(HOUR FROM check_in) > 1 OR (EXTRACT(HOUR FROM check_in) = 1 AND EXTRACT(MINUTE FROM check_in) > 30);`).Scan(&results).Error; err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no data found")
	}
	var b strings.Builder
	w := csv.NewWriter(&b)

	header := []string{"ID", "EmployeeID", "CheckIn", "CheckOut", "Date", "LeaveID", "EmployeeName", "EmployeeLastname"}
	if err := w.Write(header); err != nil {
		return nil, err
	}

	// Write rows
	for _, att := range results {
		record := []string{
			fmt.Sprintf("%d", att.ID),
			att.EmployeeID,
			att.CheckIn.Format(time.DateTime),
			att.CheckOut.Format(time.DateTime),
			att.Date,
			fmt.Sprintf("%d", att.LeaveID),
			att.EmployeeName,
			att.EmployeeLastname,
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
