DROP TABLE IF EXISTS students;
CREATE TABLE students (
  id         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(128) NOT NULL,
  lastName     VARCHAR(128) NOT NULL,
  age        INT NOT NULL,
  grade      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);