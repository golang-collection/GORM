package tools

import (
	"GORM/model"
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-11 19:28
* @Description:
**/

func TestActivityToJson(t *testing.T) {
	activity := &model.Activity{
		PlayerId:7,
		DeviceId:7,
		EventDate:"2020-08-10",
		GamesPlayed:7,
		ActivityId:4,
	}
	str, err := ActivityToJson(activity)
	if err != nil{
		t.FailNow()
	}
	fmt.Println(str)
}

func TestJsonToActivity(t *testing.T) {
	str := `{"activity_id":4,"player_id":7,"device_id":7,"event_date":"2020-08-10","games_played":7}`
	activity, err := JsonToActivity(str)
	if err != nil{
		t.FailNow()
	}
	fmt.Println(activity)
}