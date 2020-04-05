package db

import (
	"fmt"
	"my_project/logs"
	"my_project/model"
)

var defaultSchools []string

func GetSchool() []string {
	schools := make([]model.SchoolsModel, 0)
	db := db.Select("school_name").Find(&schools)
	if db.Error != nil {
		logs.Error(fmt.Sprintf("load schools err:%v", db.Error))
		return defaultSchools
	}
	schoolNames := make([]string, 0, len(schools))
	for _, school := range schools {
		schoolNames = append(schoolNames, school.SchoolName)
	}
	defaultSchools = schoolNames
	return schoolNames
}
