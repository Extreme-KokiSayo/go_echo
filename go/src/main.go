package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"database/sql" 
	_ "github.com/go-sql-driver/mysql"
)

/**
 * MAIN
 */
func main() {
	var e = createEcho()

	// routes
	e.GET("/", getList)

	e.Logger.Fatal(e.Start(":8080"))
}

/**
 * get echo instance
 */
func createEcho() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

/**
 * 
 */
func getList(c echo.Context) error {
	db := getDb()
	defer db.Close()

	var (
		msg string
		name string

		response string
	)
	q := `select 
 tp.msg, tu.name
 from posts as tp
 join users as tu on tp.user_id=tu.id
 order by tp.updated_at
 limit 5
;`

	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&msg, &name)
		if err != nil {
			log.Fatal(err)
		}

		response += "========" + "\n" + "POST: " + msg + "\n" + "USER: " + name + "\n"
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return c.String(http.StatusOK, response)
}


/*
func testFetch() {
	var (
		id int
		name string
	)

	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
*/

func dbConf() string {
	USER := "myuser"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	TABLE := "sns"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + TABLE + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return CONNECT
}

func getDb() *sql.DB {
	// Setting DB
	dbms := "mysql"
	host := dbConf()

	db, err := sql.Open(dbms, host)
	if err != nil {
		log.Fatal(err)
	}

	return db
}