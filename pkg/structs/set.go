package structs

type SetStruct struct {
	Id          int      `json:"id"`
	Name        Language `json:"name"`
	Description Language `json:"description"`
	ServiceId   int      `json:"serviceId"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}

type CreateSet struct {
	Name        Language `json:"name"`
	Description Language `json:"description"`
	ServiceId   int      `json:"serviceId"`
}
