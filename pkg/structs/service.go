package structs

type ServiceStruct struct {
	Id          int      `json:"id"`
	Name        Language `json:"name"`
	Description Language `json:"description"`
	CompanyId   int      `json:"companyId"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}
