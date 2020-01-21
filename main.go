package  main

import (
	"github.com/NJUPT-ISL/SCV/pkg/collection"
	"github.com/NJUPT-ISL/SCV/pkg/cronjob"
	"github.com/NJUPT-ISL/SCV/pkg/ops"
	"github.com/robfig/cron"
	"os"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	collection.Node = os.Getenv("NODE_NAME")
	collection.Mode = os.Getenv("MODE")
	ops.InitSCV(collection.Mode)
	ops.PrintSCV()
	ops.InitInClusterConfig()
	c := cron.New()
	cronjob.UpdateSCVJob(c, collection.Mode, collection.Node)
	cronjob.UpdateModeJob(c)
	defer ops.CleanScvLabel(collection.Node)
	defer c.Stop()
	cronjob.StartJob(c,&wg)
	wg.Wait()
}
