package main

import (
	"fmt" 
	"time"
)

func main(){
	now := time.Now()

	fmt.Println("Today : ", now.Format("2020, April 27"))

	longTimeAgo := time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC)

	different := now.Sub(longTimeAgo)

	days := int(different.Hours() / 24 )

	fmt.Printf("Dari tanggal 1 Maret 2020 sudah %d hari yang terlewati \n", days)
}