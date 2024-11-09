
package main

import (
"fmt"
"github.com/gin-gonic/gin"
"net/http"

) 



func main() {
   route:=gin.Default()

   route.GET("/",Firstroute)


   route.Run(":8080")
   

}


func Firstroute(c *gin.Context){
   fmt.Println("first function reached  !!!!!!")

   data:="first function reached  !!!!!!"

   c.JSON(http.StatusOK, gin.H{"message": data})

   return
}
