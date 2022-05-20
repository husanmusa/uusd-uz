package repo

import "github.com/husanmusa/uusd-uz/pkg/structs"

type CompanyRepoInterface interface {
	CreateCompany(company structs.CreateCompany) (structs.CompanyStruct, error)
	GetCompany(id int) (structs.CompanyStruct, error)
	GetListCompanies() ([]structs.CompanyStruct, error)
	UpdateCompany(company structs.CompanyStruct) (structs.CompanyStruct, error)
	DeleteCompany(id int) error
}

type ServiceRepoInterface interface {
	CreateService(service structs.CreateService) (structs.ServiceStruct, error)
	GetService(id int) (structs.ServiceStruct, error)
	GetListServices(companyId int) ([]structs.ServiceStruct, error)
	UpdateService(service structs.ServiceStruct) (structs.ServiceStruct, error)
	DeleteService(id int) error
}

type SetRepoInterface interface {
	CreateSet(set structs.CreateSet) (structs.SetStruct, error)
	GetSet(id int) (structs.SetStruct, error)
	GetListSets(serviceId int) ([]structs.SetStruct, error)
	UpdateSet(set structs.SetStruct) (structs.SetStruct, error)
	DeleteSet(id int) error
}

type PackageRepoInterface interface {
	CreatePackage(pack structs.CreatePackage) (structs.PackageStruct, error)
	GetPackage(id int) (structs.PackageStruct, error)
	GetListPackages(setId int) ([]structs.PackageStruct, error)
	UpdatePackage(pack structs.PackageStruct) (structs.PackageStruct, error)
	DeletePackage(id int) error
}
