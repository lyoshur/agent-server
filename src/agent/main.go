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
	conf.Tasks = make([]agentutils.Task, 0)
	// 装载前置处理任务
	if conf.CORS.Open {
		crossDomain := task.GetCrossDomain(conf.CORS.Headers)
		conf.Tasks = append(conf.Tasks, &crossDomain)
	}
	if conf.SignHeader.Open {
		headerSign := task.GetHeaderSign(conf.SignHeader.Sign)
		conf.Tasks = append(conf.Tasks, &headerSign)
	}

	// 计算前置任务-代理规则中
	conf.AgentConfig.Proxies = make([]agentutils.AgentProxy, 0)
	for i := range conf.Proxies {
		conf.Proxies[i].Tasks = make([]agentutils.Task, 0)
		if conf.Proxies[i].CORS.Open {
			crossDomain := task.GetCrossDomain(conf.CORS.Headers)
			conf.Proxies[i].Tasks = append(conf.Proxies[i].Tasks, &crossDomain)
		}
		if conf.Proxies[i].SignHeader.Open {
			headerSign := task.GetHeaderSign(conf.SignHeader.Sign)
			conf.Proxies[i].Tasks = append(conf.Proxies[i].Tasks, &headerSign)
		}
		conf.AgentConfig.Proxies = append(conf.AgentConfig.Proxies, conf.Proxies[i].AgentProxy)
	}

	// 启动服务
	agentutils.StartServer(conf.AgentConfig)
}
