package service

import "lazy/app/model"

func (demo *Demo) Get(key string) (value string) {
	demoModel := model.NewDemo()
	value = demoModel.Get(key)
	return
}
