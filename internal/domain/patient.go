package domain

type Patient struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	RG           string `json:"rg"`
	RegistryDate string `json:"registry_date"`
}

type CreatePatient struct {
	Name         string `json:"name" binding:"required"`
	Surname      string `json:"surname" binding:"required"`
	RG           string `json:"rg" binding:"required,numeric"`
	RegistryDate string `json:"registry_date" binding:"required" time_format:"2006-01-02 15:04:05"`
}

type UpdatePatient struct {
	Name         string `json:"name" binding:"required"`
	Surname      string `json:"surname" binding:"required"`
	RG           string `json:"rg" binding:"required"`
	RegistryDate string `json:"registry_date" binding:"required" time_format:"2006-01-02 15:04:05"`
}

type PatchPatientName struct {
	Name string `json:"name" binding:"required"`
}
