package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Numbers []int `json:"numbers" binding:"required"`
	Target  int   `json:"target" binding:"required"`
}

type Response struct {
	Solutions [][]int `json:"solutions"`
}



type Trg struct {
	Numbers []int `json:"numbers" binding:"required"`
	Target  int   `json:"target" binding:"required"`
}

func TargetSumHandler(ctx *gin.Context) {

	var req Trg

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}


	result := TargetSum(ctx, req.Numbers, req.Target)

	rsp:=gin.H{
		"solutions":result,
	}

	ctx.JSON(http.StatusOK,rsp)

}

func  TargetSum(gctx *gin.Context, numbers []int, traget int) [][]int {

	pMap := make(map[int]int)
	var result [][]int

	for i, nm := range numbers {

		if ind, found := pMap[traget-nm]; found {
			result = append(result, []int{ind, i})
		}
		pMap[nm] = i
	}

	fmt.Println("result", result)

	return result
}

func main() {
	r := gin.Default()


	//http://localhost:8080/find-pairs    
	//use the above url
	 

	r.POST("/find-pairs", TargetSumHandler)

	r.Run(":8080")
}
