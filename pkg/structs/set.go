package structs

type SetStruct struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ServiceId   int    `json:"serviceId"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
