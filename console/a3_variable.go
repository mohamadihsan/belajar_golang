package main

import "fmt"

func main() {
   // deklarasi type 1 menggunakan var
   var userName string = "Mohamad Ihsan"
   // deklarasi type 2 menggunaka var
   var year int
   year = 2019
   // deklarasi multiple variable menggunakan var
   var pembuka, penutup string = "Hallo,", "Masehi"
   // deklarasi tanpa menggunakan var
   detail := "Anda seorang developer"
   /* 
   setiap variable harus digunakan, jika ada variable yang tidak digunakan
   maka akan error. Jika ada variable yg tidak akan digunakan terlebih dahulu
   maka bisa diatasi dengan menampung value pada variable tersebut 
   dengan menggunakan underscore
   */
   // title := "Pengenalan Dasar" // ini akan error
   _ = "Pengenalan Dasar" // ini akan berjalan normal

   fmt.Printf("%s %s sekarang tahun %s %s. %s",
   pembuka,
   userName,
   year,
   penutup,
   detail)
}
