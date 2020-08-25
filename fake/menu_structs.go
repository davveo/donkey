package fake

type MenuList struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    []MenuListMetaData `json:"data"`
}
type MenuListMetaData struct {
	MenuID        int    `json:"menu_id"`
	ParentID      int    `json:"parent_id"`
	Name          string `json:"name"`
	Alias         string `json:"alias"`
	Icon          string `json:"icon"`
	Remark        string `json:"remark"`
	Module        string `json:"module"`
	Type          int    `json:"type"`
	URL           string `json:"url"`
	Params        string `json:"params"`
	Target        string `json:"target"`
	IsNavi        int    `json:"is_navi"`
	Sort          int    `json:"sort"`
	Status        int    `json:"status"`
	ChildrenTotal int    `json:"children_total"`
	Level         int    `json:"level"`
}
