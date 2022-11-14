package app

import (
	"data-export/app/model"
	"data-export/pkg/config"
	"data-export/pkg/console"
	"data-export/pkg/str"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"time"
)

// TimeNowInTimezone get current time, support setting time zone
func TimeNowInTimezone() time.Time {
	location, err := time.LoadLocation(config.GetString("app.timezone"))
	console.ExitIf(err)
	return time.Now().In(location)
}

func IsDebug() bool {
	return config.GetBool("app.debug")
}

func Name() string {
	return config.GetString("app.name")
}

func JwtUser(ctx *gin.Context) (user model.User) {
	value, _ := ctx.Get("jwtUser")
	user = value.(model.User)
	return
}

func ExportExcel(header []string, data []map[string]string) (string, error) {
	f := excelize.NewFile()
	r := 1 //行
	s := 1 //工作簿
	for _, item := range data {
		if r == 1000000 {
			r = 1
			s++
			f.NewSheet(fmt.Sprintf("Sheet%v", s))
		}
		sheet := fmt.Sprintf("Sheet%v", s)
		c := 1 //列
		for _, h := range header {
			if r == 1 { //标题
				name, _ := excelize.CoordinatesToCellName(c, r)
				_ = f.SetCellValue(sheet, name, h)
				name, _ = excelize.CoordinatesToCellName(c, r+1)
				_ = f.SetCellValue(sheet, name, item[h])
				c++
				continue
			}
			name, _ := excelize.CoordinatesToCellName(c, r+1)
			_ = f.SetCellValue(sheet, name, item[h])
			c++
		}
		r++
	}
	name := fmt.Sprintf("./tmp/%v.xlsx", str.RandString(32))
	err := f.SaveAs(name)
	return name, err
}
