package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)


type getAll_request struct {
	Price_sort string
	Date_sort string
	Offset int
}

type getOne_request struct {
	Id int
	Fields string
}

type setOne_request struct {
	Name string
	Link string
	Price float64
	Description string
}

type getAll_responce struct {
	Name string
	Link string
	Price float64
}

type getOne_responce struct {
	Name string
	Link string
	Price float64
	Description string
}

type setOne_responce struct {
	Id int64
	Status int
}



func main(){
	handler:=http.NewServeMux()

	handler.HandleFunc("/ad/getall", GetAll)
	handler.HandleFunc("/ad/getone", GetOne)
	handler.HandleFunc("/ad/setone", SetOne)

	fmt.Println("RUNNING")
	s:= http.Server{
		Addr: ":8080",
		Handler: handler,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}

func GetAll (w http.ResponseWriter, r *http.Request){
	var p getAll_request
	p.Offset=-1
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.Write([]byte("400"))
		return
	}
	if p.Offset<=-1 {
		w.Write([]byte("400"))
		return
	}

	db, err := sql.Open("mysql", "root:root@/?parseTime=true")
	defer db.Close()
	if err != nil {
		w.Write([]byte("500"))
		return
	}
	base_request:="select name,link,price from testdb.ads"
	if(p.Price_sort!="" && p.Date_sort!="") {
		w.Write([]byte("400"))
		return
	}
	if(p.Price_sort!="") {
			base_request+=" ORDER BY price"
			if p.Price_sort=="desc" {
				base_request+=" DESC"
			} else if p.Price_sort=="asc" {
				base_request+=" ASC"
			} else {
				w.Write([]byte("400"))
				return
			}
	}

	if p.Date_sort!="" {
			base_request+=" ORDER BY created_at"
			if p.Date_sort=="desc" {
				base_request+=" DESC"
			} else if p.Date_sort=="asc" {
				base_request+=" ASC"
			} else {
				w.Write([]byte("400"))
				return
			}
	}
	base_request+=" LIMIT 10 OFFSET " + strconv.Itoa(p.Offset)
	rows, err := db.Query(base_request)
	defer rows.Close()
	if err != nil {
		w.Write([]byte("500"))
		return
	}


	ads := []getAll_responce{}
	for rows.Next(){
		var adder getAll_responce
		err := rows.Scan(&adder.Name, &adder.Link, &adder.Price)
		if err != nil{
			fmt.Println(err)
			continue
		}
		links:=strings.Split(adder.Link,",")
		adder.Link=links[0]
		ads = append(ads, adder)
	}

	json_data, err := json.Marshal(ads)
	w.Write(json_data)
}

func GetOne (w http.ResponseWriter, r *http.Request){
	var p getOne_request
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.Write([]byte("400"))
		return
	}
	if p.Id<1 {
		w.Write([]byte("400"))
		return
	}

	db, err := sql.Open("mysql", "root:root@/testdb")
	defer db.Close()
	if err != nil {
		w.Write([]byte("500"))
		return
	}

	base_request:="select name,link,price from testdb.ads WHERE id=" + strconv.Itoa(p.Id)
	if(p.Fields!=""){
		if(p.Fields!="true"){
			w.Write([]byte("400"))
			return
		}
		base_request ="select name,link,price,description from testdb.ads WHERE id=" + strconv.Itoa(p.Id)
	}

	rows, err := db.Query(base_request)
	defer rows.Close()
	if err != nil {
		w.Write([]byte("500"))
		return
	}


	if(p.Fields=="") {
		ads := []getAll_responce{}
		for rows.Next(){
			var adder getAll_responce
			err := rows.Scan(&adder.Name, &adder.Link, &adder.Price)
			if err != nil{
				fmt.Println(err)
				continue
			}
			links:=strings.Split(adder.Link,",")
			adder.Link=links[0]
			ads = append(ads, adder)
		}

		json_data, _ := json.Marshal(ads)
		w.Write(json_data)
	} else {
		ads := []getOne_responce{}
		for rows.Next(){
			var adder getOne_responce
			err := rows.Scan(&adder.Name, &adder.Link, &adder.Price,&adder.Description)
			if err != nil{
				fmt.Println(err)
				continue
			}
			ads = append(ads, adder)
		}

		json_data, _ := json.Marshal(ads)
		w.Write(json_data)
	}

}

func SetOne (w http.ResponseWriter, r *http.Request) {
	var p setOne_request
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.Write([]byte("400"))
		return
	}
	db, err := sql.Open("mysql", "root:root@/testdb")
	defer db.Close()
	if err != nil {
		w.Write([]byte("500"))
		panic(err)
	}
	if (p.Name=="") || (p.Link=="") || (p.Price<=0) || (p.Description==""){
		w.Write([]byte("400"))
		return
	}
	links:=strings.Split(p.Link,",")
	if (len(links) > 3) || (len(p.Description)>1000) || (len(p.Name)>200){
		w.Write([]byte("400"))
		return
	}

	t := time.Now()
	ts := t.Format("2006-01-02 15:04:05")

	result , err1 := db.Exec("insert into testdb.ads (name,link,price,description,created_at) values (?,?,?,?,?)",p.Name,p.Link,p.Price,p.Description,ts)
	if err1 != nil{
		w.Write([]byte("500"))
		return
	}

	server_responce:=setOne_responce{}
	server_responce.Id, _ = result.LastInsertId()
	server_responce.Status = 1

	json_data, _ := json.Marshal(server_responce)
	w.Write(json_data)
}













