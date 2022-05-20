package structs

type PackageStruct struct {
	Id          int      `json:"id"`
	Name        Language `json:"name"`
	Description Language `json:"description"`
	Capacity    int      `json:"capacity"`
	Cost        int      `json:"cost"`
	Code        string   `json:"code"`
	SetId       int      `json:"setId"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}

type CreatePackage struct {
	Name        Language `json:"name"`
	Description Language `json:"description"`
	Capacity    int      `json:"capacity"`
	Cost        int      `json:"cost"`
	Code        string   `json:"code"`
	SetId       int      `json:"setId"`
}
