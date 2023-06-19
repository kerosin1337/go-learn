package validate

import (
	"example/web-service-gin/src/database"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var UniqueValidate validator.Func = func(fl validator.FieldLevel) bool {
	var count int64
	split := strings.Split(fl.Param(), ";")
	table := split[0]
	field := split[1]
	database.DB.Table(table).Where(fmt.Sprintf("%s = ?", field), fl.Field()).Count(&count)
	return count == 0
}
