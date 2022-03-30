package structs

type ServiceStruct struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CompanyId   int    `json:"companyId"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
