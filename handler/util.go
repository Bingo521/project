package handler

import (
	"errors"
	"my_project/dal/cache"
	"my_project/model"
	"my_project/proto_gen/class_schedule"
)

func CheckSchoolIsLegal(school string)error{
	schools:=cache.GetSchools()
	for _,schoolName := range schools{
		if schoolName == school {
			return nil
		}
	}
	return errors.New("not find school")
}

func GetClass(classInfos [][]*model.ClassInfo)[]*class_schedule.Course{
	courseInfos := []*class_schedule.Course{}
	for i,row := range classInfos{
		for j,cell := range row{
			if cell == nil ||
				cell.ClassName == ""{
				continue
			}
			courseInfo := &class_schedule.Course{
				WeekNum: class_schedule.Week(i),
				ClassNum: int32(j),
				CourseName: cell.ClassName,
				TeacherName: cell.TeacherName,
			}
			courseInfos = append(courseInfos, courseInfo)
		}
	}
	return courseInfos
}
