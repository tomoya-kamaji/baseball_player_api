package http

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Return200 is  is a helper function to return status code 200
func Return200[T any](ctx *gin.Context, body T) {
	ctx.JSON(200, &body)
}

// Return201 is  is a helper function to return status code 201
func Return201[T any](ctx *gin.Context, body T) {
	ctx.JSON(201, &body)
}

// Return204 is  is a helper function to return status code 204
func Return204(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

// Return400 is  is a helper function to return status code 400 and error
func Return400(ctx *gin.Context, err error) {
	returnAbortWith(ctx, 400, err)
}

// ReturnBadRequest is  is a helper function to return bad request
func ReturnBadRequest(ctx *gin.Context, err error) {
	Return400(ctx, err)
}

// Return401 is  is a helper function to return status code 401
func Return401(ctx *gin.Context, err error) {
	returnAbortWith(ctx, 401, err)
}

// ReturnUnauthorized is  is a helper function to return unauthorized
func ReturnUnauthorized(ctx *gin.Context, err error) {
	Return401(ctx, err)
}

// Return403 is  is a helper function to return status code 403
func Return403(ctx *gin.Context, err error) {
	returnAbortWith(ctx, 403, err)
}

// ReturnForbidden is  is a helper function to return forbidden
func ReturnForbidden(ctx *gin.Context, err error) {
	Return403(ctx, err)
}

// Return404 is  is a helper function to return status code 404
func Return404(ctx *gin.Context, err error) {
	returnAbortWith(ctx, 404, err)
}

// ReturnNotFound is  is a helper function to return not found
func ReturnNotFound(ctx *gin.Context, err error) {
	Return404(ctx, err)
}

// Return500 is  is a helper function to return status code 500
func Return500(ctx *gin.Context, err error) {
	returnAbortWith(ctx, 500, err)
}

// ReturnCSV is  is a helper function to return CSV file
func ReturnCSV(ctx *gin.Context, fileName string, contents [][]string) {
	// 書き込むCSVデータ
	csvBytes := &bytes.Buffer{}
	bw := bufio.NewWriter(csvBytes)

	// BOMをつける
	if _, err := bw.Write([]byte{0xEF, 0xBB, 0xBF}); err != nil {
		Return500(ctx, err)
		return
	}
	writer := csv.NewWriter(bw)

	// データの書き込み
	if err := writer.WriteAll(contents); err != nil {
		Return500(ctx, err)
		return
	}
	writer.Flush()
	result := csvBytes.Bytes()

	ctx.Header("Content-Disposition", "attachment; filename="+fileName+".csv")
	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Length", strconv.Itoa(len(result)))
	if _, err := ctx.Writer.Write(result); err != nil {
		Return500(ctx, err)
	}
}

// ParseQueryOpt is a helper function to parse query
func ParseQueryOpt(ctx *gin.Context, key string) *string {
	if str := ctx.Query(key); str != "" {
		return &str
	}
	return nil
}

// ParseID is a helper function to parse parameters
func ParseID(ctx *gin.Context, key string) int64 {
	if id, err := strconv.Atoi(ctx.Param(key)); err == nil {
		return int64(id)
	}
	return 0
}

// ValidateBindQuery is a helper function to BindQuery obj and ValidateStruct obj
func ValidateBindQuery(ctx *gin.Context, obj interface{}) error {
	return ctx.BindQuery(obj)
}

// ValidateBindJSON is a helper function to BindJSON obj and ValidateStruct obj
func ValidateBindJSON(ctx *gin.Context, obj interface{}) error {
	return ctx.BindJSON(obj)
}

func GetHeader(ctx *gin.Context, key string) string {
	return ctx.Request.Header.Get(key)
}

func returnAbortWith(ctx *gin.Context, code int, err error) {
	var msg string
	if err != nil {
		msg = err.Error()
	}

	ctx.AbortWithStatusJSON(code, gin.H{
		"code": code,
		"msg":  msg,
	})
}
