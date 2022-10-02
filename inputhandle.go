package main

import (
    "fmt"
    "net/http"
    "log"
    "time"
    "strconv"

)

func inputHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****inputHandler running*****")
	
	r.ParseForm()
	iyear := r.FormValue("input_year")
	imonth := r.FormValue("input_month")
	nyear,nmonth,_ := time.Now().Date()
	if iyear == ""{
	 iyear = strconv.Itoa(nyear)
	}
	if imonth == ""{
	 imonth = strconv.Itoa(int(nmonth))
	}
		
    if len(imonth)==1{
 	   imonth = "0"+imonth
    }
    
    data := struct {
             Iyear string
             Imonth string
             }{
	        Iyear: iyear,
	        Imonth: imonth,
	        }

        if err := templates["input"].Execute(w, data); err != nil {
           log.Printf("failed to execute template: %v", err)
        }
}


func inputValueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****inputValueHandler running*****")
	r.ParseForm()
	iyear := r.FormValue("input_year")
	imonth := r.FormValue("input_month")
	iday := r.FormValue("input_day")
	ikind := r.Form["input_kind"]
	iamount := r.FormValue("input_amount")
	icomment := r.FormValue("input_comment")
	
	if ikind[0] == "income"{
	  income_stmt, err := db.Prepare("insert into account_book (year,month,day,income,biko,payment,first_income,register_date,update_date) values($1,$2,$3,$4,$5,0,0,current_timestamp,current_timestamp)")
	  if err != nil {
             log.Fatal(err)
          }
         income_stmt.Exec(iyear,imonth,iday,iamount,icomment)
	}else if ikind[0] == "payment"{
	  pyament_stmt, err := db.Prepare("insert into account_book (year,month,day,payment,biko,income,first_income,register_date,update_date) values($1,$2,$3,$4,$5,0,0,current_timestamp,current_timestamp)")
	  if err != nil {
             log.Fatal(err)
          }
         pyament_stmt.Exec(iyear,imonth,iday,iamount,icomment)
	}
        
        mainHandler(w,r)
}

