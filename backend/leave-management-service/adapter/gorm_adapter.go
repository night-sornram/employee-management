package adapter

import (
	"encoding/csv"
	"fmt"
	"github.com/night-sornram/employee-management/leave-management-service/repository"
	"gorm.io/gorm"
	"math"
	"strings"
	"time"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter(db *gorm.DB) repository.LeaveRepository {
	return &GormAdapter{
		db: db,
	}
}

func (g *GormAdapter) GetAll(query repository.Query) (repository.DataJson, error) {
	var leaves []repository.Leave
	sql := `SELECT * FROM leaves a JOIN dblink('dbname=employee', 'select employee_id, first_name_en, last_name_en from employees') 
	AS employees(employee_id text, employee_name text, employee_lastname text) on a.employee_id = employees.employee_id`

	if query.Option == "All" || query.Option == "" {
		sql = fmt.Sprintf("%s WHERE 1=1", sql)
	} else if query.Option == "Month" {
		sql = fmt.Sprintf("%s WHERE (substring(CAST(a.date_start AS TEXT),1,7) = '%s' OR  substring(CAST(a.date_end AS TEXT),1,7) = '%s' ) ", sql, time.Now().Format("2006-01-02")[0:7], time.Now().Format("2006-01-02")[0:7])
	} else {
		sql = fmt.Sprintf("%s WHERE (substring(CAST(a.date_start AS TEXT),1,4) = '%s' OR  substring(CAST(a.date_start AS TEXT),1,4) = '%s' ) ", sql, time.Now().Format("2006-01-02")[0:4], time.Now().Format("2006-01-02")[0:4])
	}

	if query.Status != "" {
		if query.Status == "Pending" {
			sql = fmt.Sprintf("%s AND (a.category = 'Pending') ", sql)
		} else {
			sql = fmt.Sprintf("%s AND (a.category = 'Approved' OR a.category = 'Denied' ) ", sql)
		}
	}

	if query.Date == "" {
		if query.Name == "" {
			err := g.db.Raw(sql).Scan(&leaves).Error
			if err != nil {
				return repository.DataJson{}, err
			}
			total := len(leaves)

			sql = fmt.Sprintf("%s   LIMIT %d OFFSET %d", sql, query.PerPage, (query.Page-1)*query.PerPage)

			err = g.db.Raw(sql).Scan(&leaves).Error

			if err != nil {
				return repository.DataJson{}, err
			}

			dataJson := repository.DataJson{
				Data:     leaves,
				Total:    total,
				Page:     query.Page,
				LastPage: int(math.Ceil(float64(total) / float64(query.PerPage))),
			}
			return dataJson, nil
		} else {

			sql = fmt.Sprintf("%s AND (employees.employee_name LIKE '%%%s%%' OR employees.employee_lastname LIKE '%%%s%%')", sql, query.Name, query.Name)
			err := g.db.Raw(sql).Scan(&leaves).Error
			if err != nil {
				return repository.DataJson{}, err
			}
			total := len(leaves)

			sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, query.PerPage, (query.Page-1)*query.PerPage)

			err = g.db.Raw(sql).Scan(&leaves).Error

			if err != nil {
				return repository.DataJson{}, err
			}

			dataJson := repository.DataJson{
				Data:     leaves,
				Total:    total,
				Page:     query.Page,
				LastPage: int(math.Ceil(float64(total) / float64(query.PerPage))),
			}
			return dataJson, nil
		}
	} else {
		if query.Name == "" {
			sql = fmt.Sprintf("%s AND (substring(CAST(a.date_start AS TEXT),1,10) = '%s' OR substring(CAST(a.date_end AS TEXT),1,10) = '%s') ", sql, query.Date, query.Date)
			err := g.db.Raw(sql).Scan(&leaves).Error
			if err != nil {
				return repository.DataJson{}, err
			}
			total := len(leaves)

			sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, query.PerPage, (query.Page-1)*query.PerPage)

			err = g.db.Raw(sql).Scan(&leaves).Error

			if err != nil {
				return repository.DataJson{}, err
			}

			dataJson := repository.DataJson{
				Data:     leaves,
				Total:    total,
				Page:     query.Page,
				LastPage: int(math.Ceil(float64(total) / float64(query.PerPage))),
			}
			return dataJson, nil
		} else {
			sql = fmt.Sprintf("%s AND (substring(CAST(a.date_start AS TEXT),1,10) = '%s' OR substring(CAST(a.date_end AS TEXT),1,10) = '%s') AND (employees.employee_name LIKE '%%%s%%' OR employees.employee_lastname LIKE '%%%s%%')", sql, query.Date, query.Date, query.Name, query.Name)
			err := g.db.Raw(sql).Scan(&leaves).Error
			if err != nil {
				return repository.DataJson{}, err
			}
			total := len(leaves)

			sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, query.PerPage, (query.Page-1)*query.PerPage)

			err = g.db.Raw(sql).Scan(&leaves).Error

			if err != nil {
				return repository.DataJson{}, err
			}

			dataJson := repository.DataJson{
				Data:     leaves,
				Total:    total,
				Page:     query.Page,
				LastPage: int(math.Ceil(float64(total) / float64(query.PerPage))),
			}
			return dataJson, nil
		}
	}
}

func (g *GormAdapter) GetByID(id int) (repository.Leave, error) {
	var leave repository.Leave
	if err := g.db.First(&leave, id).Error; err != nil {
		return leave, err
	}
	return leave, nil
}

func (g *GormAdapter) Create(leave repository.Leave) (repository.Leave, error) {
	if err := g.db.Create(&leave).Error; err != nil {
		return leave, err
	}
	return leave, nil
}

func (g *GormAdapter) Update(id int, leave repository.Leave) (repository.Leave, error) {
	if err := g.db.Model(&leave).Where("id = ?", id).Updates(leave).Error; err != nil {
		return leave, err
	}
	return leave, nil
}

func (g *GormAdapter) Delete(id int) error {
	if err := g.db.Delete(&repository.Leave{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormAdapter) UpdateStatus(id int, leave repository.Leave) (repository.Leave, error) {
	var existingLeave repository.Leave
	if err := g.db.Where("id = ?", id).First(&existingLeave).Error; err != nil {
		return leave, err
	}

	existingLeave.Status = leave.Status
	existingLeave.ManagerOpinion = leave.ManagerOpinion

	if err := g.db.Save(&existingLeave).Error; err != nil {
		return existingLeave, err
	}
	return existingLeave, nil
}

func (g *GormAdapter) GetAllMe(query repository.Query, eid string) (repository.DataJson, error) {
	var leaves []repository.Leave
	sql := `SELECT * FROM leaves a JOIN dblink('dbname=employee', 'select employee_id, first_name_en, last_name_en from employees') 
	AS employees(employee_id text, employee_name text, employee_lastname text) on a.employee_id = employees.employee_id`

	if query.Option == "All" || query.Option == "" {
		sql = fmt.Sprintf("%s WHERE 1=1", sql)
	} else if query.Option == "Month" {
		sql = fmt.Sprintf("%s WHERE (substring(CAST(a.date_start AS TEXT),1,7) = '%s' OR  substring(CAST(a.date_end AS TEXT),1,7) = '%s' ) ", sql, time.Now().Format("2006-01-02")[0:7], time.Now().Format("2006-01-02")[0:7])
	} else {
		sql = fmt.Sprintf("%s WHERE (substring(CAST(a.date_start AS TEXT),1,4) = '%s' OR  substring(CAST(a.date_start AS TEXT),1,4) = '%s' ) ", sql, time.Now().Format("2006-01-02")[0:4], time.Now().Format("2006-01-02")[0:4])
	}

	if query.Date == "" {
		err := g.db.Raw(sql).Scan(&leaves).Error
		if err != nil {
			return repository.DataJson{}, err
		}
		total := len(leaves)

		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, query.PerPage, (query.Page-1)*query.PerPage)

		err = g.db.Raw(sql).Scan(&leaves).Error

		if err != nil {
			return repository.DataJson{}, err
		}

		dataJson := repository.DataJson{
			Data:     leaves,
			Total:    total,
			Page:     query.Page,
			LastPage: int(math.Ceil(float64(total) / float64(query.PerPage))),
		}
		return dataJson, nil
	} else {
		sql = fmt.Sprintf("%s AND (substring(CAST(a.date_start AS TEXT),1,10) = '%s' OR substring(CAST(a.date_end AS TEXT),1,10) = '%s' )", sql, query.Date, query.Date)
		err := g.db.Raw(sql).Scan(&leaves).Error
		if err != nil {
			return repository.DataJson{}, err
		}
		total := len(leaves)

		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, query.PerPage, (query.Page-1)*query.PerPage)

		err = g.db.Raw(sql).Scan(&leaves).Error

		if err != nil {
			return repository.DataJson{}, err
		}

		dataJson := repository.DataJson{
			Data:     leaves,
			Total:    total,
			Page:     query.Page,
			LastPage: int(math.Ceil(float64(total) / float64(query.PerPage))),
		}
		return dataJson, nil
	}

}

func (g *GormAdapter) GetCSV(query string) ([]byte, error) {
	var results []repository.Leave
	if err := g.db.Raw(`SELECT * FROM leaves a JOIN dblink('dbname=employee', 'select employee_id, first_name_en, last_name_en from employees') 
	AS employees(employee_id text, employee_name text, employee_lastname text) on a.employee_id = employees.employee_id`).Scan(&results).Error; err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no data found")
	}
	var b strings.Builder
	w := csv.NewWriter(&b)

	header := []string{"ID", "EmployeeID", "DateStart", "DateEnd", "Reason", "Category", "ManagerOpinion", "Status", "Manager", "EmployeeName", "EmployeeLastname"}
	if err := w.Write(header); err != nil {
		return nil, err
	}

	// Write rows
	for _, leave := range results {
		record := []string{
			fmt.Sprintf("%d", leave.ID),
			leave.EmployeeID,
			leave.DateStart.Format(time.DateTime),
			leave.DateEnd.Format(time.DateTime),
			leave.Reason,
			leave.Category,
			leave.ManagerOpinion,
			leave.Status,
			leave.ManagerOpinion,
			leave.EmployeeName,
			leave.EmployeeLastname,
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
