package structs

type CompanyStruct struct {
	Id     int      `json:"id"`
	Name   Language `json:"name"`
	Cover  string   `json:"cover"`
	Slogan Language `json:"slogan"`
	//Services  ServiceStruct `json:"serviceStruct"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
