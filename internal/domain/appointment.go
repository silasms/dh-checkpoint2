package domain

type Appointment struct {
	Id              int     `json:"id"`
	Patient         Patient `json:"patient"`
	Dentist         Dentist `json:"dentist"`
	AppointmentDate string  `json:"appointment_date"`
}

type CreateAppointment struct {
	PatientId       int    `json:"patient_id" binding:"omitempty,required"`
	DentistId       int    `json:"dentist_id" binding:"omitempty,required"`
	AppointmentDate string `json:"appointment_date" binding:"omitempty,required"`
}

type UpdateAppointment struct {
	PatientId       int    `json:"patient_id" binding:"omitempty,required"`
	DentistId       int    `json:"dentist_id" binding:"omitempty,required"`
	AppointmentDate string `json:"appointment_date" binding:"omitempty,required"`
}

type PatchAppointment struct {
	AppointmentDate string `json:"appointment_date" binding:"omitempty,required"`
}