package model

type APIResponse struct {
	Status    bool        `json:"status" example:"false" enums:"true,false"`
	ErrorCode string      `json:"errorcode" example:"CE10001"`
	Message   string      `json:"message" example:"INTERNAL SERVER ERROR"`
	Data      interface{} `json:"data"`
}

type APISuccessResponse struct {
	Status bool        `json:"status" example:"false" enums:"true,false"`
	Data   interface{} `json:"data"`
}

type APIFailureResponse struct {
	Status    bool   `json:"status" example:"false" enums:"true,false"`
	ErrorCode string `json:"errorcode" example:"CE10001"`
	Message   string `json:"message" example:"INTERNAL SERVER ERROR"`
}

type HealthApiResponse struct {
	Resource string `json:"resource"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

type BoredApiRequest struct {
	UserId string `json:"user_id"`
}

type BoredApiResponse struct {
	Key      string `json:"key"`
	Activity string `json:"activity"`
}

type SchedulerResponse struct {
	Count int    `json:"count"`
	Key   string `json:"key"`
}
