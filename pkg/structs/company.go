package structs

type CompanyStruct struct {
	Id        int           `json:"id"`
	Name      string        `json:"name"`
	Cover     string        `json:"cover"`
	Slogan    string        `json:"slogan"`
	Services  ServiceStruct `json:"serviceStruct"`
	CreatedAt string        `json:"createdAt"`
	UpdatedAt string        `json:"updatedAt"`
}
