package Utils

import (
	"time"
	"fmt"
	"strconv"
)

func ConvertTimestampToDate(input int64) string {
	tm := time.Unix(input, 0)
	fmt.Println(tm)
	return string(tm.Format("2006-01-02 15:04:05"))
}


	
func ShouldResetPassword(psw_created_date string) bool{
 
	i, err := strconv.ParseInt(psw_created_date, 10, 64)
    if err != nil {
		panic(err)
	}
	if(time.Now().Unix()-i > 0){
      	return false
	} else {
	    return true
	}
}


