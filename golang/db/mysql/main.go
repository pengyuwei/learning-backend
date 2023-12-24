package main

import (
	"database/sql"
	"fmt"
	"time"

	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       int       `gorm:"column:id"`
	Name     string    `gorm:"column:name"`
	Age      int       `gorm:"column:age"`
	CreateAt time.Time `json:"qadate" gorm:"column:create_at"`
}
type Config struct {
	Host     string `json:"Host" gorm:"default:testdb.host"`
	Port     int
	User     string
	Password string
	Db       string `json:"Host" gorm:"default:null"`
}

func getConn() (*gorm.DB, error) {
	var err error
	var db_conn *gorm.DB = nil
	cfg := Config{Host: "testdb.host", Port: 3306, User: "test", Password: "123456", Db: "test"}

	dsn := fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Db)
	fmt.Printf("dsn=%v\n", dsn)
	// dsn := "test:123456@(testdb.host:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db_conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db_conn = db_conn.Debug()

	return db_conn, err
}

func gorm_example() {
	db_conn, err := getConn()
	if err != nil {
		panic("failed to connect database")
		return
	}

	var user User
	var users []User

	// Create
	fmt.Println("\n--------------------Create--------------------")
	newUser := User{Name: "0xFF", Age: 42}
	db_conn.Create(&newUser)
	newUsers := []User{{Name: "Hapless", Age: 24}, {Name: "The Cooler", Age: 25}, {Name: "A poor fish", Age: 26}}
	db_conn.Create(&newUsers)

	// Delete
	fmt.Println("\n--------------------Delete--------------------")
	db_conn.Delete(&user, 24)
	db_conn.Delete(&User{}, "age = ?", 25)
	db_conn.Where("age = ?", 26).Delete(&User{})

	// Update
	fmt.Println("\n--------------------Update--------------------")
	db_conn.Model(&user).Update("Age", 18) // Warning: WHERE conditions required
	db_conn.Model(&user).Where("age > ?", 42).Updates(User{Name: "0xFF", Age: 81})
	db_conn.Table("user").Model(&user).Where("age > ?", 42).Updates(map[string]interface{}{"Name": "0xFF", "Age": 18})
	db_conn.Table("user").Select("name, age").Order("id desc").First(&user, "age = ?", user.Age)
	ret := db_conn.Model(&user).Updates(User{Name: "NewUser", Age: 81}) // Error if no record found
	if ret.Error != nil {
		fmt.Println("Update can't create new record:", ret.Error)
	}

	// Read
	fmt.Println("\n--------------------Read--------------------")
	rs := db_conn.First(&user, 1) // find user with id 1
	if rs.Error != nil {
		if errors.Is(rs.Error, gorm.ErrRecordNotFound) {
			fmt.Println("No record found")
		} else {
			fmt.Println("Error occurred:", rs.Error)
		}
		return
	}
	fmt.Printf("user=%v\n", user)

	db_conn.First(&user, "name = ?", "Tom") // find user with name Tom
	fmt.Printf("user=%s\n", user.Name)
	db_conn.Where("name = ?", "Tom").First(&user) // find user with name Tom
	db_conn.First(&user, "age > ?", 20)           // find first user with age bigger than 20
	db_conn.Find(&users, "name = ?", "Tom")       // find all users with name Tom
	db_conn.Where("name <> ?", "Tom").Find(&users)
	db_conn.Find(&users) // find all users
	for _, user := range users {
		fmt.Printf("users=%v\n", user)
	}

	db_conn.Table("test.users").First(&user, "DATE(create_at) = DATE(?)", "2023-12-23")
	fmt.Println("test datetime: ", user)
}

func sql_example() {
	db_conn, _ := getConn()
	var sqldb *sql.DB
	var err error
	sqldb, err = db_conn.DB()

	rows, err := sqldb.Query("SELECT id,name,age FROM users WHERE id = ?", 1)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id=%v, name=%s, age=%v\n", id, name, age)
	}

	result, err := sqldb.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Hapless", 24)
	if err != nil {
		panic(err)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("lastInsertID=%v\n", lastInsertID)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("rowsAffected=%v\n", rowsAffected)

	result, err = sqldb.Exec("UPDATE users SET age = ? WHERE id = ?", 24, 99)
	if err != nil {
		panic(err)
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("rowsAffected=%v\n", rowsAffected)

	result, err = sqldb.Exec("DELETE FROM users WHERE id = ?", 1)
	if err != nil {
		panic(err)
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("rowsAffected=%v\n", rowsAffected)
}

func main() {
	gorm_example()
	sql_example()
}
