package conn

/**
* @Author: super
* @Date: 2020-08-11 16:09
* @Description: gorm自动映射操作activitys表
**/

import (
	"GORM/model"
	"GORM/tools"
	"errors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func SelectActivityById(id int) (*model.Activity, error) {
	db := tools.GetDB()

	activity := &model.Activity{}

	result := db.Where("activity_id = ?", id).Find(activity)
	if result.RecordNotFound(){
		return nil, errors.New("wrong id")
	}

	return  activity, nil
}

func SelectActivityRandom()(*model.Activity, error){
	db := tools.GetDB()

	activity := &model.Activity{}

	result := db.Take(activity)

	if result.RecordNotFound(){
		return nil, errors.New("empty table")
	}

	return  activity, nil
}


func InsertActivity(activity *model.Activity) error {
	db := tools.GetDB()

	result := db.Create(activity)

	if result.RowsAffected == int64(0){
		return errors.New("insert error")
	}

	return  nil
}

func UpdateActivity(activity *model.Activity) error {
	db := tools.GetDB()

	result := db.Model(activity).Where("activity_id = ?", activity.ActivityId).Updates(activity)

	if result.RowsAffected == int64(0){
		return errors.New("update error")
	}

	return  nil
}

func DeleteActivity(id int) error {
	db := tools.GetDB()

	result := db.Where("activity_id = ?", id).Delete(model.Activity{})

	if result.RowsAffected == int64(0){
		return errors.New("delete error")
	}

	return  nil
}