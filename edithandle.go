package main

import (
    "fmt"
    "net/http"
    "log"
    "time"
    "strconv"
    "math"
)

func openEditHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****openEditHandler running*****")
	r.ParseForm()
	ibookId := r.FormValue("bookId")
	fmt.Println("*ibookId :",ibookId)
	var (
	bookId int
	year string
	month string
	day string
	income int
	payment int
	first_income int
	comment string
	)
         if err := db.QueryRow("SELECT account_book_id,year,month,day,coalesce(income,0) as income,coalesce(payment,0) as payment,coalesce(first_income,0) as first_income,biko FROM account_book WHERE account_book_id = $1", ibookId).Scan(&bookId,&year,&month,&day,&income,&payment,&first_income,&comment);err != nil {
    log.Fatal(err)}
         
         data := struct {
          BookId int
          Year   string
          Month  string
          Day    string 
          Income  int 
          Payment int
          First_income int
          Comment string
         }{
          BookId: bookId,
          Year:   year,
          Month:  month,
          Day:    day, 
          Income:  income, 
          Payment: payment,
          First_income: first_income,
          Comment: comment,
         }
         

        if err := templates["edit"].Execute(w, data); err != nil {
           log.Printf("failed to execute template: %v", err)
        }

}

func editValueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****editValueHandler running*****")
	r.ParseForm()
	ibookId := r.FormValue("bookId")
	iyear := r.FormValue("input_year")
	imonth := r.FormValue("input_month")
	iday := r.FormValue("input_day")
	iincome := r.FormValue("input_income")
	ipayment := r.FormValue("input_payment")
	ifirst := r.FormValue("input_first")
	icomment := r.FormValue("input_comment")
	
	_ = iday
	
	var (
	bookId int
	year string
	month string
	day string
	income int
	payment int
	first_income int
	comment string
	)
	
	niyear,nimonth,niday := time.Now().Date()
	 nyear := strconv.Itoa(niyear)
	 nmonth := strconv.Itoa(int(nimonth))
	 nday := strconv.Itoa(niday)
    if len(nmonth)==1{
 	   nmonth = "0"+nmonth
    }
    if len(nday)==1{
 	   nday = "0"+nday    
    }
    
    if iyear+imonth < nyear+nmonth{
	
	    if err := db.QueryRow("SELECT account_book_id,year,month,day,coalesce(income,0) as income,coalesce(payment,0) as payment,coalesce(first_income,0) as first_income FROM account_book WHERE account_book_id = $1", ibookId).Scan(&bookId,&year,&month,&day,&income,&payment,&first_income);err != nil {
    	log.Fatal(err)}
    	
    	cifirst,_ := strconv.Atoi(ifirst)
    	ciincome,_ := strconv.Atoi(iincome)
    	cipayment,_ := strconv.Atoi(ipayment)
    	
    	
    	cresult :=(cifirst+ciincome-cipayment)-(first_income+income-payment)    	
    	if cresult<0{

    		cpayment := math.Abs(float64(cresult))
    		comment = year+"/"+month+"/"+day+"の差分登録"
	  		pyament_stmt, err := db.Prepare("insert into account_book (year,month,day,payment,biko,income,first_income,register_date,update_date) values($1,$2,$3,$4,$5,0,0,current_timestamp,current_timestamp)")
	  		if err != nil {
            	log.Fatal(err)
          	}
			pyament_stmt.Exec(nyear,nmonth,nday,cpayment,comment)
       	}else{
    		income = cresult
     		comment = year+"/"+month+"/"+day+"の差分登録"
			income_stmt, err := db.Prepare("insert into account_book (year,month,day,income,biko,payment,first_income,register_date,update_date) values($1,$2,$3,$4,$5,0,0,current_timestamp,current_timestamp)")
	  		if err != nil {
            	log.Fatal(err)
        	}
         	income_stmt.Exec(nyear,nmonth,nday,income,comment)
    	}    	
    }
	
	stmt, err := db.Prepare("update account_book set income = $1, payment =$2,first_income = $3, biko =$4, update_date = current_timestamp where account_book_id = $5 ;")
	  if err != nil {
             log.Fatal(err)
          }
         stmt.Exec(iincome,ipayment,ifirst,icomment,ibookId)

        historyHandler(w,r)

}

func valueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****valueDeleteHandler running*****")
	r.ParseForm()
	ibookId := r.FormValue("bookId")
	
	stmt, err := db.Prepare("delete from account_book where account_book_id = $1")
	  if err != nil {
             log.Fatal(err)
          }
         stmt.Exec(ibookId)
    
        historyHandler(w,r)
}

