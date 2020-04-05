package model

type SchoolsModel struct {
	SchoolName string `json:"school_name",grom:"school_name"`
}

func (s *SchoolsModel)TableName()string{
	return "school"
}