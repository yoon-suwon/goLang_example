package main

import(
    "strconv"
    "fmt"
    "time"
    "net/http"
    "log"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****mainHandler running*****")

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
    
    slcDate := iyear+imonth


	var (
	first_income_flag bool
	total_payment int
	total_income int
	balance int
	fif int
	)
	first_stmt := "SELECT first_income_flag FROM account_book WHERE year = $1 and month = $2 group by first_income_flag"
	rows,err := db.Query(first_stmt, iyear,imonth)
	 if err != nil {
        fmt.Println(err)
    }
    
      fmt.Println("*iyear,imonth",iyear,imonth,slcDate)
	fif = 0
	for rows.Next() {
	    rows.Scan(&first_income_flag)	    
            if first_income_flag == true{
	       fif = 1
	    }
	}

	if fif != 1 {
	 ifirst_stmt, err := db.Prepare("insert into account_book (year,month,day,first_income,first_income_flag,biko,register_date,update_date) values($1,$2,'01',(select (coalesce(sum(first_income),0) + coalesce(sum(income),0) - coalesce(sum(payment),0)) as first_income from account_book where year || month <= $3 group by year,month order by year,month desc limit 1),true,'frist income',current_timestamp,current_timestamp)")
	  if err != nil {
             log.Fatal(err)
          }
         ifirst_stmt.Exec(iyear,imonth,slcDate) 
	}
	 if err := db.QueryRow("SELECT coalesce(sum(income),0) as total_income, coalesce(sum(payment),0) as total_payment, (coalesce(sum(first_income),0) + coalesce(sum(income),0) - coalesce(sum(payment),0)) as balance FROM account_book WHERE year = $1 and month = $2", iyear, imonth).Scan(&total_income,&total_payment,&balance);err != nil {
	 	     fmt.Println("*mainerror 2")
    log.Fatal(err)
}
        data := struct {
                Iyear string
                Imonth string
		Income   int
		Payment int
		Balance    int
	}{
	        Iyear: iyear,
	        Imonth: imonth,
		Income: total_income,
		Payment: total_payment,
		Balance:    balance,
	}

	
    if err := templates["main"].Execute(w, data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }
}
