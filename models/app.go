package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"database/sql"
	"beego-mvc-api/utils"
)

type App struct {
	Id          int       `orm:"column(id);auto" sql:"id"`
	CreateDate  time.Time `orm:"column(create_date);type(datetime)" sql:"create_date"`
	AppCode     string    `orm:"column(app_code);size(255)" sql:"app_code"`
	AppName     string    `orm:"column(app_name);size(255)" sql:"app_name"`
	PublishDate time.Time `orm:"column(publish_date);type(date);null" sql:"publish_date"`
}
type Apps struct{
	Apps []App `json:"apps"`
}
func (t *App) TableName() string {
	return "app"
}

func init() {
	orm.RegisterModel(new(App))
}

// AddApp insert a new App into database and returns
// last inserted Id on success.
func AddApp(m *App) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAppById retrieves App by Id. Returns error if
// Id doesn't exist
func GetAppById(id int) (v *App, err error) {
	o := orm.NewOrm()
	v = &App{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllApp retrieves all App matches certain condition. Returns empty list if
// no records exist
func GetAllApp(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(App))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []App
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateApp updates App by Id and returns error if
// the record to be updated doesn't exist
func UpdateAppById(m *App) (err error) {
	o := orm.NewOrm()
	v := App{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteApp deletes App by Id and returns error if
// the record to be deleted doesn't exist
func DeleteApp(id int) (err error) {
	o := orm.NewOrm()
	v := App{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&App{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

var con *sql.DB;
func GetApps() Apps{
	con := utils.CreateCon()
	sqlStatement := "select id,create_date,app_code,app_name,publish_date from app"
	result := Apps{}
	rows, err := con.Query(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}
	//defer rows.Close()

	/**
	* one way get db rows to struct
	 */
	for rows.Next() {
		record := App{}
		utils.ToStruct(rows, &record)
		fmt.Println("-------------",record.Id, record.CreateDate, record.AppCode,record.AppName,record.PublishDate)
		result.Apps = append(result.Apps, record)
	}

	/**
	* the other way to get db row data to struct
	 */
	//for rows.Next() {
	//	app := App{}
	//	err2 := rows.Scan(&app.Id,&app.CreateDate,&app.AppCode,&app.AppName,&app.PublishDate)
	//	if err2 != nil {
	//		fmt.Print(err2)
	//	}
	//	result.Apps = append(result.Apps, app)
	//}
	return result
}
