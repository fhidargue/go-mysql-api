package utils

import (
	"database/sql"
	"fmt"
	"log"
	"mysql-backend/models"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() {
	var err error

	cfg := mysql.Config {
		User: "root",
		Passwd: "root",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "school",
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to MYSQL on port :3306")
}

func GetStudentsDB() ([]models.Student, error) {
	var students []models.Student

	rows, err := db.Query("select * from students")

	if err != nil {
		return nil, fmt.Errorf("there was an issue querying the students")
	}

	defer rows.Close()

	for rows.Next() {
		var student models.Student

		if err := scanStudentRows(rows, &student); err != nil {
			return nil, fmt.Errorf("there was an issue scanning the students")
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("there was an issue while getting the students")
	}

	return students, nil
}

func GetStudentByIdDB(id int64) (models.Student, error) {
	var student models.Student

	row := db.QueryRow("select * from students where id = ?", id)

	if err := scanStudentRow(row, &student); err != nil {
		if err == sql.ErrNoRows {
			return student, fmt.Errorf("student with Id: %v, not found", id)
		}

		return student, fmt.Errorf("student with Id: %v, not found", id)
	}

	return student, nil
}

func AddStudentDB(student models.Student) (int64, error) {
	result, err := db.Exec("insert into students (name, lastName, age, grade) values (?, ?, ?, ?)", 
	student.Name, student.LastName, student.Age, student.Grade)

	if err != nil {
		return 0, fmt.Errorf("error adding student: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error getting the student ID: %v", err)
	}

	return id, nil
}

func UpdateStudentDB(id int64, student models.StudentPatch) error {
	query := "UPDATE students SET "
	var args []interface{}
	
	if student.Name != "" {
		query += "name = ?, "
		args = append(args, student.Name)
	}

	if student.LastName != "" {
		query += "lastName = ?, "
		args = append(args, student.LastName)
	}

	if student.Age != nil {
		query += "age = ?, "
		args = append(args, *student.Age)
	}

	if student.Grade != nil {
		query += "grade = ?, "
		args = append(args, *student.Grade)
	}
	
	if len(args) == 0 {
		return fmt.Errorf("no fields to update")
	}
	
	query = query[:len(query)-2] + " WHERE id = ?"
	args = append(args, id)
	
	_, err := db.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("error updating student with ID %d: %v", id, err)
	}

	return nil
}

func DeleteStudentDB(id int64) (int64, error) {
	_, err := db.Exec("delete from students where id = ?", id)

	if err != nil {
		return 0, fmt.Errorf("error adding student: %v", err)
	}

	return id, nil
}

func scanStudentRows(rows *sql.Rows, student *models.Student) error {
	return rows.Scan(&student.ID, &student.Name, &student.LastName, &student.Age, &student.Grade)
}

func scanStudentRow(row *sql.Row, student *models.Student) error {
	return row.Scan(&student.ID, &student.Name, &student.LastName, &student.Age, &student.Grade)
}