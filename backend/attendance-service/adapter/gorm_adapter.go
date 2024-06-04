package adapter

import (
	"fmt"
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

	if query.Option == "All" {
		sql = fmt.Sprintf("%s WHERE 1=1", sql)
	} else if query.Option == "Month" {
		sql = fmt.Sprintf("%s WHERE substring(CAST(a.date AS TEXT),1,7) = '%s'  ", sql, time.Now().Format("2006-01-02")[0:7])
	} else {
		sql = fmt.Sprintf("%s WHERE substring(CAST(a.date AS TEXT),1,4) = '%s'  ", sql, time.Now().Format("2006-01-02")[0:4])
	}

	fmt.Println(sql)

	if query.Date == "" {
		if query.Name == "" {
			err := g.db.Find(&attendances).Error
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
		return attendance, err
	}
	return attendance, nil
}

/*
package adapter

import (
	"fmt"
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

	if query.Date == "" {
		if query.Name == "" {
			err := g.db.Find(&attendances).Error
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

			sql = fmt.Sprintf("%s WHERE employees.employee_name LIKE '%%%s%%' OR employees.employee_lastname LIKE '%%%s%%'", sql, query.Name, query.Name)
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
			sql = fmt.Sprintf("%s WHERE a.date = '%s'", sql, query.Date)
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
			sql = fmt.Sprintf("%s WHERE a.date = '%s' AND (employees.employee_name LIKE '%%%s%%' OR employees.employee_lastname LIKE '%%%s%%')", sql, query.Date, query.Name, query.Name)
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

	if query.Date == "" {
		err := g.db.Where("employee_id = ?", eid).Find(&attendances).Error
		if err != nil {
			return repository.DataJson{}, err
		}
		total := len(attendances)
		err = g.db.Where("employee_id = ? ", eid).Limit(query.PerPage).Offset((query.Page - 1) * query.PerPage).Find(&attendances).Error
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
		err := g.db.Where("employee_id = ? AND date = ? ", eid, query.Date).Find(&attendances).Error
		if err != nil {
			return repository.DataJson{}, err
		}
		total := len(attendances)
		err = g.db.Where("employee_id = ? AND date = ? ", eid, query.Date).Limit(query.PerPage).Offset((query.Page - 1) * query.PerPage).Find(&attendances).Error
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
		return attendance, err
	}
	return attendance, nil
}

*/
