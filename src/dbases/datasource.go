package dbases

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

//MonitorMsg 监控指标信息
type MonitorMsg struct {
	Address     string `json:"address"`
	MonitorItem string `json:"monitoritem"`
}

//QueryResult 创建一个方法
func (m *MonitorMsg) QueryResult(proapi string) string {
	// 创建Prometheus API客户端
	config := api.Config{
		Address: proapi,
	}
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	api := v1.NewAPI(client)

	// 构建查询参数
	query := m.MonitorItem + "{IP=" + "'" + m.Address + "'" + "}"
	// startTime := time.Now().Add(-1 * time.Second)
	startTime, endTime := time.Now(), time.Now()
	step := 1 * time.Second
	queryRange := v1.Range{
		Start: startTime,
		End:   endTime,
		Step:  step,
	}

	// 执行查询
	ctx := context.Background()
	result, warnings, err := api.QueryRange(ctx, query, queryRange)
	if err != nil {
		panic(err)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}

	// 处理查询结果
	return result.String()
}
