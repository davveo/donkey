package fake

type AuthGroupList struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		GroupID     int    `json:"group_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Module      string `json:"module"`
		System      int    `json:"system"`
		Sort        int    `json:"sort"`
		Status      int    `json:"status"`
	} `json:"data"`
}
