package main

import (
	"github.com/robfig/cron"
	"github.com/va-len-tine/socks5/utils"
	"log"
	"os/exec"
)

var ch = make(chan string, 1)
var cmd *exec.Cmd
var ss string
var err error
var port = "10808"
var SSPath = "ss.txt"
var SSUrl = "https://bulink.me/sub/mruxq/ss"
var Interval = "0 */10 * * * *"

func main()  {
	log.Printf("正在获取代理...")
	ss,err = utils.DfShadowsocks.GetAvailSS(SSUrl, SSPath)
	if err != nil {
		log.Fatal(err)
	}

	cmd = utils.DfShadowsocks.NewSock5Proxy(ss, port)
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer cmd.Process.Kill()
	log.Printf("代理启动成功！端口:10808 PID:%d %s\n\n", cmd.Process.Pid, ss)

	// 开启定时任务自动切换代理
	cron2 := cron.New()
	err = cron2.AddFunc(Interval, test)
	if err != nil {
		log.Fatal("添加定时任务失败！")
	}
	cron2.Start()
	defer cron2.Stop()
	for {
		select {
		case ss = <-ch:
			cmd.Process.Kill()
			cmd = utils.DfShadowsocks.NewSock5Proxy(ss ,port)
			cmd.Start()
			log.Printf("自动切换代理！端口:10808 PID:%d %s\n\n", cmd.Process.Pid, ss)
		}
	}
}

func test()  {
	log.Println("开启自动测速...")
	ss,err = utils.DfShadowsocks.GetFastSS(SSUrl, SSPath)
	if err != nil {
		log.Fatal(err)
	}
	ch <- ss
}
