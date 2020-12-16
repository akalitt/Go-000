package model

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	//"github.com/sirupsen/logrus"
	//"github.com/bbkdevadmin/bbm/plugin/logger"
)

var DB *gorm.DB

func ResetDatabase() {
	DB.AutoMigrate(
		&Tag{},
	)
}

// Database is a middleware function that initializes the datastore and attaches to
// the context of every request context.
func Database() {
	var err error

	DB, err = gorm.Open(
		viper.GetString("db.driver"),
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
			viper.GetString("db.host"),
			viper.GetString("db.port"),
			viper.GetString("db.user"),
			viper.GetString("db.name"),
			viper.GetString("db.ssl"),
			viper.GetString("db.password")),
	)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	DB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	DB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	DB.Callback().Delete().Replace("gorm:delete", deleteCallback)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	DB.CreateTable(&Tag{})

}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		fmt.Println("nowTime", nowTime)
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			createTimeField.Set(nowTime)
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			modifyTimeField.Set(nowTime)
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", time.Now())
	}
}

// deleteCallback will set `DeletedOn` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
