package alistlib

import (
	"encoding/json"
	"github.com/alist-org/alist/v3/internal/bootstrap"
	"github.com/alist-org/alist/v3/internal/db"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/internal/op"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

func Backup() string {
	db := db.GetDb()
	//var users []model.User
	var storages []model.Storage
	//var settings []model.SettingItem
	//var metas []model.Meta

	//if err := db.Find(&users).Error; err != nil {
	//	log.Printf("Error fetching users: %v", err)
	//	return ""
	//}
	if err := db.Find(&storages).Error; err != nil {
		log.Printf("Error fetching storages: %v", err)
		return ""
	}
	//if err := db.Find(&settings).Error; err != nil {
	//	log.Printf("Error fetching setting items: %v", err)
	//	return ""
	//}
	//if err := db.Find(&metas).Error; err != nil {
	//	log.Printf("Error fetching metas: %v", err)
	//	return ""
	//}

	backupData := map[string]interface{}{
		//"users":    users,
		"storages": storages,
		//"settings": settings,
		//"metas":    metas,
	}

	jsonData, err := json.Marshal(backupData)
	if err != nil {
		log.Printf("Error marshalling backup data: %v", err)
		return ""
	}

	return string(jsonData)
}

func Restore(jsonData string) {
	db := db.GetDb()
	var data map[string]json.RawMessage

	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		log.Printf("Error unmarshalling JSON data: %v", err)
		return
	}

	// 恢复Users
	if usersData, ok := data["users"]; ok {
		var users []model.User
		if err := json.Unmarshal(usersData, &users); err == nil {
			for _, user := range users {
				if user.Password != "" {
					user.SetPassword(user.Password)
				}
				err := db.Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "id"}},
					UpdateAll: true,
				}).Create(&user).Error
				if err != nil {
					log.Printf("Error restoring user: %v", err)
				}
			}
		} else {
			log.Printf("Error unmarshalling users data: %v", err)
		}
	}
	// 恢复Storages
	if storagesData, ok := data["storages"]; ok {
		var storages []model.Storage
		if err := json.Unmarshal(storagesData, &storages); err == nil {
			for _, storage := range storages {
				err := db.Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "id"}},
					UpdateAll: true,
				}).Create(&storage).Error
				if err != nil {
					log.Printf("Error restoring storages: %v", err)
				}
			}
			// 更新数据表之后 还需要刷新内存中缓存的数据
			if IsRunning("http") {
				op.ClearStorageCache()
				bootstrap.LoadStorages()
			}
		} else {
			log.Printf("Error unmarshalling storages data: %v", err)
		}
	}
	// 恢复Settings
	if settingsData, ok := data["settings"]; ok {
		var settings []model.SettingItem
		if err := json.Unmarshal(settingsData, &settings); err == nil {
			err := op.SaveSettingItems(settings)
			if err != nil {
				log.Printf("Error restoring settings: %v", err)
			}
		} else {
			log.Printf("Error unmarshalling settings data: %v", err)
		}
	}
	// 恢复Metas
	if metasData, ok := data["metas"]; ok {
		var metas []model.Meta
		if err := json.Unmarshal(metasData, &metas); err == nil {
			for _, meta := range metas {
				err := db.Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "path"}},
					UpdateAll: true,
				}).Create(&meta).Error
				if err != nil {
					log.Printf("Error restoring metas: %v", err)
				}
			}
		} else {
			log.Printf("Error unmarshalling metas data: %v", err)
		}
	}
}
