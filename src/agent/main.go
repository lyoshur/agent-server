package main

import (
	"agent/config"
	"agent/task"
	"github.com/lyoshur/agentutils"
)

func main() {
	// 加载配置文件
	conf := config.Load()

	// 计算前置任务
	tasks := make([]agentutils.Task, 0)
	// 装载前置处理任务
	if conf.CORS.Open {
		crossDomain := task.GetCrossDomain(conf.CORS.Headers)
		tasks = append(tasks, &crossDomain)
	}
	if conf.SignHeader.Open {
		headerSign := task.GetHeaderSign(conf.SignHeader.Sign)
		tasks = append(tasks, &headerSign)
	}

	// 启动服务
	agentutils.StartServer(conf.AgentConfig, tasks)
}
