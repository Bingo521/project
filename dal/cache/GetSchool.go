package cache

import "my_project/dal/db"

func GetSchools()[]string{
	return db.GetSchool()
}
