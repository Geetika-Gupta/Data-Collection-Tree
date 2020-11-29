package util

//PInsert : will process insert web request
func PInsert(req model.WebRequest) (data string, err error) {
	var (
		webReq, timeSpent int
		device, country string
		dim map[string]map[string]int
		met map[string]int
		exists bool
	)
	for _, val := range req.Met {
		if val.Key == "webreq" {
			webReq += val.Value
		} else if val.Key == "timespent" {
			timeSpent += val.Value
		}
	}
	for _, val := range req.Dim {
		if val.Key == "device" {
			device = val.Value
		} else if val.Key == "country" {
			country = val.Value
		}
	}
	if country != "" && device != "" {
		if dim, exists = constant.COUNTRY[country]; !exists {
			dim = make(map[string]map[string]int)
		}
		if met, exists = dim[device]; !exists {
			met = make(map[string]int)
		}
		met["webreq"] += webReq
		met["timespent"] += timeSpent
		dim[device] = met
		constant.COUNTRY[country] = dim
	} else {
		err = perror.CustomError("INVALID REQUEST")
		return
	}
	data = "WEB REQUEST ADDED SUCCESSFULLY!!!"
	return
}

//PGet : will process get web request
func PGet(req model.Dimension) (data model.WebResponse, err error) {
	var country string
	webReq := 0
	timeSpent := 0
	for _, val := range req {
		if val.Key == "country" {
			country = val.Value
		}
	}
	if country == "" {
		err = perror.CustomError("Invalid Country")
		return
	}
	if dim, ex := constant.COUNTRY[country]; ex {
		for _, val := range dim {
			for type, currVal := range val {
				if type == "webreq" {
					webReq += currVal
				} else if type == "timespent" {
					timeSpent += currVal
				}
			}
		}
	} 
	data = make(model.WebResponse{
		Dim: []model.Data{
			{
				Key: "country",
				Value: country,
			},
		},
		Met: []model.Data{
			{
				Key: "webreq",
				Value: webReq,
			},
			{
				Key: "timespent",
				Value: timeSpent,
			},
		}
	})
	return
}