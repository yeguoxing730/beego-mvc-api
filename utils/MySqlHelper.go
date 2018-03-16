package utils
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Unknwon/goconfig"
	"fmt"
	"reflect"
	"errors"
)

var db *sql.DB
func init()  {
	cfg,err:=goconfig.LoadConfigFile("./conf/mysql.conf")
	if err!=nil{
		panic("加载mysql配置失败")
	}
	host,err := cfg.GetValue("mysql","host");
	if err != nil{
		panic(err)
	}
	username,err := cfg.GetValue("mysql","username");
	if err != nil{
		panic(err)
	}
	password,err := cfg.GetValue("mysql","password");
	if err != nil{
		panic(err)
	}
	maxOpen:=cfg.MustInt("mysql","maxOpen")
	maxIdle:=cfg.MustInt("mysql","maxIdle")
	db, err = sql.Open("mysql", username+":"+password+"@"+host)
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)
	db.Ping()
}
/**
执行sql返回执行是否成功
 */
func ExecSql(sql string,args ...interface{}) (err error){
	if args==nil{
		_,err=db.Exec(sql)
	}else{
		_,err=db.Exec(sql,args...)
	}
	return
}
/*Create mysql connection*/
func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "root:Welcome1@tcp(localhost:3306)/mysql?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}

/*end mysql connection*/



func ToStruct(rows *sql.Rows, to interface{}) error {
	v := reflect.ValueOf(to)
	if v.Elem().Type().Kind() != reflect.Struct {
		return errors.New("Expect a struct")
	}

	scan_dest := []interface{}{}
	column_names, _ := rows.Columns()

	addr_by_column_name := map[string]interface{}{}

	for i := 0; i < v.Elem().NumField(); i++ {
		one_value := v.Elem().Field(i)
		column_name := v.Elem().Type().Field(i).Tag.Get("sql")
		if column_name == "" {
			column_name = one_value.Type().Name()
		}
		addr_by_column_name[column_name] = one_value.Addr().Interface()
	}

	for _, column_name := range column_names {
		scan_dest = append(scan_dest, addr_by_column_name[column_name])
	}

	return rows.Scan(scan_dest...)

}

