package class_schdule

import "my_project/model"

type BoZhouLoader struct {
	openId string
	schoolName string
	stuId string
	stuPassword string
	code string
	weekInfo [][][]*model.ClassInfo
	firstWeekDate int64
}

func GetBoZhouLoader(openId string,schoolName string,stuId string,stuPassword string,code string)ClassSchduler{
	return NewBoZhouLoader(schoolName,stuId,stuPassword,code)
}
func NewBoZhouLoader(schoolName string,stuId string,stuPassword string,code string)*BoZhouLoader{
	return &BoZhouLoader{
		schoolName: schoolName,
		stuId: stuId,
		stuPassword: stuPassword,
		code: code,
	}
}

func (bz *BoZhouLoader)Load()error{

	return nil
}

func (bz *BoZhouLoader)GetCourse()([][][]*model.ClassInfo,error){
	weekInfo := make([][][]*model.ClassInfo,0)

	return weekInfo,nil
}

func (bz *BoZhouLoader)GetFirsetWeekData()int64{

	return 0
}

func (bz *BoZhouLoader)GetErrorInfo()string{
	return "导入失败"
}