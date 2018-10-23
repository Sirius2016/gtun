package god

import (
	"github.com/ICKelin/glog"
	"github.com/ICKelin/gtun/god/controller"
	"github.com/gin-gonic/gin"
)

func Main() {
	opts, err := ParseArgs()
	if err != nil {
		glog.FATAL(err)
	}

	config, err := ParseConfig(opts.confPath)
	if err != nil {
		glog.FATAL(err)
	}

	go func() {
		d := controller.NewGtund(config.GtundConfig)
		glog.FATAL(d.Run())
	}()

	go func() {
		c := controller.NewGtun(config.GtunConfig)
		glog.FATAL(c.Run())
	}()

	engine := gin.Default()
	err = engine.Run(config.Listener)
	if err != nil {
		glog.ERROR(err)
	}
}
