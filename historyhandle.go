package main
 
import (
    "fmt"
    _ "github.com/lib/pq"
    "net/http"
    "log"
    "strconv"
    "time"
)


func historyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****historyHandler running*****")

	//html?? ??????
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
    
	stmt := "SELECT (year||'/'||month||'/'||day) as date,coalesce(income,0) as income,coalesce(payment,0) as payment,coalesce(first_income,0) as first_income,coalesce(biko,'') as biko, account_book_id FROM account_book WHERE year = $1 and month = $2 order by date,account_book_id"
	rows,err := db.Query(stmt, iyear,imonth)
	 if err != nil {
        fmt.Println(err)
    }
        type History struct {
         Date    string 
         Income  int 
         Payment int
         First_income int
         Comment string
         BookId  int
         }
         type Output struct{
          Iyear string
          Imonth string
          Hist []History
         }
       
        var his []History
        var mst History
        for rows.Next() {
	    err = rows.Scan(&mst.Date,&mst.Income,&mst.Payment,&mst.First_income,&mst.Comment,&mst.BookId)	    
            his = append(his,mst)

         }
        var out Output
        out.Iyear = iyear
        out.Imonth = imonth
        out.Hist = his 
	fmt.Println(out)
	
       if err := templates["history"].Execute(w, out); err != nil {
           log.Printf("failed to execute template: %v", err)
        }
}

