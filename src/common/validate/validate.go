package validate

import (
	"example/web-service-gin/src/database"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
	"time"
)

var UniqueValidate validator.Func = func(fl validator.FieldLevel) bool {
	var count int64
	split := strings.Split(fl.Param(), ";")
	table := split[0]
	field := split[1]
	database.DB.Table(table).Where(fmt.Sprintf("%s = ?", field), fl.Field()).Count(&count)
	return count == 0
}

func DateValidate(fl validator.FieldLevel) bool {
	fmt.Print(123)
	layout := "2006-01-02"
	_, err := time.Parse(layout, fl.Field().String())
	return err == nil
}
