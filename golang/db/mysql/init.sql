CREATE SCHEMA `test` DEFAULT CHARACTER SET utf8mb4 ;

CREATE TABLE `test`.`users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(45) CHARACTER SET utf8mb4 DEFAULT NULL,
    `age` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

INSERT INTO `test`.`users` (`name`, `age`) VALUES ('Natalia', '20');
INSERT INTO `test`.`users` (`name`, `age`) VALUES ('Tom', '21');
INSERT INTO `test`.`users` (`name`, `age`) VALUES ('Mary', '22');
INSERT INTO `test`.`users` (`name`, `age`) VALUES ('Jack', '23');

CREATE USER 'test'@'%' IDENTIFIED BY '123456';
GRANT ALL PRIVILEGES ON test.* TO 'test'@'%';
flush privileges;
