package funcs

import (
	"github.com/open-falcon/common/model"
	"github.com/toolkits/nux"
	"log"
)

// 采集负载数据
func LoadAvgMetrics() []*model.MetricValue {
	load, err := nux.LoadAvg()
	if err != nil {
		log.Println(err)
		return nil
	}

	return []*model.MetricValue{
		GaugeValue("load.1min", load.Avg1min),
		GaugeValue("load.5min", load.Avg5min),
		GaugeValue("load.15min", load.Avg15min),
	}

}
