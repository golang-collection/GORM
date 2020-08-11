package model

/**
* @Author: super
* @Date: 2020-08-11 16:19
* @Description:
**/


type Activity struct {
	//gorm.Model
	ActivityId int `gorm:"column:activity_id" gorm:"PRIMARY_KEY" json:"activity_id"`
	PlayerId int `gorm:"column:player_id" json:"player_id"`
	DeviceId int `gorm:"column:device_id" json:"device_id"`
	EventDate string `gorm:"column:event_date" json:"event_date"`
	GamesPlayed int `gorm:"column:games_played" json:"games_played"`
}