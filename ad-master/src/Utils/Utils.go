package Utils

import (
	b64 "encoding/base64"
	"fmt"

)

func EncodeBase64(input string) string {
	sEnc := b64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println(sEnc)
	return sEnc
}

func DecodeBase64(input string) string {

    sDec, _ := b64.StdEncoding.DecodeString(input)
    fmt.Println(string(sDec))
    fmt.Println()
	return string(sDec)
}


// func ConvertTimestampToDate(input string) string {

// 	i, err := strconv.ParseInt(input, 10, 64)
//     if err != nil {
// 		panic(err)
// 	}
// 	tm := time.Unix(i, 0)
// 	fmt.Println(tm)
// 	return string(tm.Format("2006-01-02 15:04:05"))

// }
