package app

import (
	"data-export/app/model"
	"data-export/pkg/config"
	"data-export/pkg/console"
	"data-export/pkg/str"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
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

func ChunkFindExportExcel(Db *gorm.DB) (string, error) {
	rows, err := Db.Rows()
	if err != nil {
		return "", err
	}
	var (
		columns []string
		list    []map[string]string
		size    = 20000
		f       = excelize.NewFile()
		r       = 1 //行
		s       = 1 //工作簿
		sw, _   = f.NewStreamWriter(fmt.Sprintf("Sheet%v", s))
	)

	columns, err = rows.Columns()
	if err != nil {
		return "", err
	}
	columnLength := len(columns)

	cache := make([]interface{}, columnLength)
	for index := range cache {
		var a interface{}
		cache[index] = &a
	}

	for rows.Next() {
		err = rows.Scan(cache...)
		if err != nil {
			return "", err
		}

		item := make(map[string]string, columnLength)
		for i, data := range cache {
			v := *data.(*interface{})
			switch v.(type) {
			case time.Time:
				item[columns[i]] = v.(time.Time).Format("2006-01-02 15:04:05")
			case []uint8:
				item[columns[i]] = string(v.([]byte))
			}
		}
		list = append(list, item)

		if len(list) == size {
			sw, r, s = exportExcel(columns, list, f, sw, r, s)
			list = nil
		}
	}
	if len(list) > 0 {
		sw, r, s = exportExcel(columns, list, f, sw, r, s)
		list = nil
	}

	_ = sw.Flush()
	name := fmt.Sprintf("./tmp/%v.xlsx", str.RandString(32))
	err = f.SaveAs(name)
	_ = f.Close()
	return name, err
}

func exportExcel(header []string, data []map[string]string, f *excelize.File, sw *excelize.StreamWriter, r int, s int) (*excelize.StreamWriter, int, int) {
	h := funk.Map(header, func(i interface{}) interface{} {
		return i
	}).([]interface{})

	for _, item := range data {
		if r > 1000000 {
			_ = sw.Flush()
			r = 1
			s++
			sheet := fmt.Sprintf("Sheet%v", s)
			f.NewSheet(sheet)
			sw, _ = f.NewStreamWriter(sheet)
		}
		v := funk.Map(header, func(i string) interface{} {
			return item[i]
		}).([]interface{})
		if r == 1 {
			cell, _ := excelize.CoordinatesToCellName(1, r)
			_ = sw.SetRow(cell, h)
			r++
			cell, _ = excelize.CoordinatesToCellName(1, r)
			_ = sw.SetRow(cell, v)
			r++
			continue
		}
		cell, _ := excelize.CoordinatesToCellName(1, r)
		_ = sw.SetRow(cell, v)
		r++
	}
	return sw, r, s
}
