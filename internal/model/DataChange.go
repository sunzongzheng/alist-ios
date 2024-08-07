package model

import (
	"context"
	"github.com/alist-org/alist/v3/eventbus"
	"gorm.io/gorm"
)

type DataChangeEvent struct {
	Model string
}

func NotifyDataChange(model string) {
	eventbus.Publish[*DataChangeEvent]()(context.Background(), &DataChangeEvent{Model: model})
}

// 存储
func (t *Storage) AfterSave(*gorm.DB) (err error) {
	NotifyDataChange("Storage")
	return nil
}

func (t *Storage) AfterDelete(*gorm.DB) (err error) {
	NotifyDataChange("Storage")
	return nil
}

/*
// 用户
func (t *User) AfterSave(*gorm.DB) (err error) {
	NotifyDataChange("User")
	return nil
}

func (t *User) AfterDelete(*gorm.DB) (err error) {
	NotifyDataChange("User")
	return nil
}

// 设置
func (t *SettingItem) AfterSave(*gorm.DB) (err error) {
	NotifyDataChange("SettingItem")
	return nil
}

func (t *SettingItem) AfterDelete(*gorm.DB) (err error) {
	NotifyDataChange("SettingItem")
	return nil
}

// meta
func (t *Meta) AfterSave(*gorm.DB) (err error) {
	NotifyDataChange("Meta")
	return nil
}

func (t *Meta) AfterDelete(*gorm.DB) (err error) {
	NotifyDataChange("Meta")
	return nil
}
*/
