package user

//InitializeRoutes defines endpoints.
func InitializeRoutes(router *gin.Engine) {
	user := router.Group("/v1")
	{
		user.POST("/insert", insert())
		user.POST("/query", get())
	}
}

//will insert new dimension
func insert() func(context *gin.Context) {
	return func(context *gin.Context) {
		var (
			data string
			req model.WebRequest
			err error
		)
		if req, err = util.VInsert(context); err == nil {
			data, err = util.PInsert(req)
		}
		util.BuildResponse(data, err, context)
	}
}

//will get metrics for given country
func get() func(context *gin.Context) {
	return func(context *gin.Context) {
		var (
			data string
			req model.Dimension
			err error
		)
		if req, err = util.VGet(context); err == nil {
			data, err = util.PGet(req)
		}
		util.BuildResponse(data, err, context)
	}
}
