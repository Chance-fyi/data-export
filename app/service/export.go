package service

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/pkg/app"
	"data-export/pkg/g"
	ws "data-export/pkg/websocket"
	"fmt"
	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/panjf2000/ants/v2"
	"strconv"
	"time"
)

var ExportQueue = llq.New()

func CreateExport(r api.CreateExportRequest, userId uint) error {
	var (
		us     model.UserSql
		s      model.Sql
		export model.Export
	)

	g.DB().Table(fmt.Sprintf("%v us", us.TableName())).
		Joins(fmt.Sprintf("left join %v s on s.id = us.sql_id", s.TableName())).
		Where("us.id = ?", r.Id).
		Select("us.id user_sql_id", "s.id sql_id", "s.database_id").
		Scan(&export)

	export.UserId = userId
	export.Filename = r.Filename
	export.Sql = r.Sql

	err := g.DB().Create(&export).Error
	if err != nil {
		return err
	}
	ExportQueue.Enqueue(export.Id)
	return nil
}

func ExportQueueConsumer(size int) {
	pool, _ := ants.NewPool(size)

	for {
		_ = pool.Submit(exportDownloadExcel)
		time.Sleep(time.Millisecond)
	}
}

func exportDownloadExcel() {
	id, ok := ExportQueue.Dequeue()
	if !ok {
		return
	}

	var e model.Export
	g.DB().First(&e, id)
	e.Status = 3

	Db, err := getDb(e.DatabaseId)
	if err != nil {
		e.ErrorMsg = err.Error()
		g.DB().Model(&e).Updates(e)
		return
	}

	name, err := app.ChunkFindExportExcel(Db.Raw(e.Sql))
	if err != nil {
		e.ErrorMsg = err.Error()
		g.DB().Model(&e).Updates(e)
		return
	}
	e.Status = 1
	e.Path = name
	g.DB().Model(&e).Updates(e)
	ws.SendUseTag(strconv.Itoa(int(e.UserId)), api.WsResponse{
		Type: "export",
		Data: api.WsExportResponse{
			Name: e.Filename,
		},
	})
}

func ExportList(r api.ExportListRequest, userId uint) (export []api.ExportListItem, count int64) {
	Db := g.DB().Model(&model.Export{})

	if r.Filename != "" {
		Db.Where("filename like ?", "%"+r.Filename+"%")
	}

	if r.Status != "" {
		Db.Where("status = ?", r.Status)
	}

	Db.Where("user_id = ?", userId)
	Db.Order("id DESC")
	Db.Count(&count)
	Db.Offset((r.Current - 1) * r.PageSize).Limit(r.PageSize).Find(&export)
	return
}

func ExportDownload(r api.ExportDownloadRequest) (e model.Export) {
	_ = g.DB().First(&e, r.Id)
	e.Status = 2
	g.DB().Model(&e).Updates(e)
	return
}
