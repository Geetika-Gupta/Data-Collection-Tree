package util

//Response handler.
func BuildResponse(data interface{}, err error, context *gin.Context) {
	var status int
	if err == nil {
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}
	context.JSON(
		status,
		gin.H{
			"status": status,
			"result": data,
			"error":  err,
		},
	)
}
