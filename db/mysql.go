package main

import (
  "fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

/*
CREATE TABLE `test`.`test_tables` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(45) CHARACTER SET utf8mb4 DEFAULT NULL,
    `age` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4'

INSERT INTO `test`.`test_tables` (`name`, `age`) VALUES ('Tom', '18');
INSERT INTO `test`.`test_tables` (`name`, `age`) VALUES ('Mary', '24');

CREATE USER 'test'@'localhost' IDENTIFIED BY '123456';
flush privileges;
grant all privileges on test.* to test@localhost identified by '123456';
*/

type Test_Table struct {
  ID int
  Name string
  Age int
}

func sqlselect(keyword string) (int) {
  dsn := "test:123456@(127.0.0.1:33305)/test?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
    return -1
  }

  rs := []Test_Table{}
  like := fmt.Sprintf("%v%v%v", "%", keyword, "%")
  db.Limit(10).Where("name<>'-1' and name LIKE ?", like).Find(&rs)

  for i := 0; i < len(rs); i++ {
    itor := rs[i]
    fmt.Printf("%v:%v\n", itor.Name, itor.Age)
  }
  
  return 0
}

func main() {
  sqlselect("Tom")
}