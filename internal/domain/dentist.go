package domain

type Dentist struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Registry string `json:"registry"`
}

type CreateDentist struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Registry string `json:"registry" binding:"required"`
}

type UpdateDentist struct {
	Name     string `json:"name" binding:"omitempty,required"`
	Surname  string `json:"surname" binding:"omitempty,required"`
	Registry string `json:"registry" binding:"omitempty,required"`
}

type PatchDentistName struct {
	Name string `json:"name" binding:"required"`
}
