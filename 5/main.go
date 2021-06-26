package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
)

type Teacher struct {
	TeacherName string `json:"teachername"`
}

type JsonResponse struct {
	Type    string    `json:"type"`
	Data    []Teacher `json:"data"`
	Message string    `json:"message"`
}

func main() {
	router := mux.NewRouter()

	// Get all teacher
	router.HandleFunc("/teacher/", GetTeacher).Methods("GET")

	// Create a teacher data
	router.HandleFunc("/teacher/", CreateTeacherData).Methods("POST")

	// Update a teacher data
	router.HandleFunc("/updateteacher/", UpdateTeacherData).Methods("POST")

	// Delete a specific teacher data by the id
	router.HandleFunc("/teacher/{teacherid}", DeleteTeacher).Methods("DELETE")

	// Search teacher data
	router.HandleFunc("/teacher/{filter}", SearchTeacher).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", router))
}

// Get all teacher data
func GetTeacher(w http.ResponseWriter, r *http.Request) {
	simpleAuth := "KillingOfaSacredDeer"
	ua := r.Header.Get("Authorization")
	if simpleAuth != ua {
		var response = JsonResponse{Type: "failure", Message: "Admin not recognized"}

		json.NewEncoder(w).Encode(response)
	} else {
		db := setupDB()

		printMessage("Getting teachers data...")

		rows, err := db.Query("SELECT * FROM teacher_test")

		checkErr(err)
		var teacher []Teacher
		for rows.Next() {
			var teacher_id int
			var teacherName string

			err = rows.Scan(&teacher_id, &teacherName)

			checkErr(err)

			teacher = append(teacher, Teacher{TeacherName: teacherName})
		}

		var response = JsonResponse{Type: "success", Data: teacher}

		json.NewEncoder(w).Encode(response)
	}
}

// Create a teacher data
func CreateTeacherData(w http.ResponseWriter, r *http.Request) {
	simpleAuth := "KillingOfaSacredDeer"
	ua := r.Header.Get("Authorization")
	if simpleAuth != ua {
		var response = JsonResponse{Type: "failure", Message: "Admin not recognized"}

		json.NewEncoder(w).Encode(response)
	} else {
		teacherName := r.FormValue("teachername")

		var response = JsonResponse{}

		if teacherName == "" {
			response = JsonResponse{Type: "error", Message: "You are missing teachername parameter."}
		} else {
			db := setupDB()

			printMessage("Inserting teacher data into DB")

			fmt.Println("Inserting new teacher with name: " + teacherName)

			var lastInsertID int
			err := db.QueryRow("INSERT INTO teacher_test(teacher_name) VALUES($1) returning teacher_id;", teacherName).Scan(&lastInsertID)

			checkErr(err)

			response = JsonResponse{Type: "success", Message: "The data has been inserted successfully!"}
		}

		json.NewEncoder(w).Encode(response)
	}
}

// Update teacher data
func UpdateTeacherData(w http.ResponseWriter, r *http.Request) {
	simpleAuth := "KillingOfaSacredDeer"
	ua := r.Header.Get("Authorization")
	if simpleAuth != ua {
		var response = JsonResponse{Type: "failure", Message: "Admin not recognized"}

		json.NewEncoder(w).Encode(response)
	} else {
		teacherName := r.FormValue("teachername")
		teacherId := r.FormValue("teacherid")
		fmt.Println(teacherName)
		fmt.Println(string(teacherId))
		var response = JsonResponse{}

		if teacherName == "" {
			response = JsonResponse{Type: "error", Message: "You are missing teachername parameter."}
		} else {

			db := setupDB()

			printMessage("Inserting teacher data into DB")

			fmt.Println("Inserting new teacher with name: " + teacherName)

			res, err := db.Exec("UPDATE teacher_test SET teacher_name = $1 where teacher_id = $2;", teacherName, teacherId)

			checkErr(err)
			fmt.Println(res.RowsAffected)

			response = JsonResponse{Type: "success", Message: "The data has been changed successfully!"}
		}

		json.NewEncoder(w).Encode(response)
	}
}

// Delete a teacher data
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	simpleAuth := "KillingOfaSacredDeer"
	ua := r.Header.Get("Authorization")
	if simpleAuth != ua {
		var response = JsonResponse{Type: "failure", Message: "Admin not recognized"}

		json.NewEncoder(w).Encode(response)
	} else {
		params := mux.Vars(r)

		teacherID := params["teacherid"]

		var response = JsonResponse{}

		if teacherID == "" {
			response = JsonResponse{Type: "error", Message: "You are missing teacherid parameter."}
		} else {
			db := setupDB()

			printMessage("Deleting teacher data from DB")

			_, err := db.Exec("DELETE FROM teacher_test where teacher_id = $1", teacherID)
			checkErr(err)

			response = JsonResponse{Type: "success", Message: "The teacher data has been deleted successfully!"}
		}

		json.NewEncoder(w).Encode(response)
	}
}

// Get filtered data
func SearchTeacher(w http.ResponseWriter, r *http.Request) {
	simpleAuth := "KillingOfaSacredDeer"
	ua := r.Header.Get("Authorization")
	if simpleAuth != ua {
		var response = JsonResponse{Type: "failure", Message: "Admin not recognized"}

		json.NewEncoder(w).Encode(response)
	} else {
		params := mux.Vars(r)

		filter := strings.ToLower(params["filter"])

		db := setupDB()

		printMessage("Getting teachers data...")
		printMessage(filter)

		rows, err := db.Query("SELECT * FROM teacher_test WHERE LOWER(teacher_name) LIKE '%" + filter + "%'")

		checkErr(err)
		var teacher []Teacher
		for rows.Next() {
			var teacher_id int
			var teacherName string

			err = rows.Scan(&teacher_id, &teacherName)

			checkErr(err)

			teacher = append(teacher, Teacher{TeacherName: teacherName})
		}

		var response = JsonResponse{Type: "success", Data: teacher}

		json.NewEncoder(w).Encode(response)
	}
}

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
