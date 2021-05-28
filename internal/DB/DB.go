package DB

import (
	"fmt"
	config "lab3/config"
	prod "lab3/internal/products"
	"github.com/jmoiron/sqlx"
	_ "github.com/identitii/gdbc/postgresql"
)

var Connect *sqlx.DB
var conf = config.GetConfig()

func init() {
	Connect = ConnectToDb()
}
var products *[]prod.Product

func ConnectToDb() *sqlx.DB {
	driver := conf.Driver
	host := "host=" + conf.HostPostgre
	user := "user=" + conf.UserNamePostgre + ":"
	password := "password=" + conf.PasswordPostgre
	dbname := "dbname=" + conf.DBname
	dsn := user + password + "@" + driver+ ":postgresql://" + host + "/" + dbname

	con, err := sqlx.Connect(driver, dsn)
	if err != nil {
		fmt.Println("Подключение к базе невозможно.", err)
		return  nil
	}
	return con
}

func InsertProducts(prod *[]prod.Product, flags string){
	tableName := ""
	switch flags{
		case "1":
			tableName = "keyboards"
		case "2":
			tableName = "mouses"
	}

	tx := Connect.MustBegin()
	_, err := tx.NamedExec("INSERT INTO " + tableName + " VALUES (:name, :type, :rate, :price)", &prod)
		if err != nil {
			fmt.Println("Вставка невозможна.")
		}
	tx.Rollback()
	tx.Commit()
	}


func SelectAllProducts(){
	var prods *[]prod.Product
	tx := Connect.MustBegin()
	_, err := tx.NamedExec("select p.*, t.id from keyboads k join types t on p.type = t.name union select m.*,t.id from mouses m join types t on m.type = t.name", &prods)
		if err != nil {
			fmt.Println("Выборка невозможна.")
		}
	tx.Rollback()
	tx.Commit()
}

func UserSelect(){
	var prods *[]prod.Product
	var typeName string
	var param string
	fmt.Println("Какую выборку хотите осуществить?\nВведите (1) - по типу\n(2) - по рейтингу\n(3) - по цене")
	fmt.Scanf("%s\n", &param)
	switch param {
		case "1":
			fmt.Println("Введите тип продукта(wire/wireless)")
			param = "type = '"
			fmt.Scanf("%s\n", &typeName)
			param = param + typeName + "'"
		case "2":
			fmt.Println("Введите рейтинг")
			fmt.Scanf("%s\n", &typeName)
			fmt.Println("(1) - больше\n(2) - меньше")
			fmt.Scanf("%s\n", &param)
			switch param {
				case "1":
					param = "rate > " + typeName
				case "2":
					param = "rate < " + typeName
			}
		case "3":
			fmt.Println("Введите цену")
			fmt.Scanf("%s\n", &typeName)
			fmt.Println("(1) - больше\n(2) - меньше")
			fmt.Scanf("%s\n", &param)
			switch param {
				case "1":
					param = "price > " + typeName
				case "2":
					param = "price < " + typeName
			}
	}
	tx := Connect.MustBegin()
	_, err := tx.NamedExec("select p.*, t.id from keyboads k join types t on p.type = t.name union select m.*,t.id from mouses m join types t on m.type = t.name where " + param, &prods)
		if err != nil {
			fmt.Println("Выборка невозможна.")
		}
	tx.Rollback()
	tx.Commit()
	fmt.Println(&prods)
}

func UpdDel(){

	tx := Connect.MustBegin()
	errUpd := tx.MustExec("Update keyboards set price = 0 where id = ?", 1)
	if errUpd != nil {
		fmt.Println("Невозможно обновить запись.")
	} else {
		fmt.Println("Запись обновлена.")
	}
	errDel := tx.MustExec("delete from keyboards id = ?", 4 )
	if errDel != nil {
		fmt.Println("Невозможно удалить запись.")
	} else {
		fmt.Println("Запись удалена.")
	}
	tx.Rollback()
	tx.Commit()
}