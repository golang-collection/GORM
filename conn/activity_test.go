package conn

import (
	"GORM/model"
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-11 17:10
* @Description:
**/


func TestSelectActivityById(t *testing.T) {
	activity, err := SelectActivityById(1)
	if err != nil{
		t.FailNow()
	}
	fmt.Println(activity)
}

func TestSelectActivityRandom(t *testing.T) {
	activity, err := SelectActivityRandom()
	if err != nil{
		t.FailNow()
	}
	fmt.Println(activity)
}

func TestInsertActivity(t *testing.T) {
	activity := &model.Activity{
		PlayerId:2,
		DeviceId:2,
		EventDate:"2020-08-10",
		GamesPlayed:5,
		ActivityId:5,
	}
	err := InsertActivity(activity)
	if err != nil {
		t.FailNow()
	}
}

func TestUpdateActivity(t *testing.T) {
	activity := &model.Activity{
		PlayerId:7,
		DeviceId:7,
		EventDate:"2020-08-10",
		GamesPlayed:7,
		ActivityId:4,
	}
	err := UpdateActivity(activity)
	if err != nil {
		t.FailNow()
	}
}

func TestDeleteActivity(t *testing.T) {
	err := DeleteActivity(3)
	if err != nil{
		t.FailNow()
	}
}