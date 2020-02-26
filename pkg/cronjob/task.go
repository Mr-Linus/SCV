package cronjob

import (
	"github.com/NJUPT-ISL/SCV/pkg/collection"
	"github.com/NJUPT-ISL/SCV/pkg/log"
	"github.com/NJUPT-ISL/SCV/pkg/ops"
	"github.com/robfig/cron"
	"os"
	"sync"
)

func UpdateSCVJob(c *cron.Cron, name string) {
	if err := c.AddFunc("*/5 * * * * ?", func() {
		ops.UpdateScvLabel(name)
	}); err != nil {
		log.ErrPrint(err)
	}
}

func UpdateModeJob(c *cron.Cron) {
	if err := c.AddFunc("0 */5 * * * ?", func() {
		collection.Mode = os.Getenv("MODE")
	}); err != nil {
		log.ErrPrint(err)
	}
}
func StartJob(c *cron.Cron, w *sync.WaitGroup) {
	w.Add(1)
	c.Start()
}
