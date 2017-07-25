package service

import "lazy/app/model"

func (demo *Demo) Set(key, value string) (err error) {
	demoModel := model.NewDemo()
	demoModel.Set(key, value)
	return
}
