package repo

import "github.com/husanmusa/uusd-uz/pkg/structs"

type CompanyRepoInterface interface {
	CreateCompany(company structs.CompanyStruct) (structs.CompanyStruct, error)
	GetCompany(id int) (structs.CompanyStruct, error)
	GetListCompany() ([]structs.CompanyStruct, error)
	UpdateCompany(company structs.CompanyStruct) (structs.CompanyStruct, error)
	DeleteCompany(id int) error
}

type ServiceRepoInterface interface {
	CreateService(service structs.ServiceStruct) (structs.ServiceStruct, error)
	GetService(id int) (structs.ServiceStruct, error)
	GetListService() ([]structs.ServiceStruct, error)
	UpdateService(service structs.ServiceStruct) (structs.ServiceStruct, error)
	DeleteService(id int) error
}

type SetRepoInterface interface {
	CreateSet(set structs.SetStruct) (structs.SetStruct, error)
	GetSet(id int) (structs.SetStruct, error)
	GetListSet() ([]structs.SetStruct, error)
	UpdateSet(set structs.SetStruct) (structs.SetStruct, error)
	DeleteSet(id int) error
}

type PackageRepoInterface interface {
	CreatePackage(pack structs.PackageStruct) (structs.PackageStruct, error)
	GetPackage(id int) (structs.PackageStruct, error)
	GetListPackage() ([]structs.PackageStruct, error)
	UpdatePackage(pack structs.PackageStruct) (structs.PackageStruct, error)
	DeletePackage(id int) error
}
