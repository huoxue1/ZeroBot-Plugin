package chouxianghua

import (
	"os"

	log "github.com/sirupsen/logrus"

	sql "github.com/FloatTech/sqlite"
	"github.com/FloatTech/zbputils/file"
	"github.com/FloatTech/zbputils/process"

	"github.com/FloatTech/zbputils/control/order"
)

const (
	dbpath = "data/ChouXiangHua/"
	dbfile = dbpath + "cxh.db"
)

var db = &sql.Sqlite{DBPath: dbfile}

// 加载数据库
func init() {
	go func() {
		defer order.DoneOnExit()()
		process.SleepAbout1sTo2s()
		// os.RemoveAll(dbpath)
		_ = os.MkdirAll(dbpath, 0755)
		_, _ = file.GetLazyData(dbfile, false, true)
		err := db.Create("pinyin", &pinyin{})
		if err != nil {
			panic(err)
		}
		n, err := db.Count("pinyin")
		if err != nil {
			panic(err)
		}
		log.Printf("[chouxianghua]读取%d条拼音", n)
	}()
}
