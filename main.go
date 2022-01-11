package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	// 插件控制
	// webctrl "github.com/FloatTech/zbpctrl/web" // web 后端控制

	"github.com/fumiama/go-registry"
	"github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"

	_ "github.com/FloatTech/ZeroBot-Plugin-Dynamic/dyloader"
)

var (
	contents = []string{
		"* OneBot + ZeroBot + Golang",
		"* Version 1.2.3d - 2022-01-07 20:09:30 +0800 CST",
		"* Copyright © 2020 - 2021 FloatTech. All Rights Reserved.",
		"* Project: https://github.com/FloatTech/ZeroBot-Plugin",
	}
	nicks  = []string{"ATRI", "atri", "亚托莉", "アトリ"}
	banner = strings.Join(contents, "\n")
	token  *string
	url    *string
	adana  *string
	prefix *string
	reg    = registry.NewRegReader("reilia.fumiama.top:32664", "fumiama")
)

func init() {
	// 解析命令行参数
	d := flag.Bool("d", false, "Enable debug level log and higher.")
	w := flag.Bool("w", false, "Enable warning level log and higher.")
	h := flag.Bool("h", false, "Display this help.")
	// 解析命令行参数，输入 `-g 监听地址:端口` 指定 gui 访问地址，默认 127.0.0.1:3000
	// g := flag.String("g", "127.0.0.1:3000", "Set web gui listening address.")

	// 直接写死 AccessToken 时，请更改下面第二个参数
	token = flag.String("t", "", "Set AccessToken of WSClient.")
	// 直接写死 URL 时，请更改下面第二个参数
	url = flag.String("u", "ws://127.0.0.1:6700", "Set Url of WSClient.")
	// 默认昵称
	adana = flag.String("n", "椛椛", "Set default nickname.")
	prefix = flag.String("p", "/", "Set command prefix.")

	flag.Parse()
	if *h {
		printBanner()
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	} else {
		if *d && !*w {
			logrus.SetLevel(logrus.DebugLevel)
		}
		if *w {
			logrus.SetLevel(logrus.WarnLevel)
		}
	}
	// 启用 gui
	// webctrl.InitGui(*g)
}

func printBanner() {
	fmt.Print(
		"\n======================[ZeroBot-Plugin]======================",
		"\n", banner, "\n",
		"----------------------[ZeroBot-公告栏]----------------------",
		"\n", getKanban(), "\n",
		"============================================================\n",
	)
}

func getKanban() string {
	err := reg.Connect()
	if err != nil {
		return err.Error()
	}
	defer reg.Close()
	text, err := reg.Get("ZeroBot-Plugin/kanban")
	if err != nil {
		return err.Error()
	}
	return text
}

func main() {
	printBanner()
	// 帮助
	zero.OnFullMatchGroup([]string{"/help", ".help", "菜单"}, zero.OnlyToMe).SetBlock(true).FirstPriority().
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text(banner, "\n可发送\"/服务列表\"查看 bot 功能"))
		})
	zero.OnFullMatch("查看zbp公告", zero.OnlyToMe, zero.AdminPermission).SetBlock(true).FirstPriority().
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text(getKanban()))
		})
	zero.RunAndBlock(
		zero.Config{
			NickName:      append([]string{*adana}, nicks...),
			CommandPrefix: *prefix,
			// SuperUsers 某些功能需要主人权限，可通过以下两种方式修改
			// SuperUsers: []string{"12345678", "87654321"}, // 通过代码写死的方式添加主人账号
			SuperUsers: flag.Args(), // 通过命令行参数的方式添加主人账号
			Driver:     []zero.Driver{driver.NewWebSocketClient(*url, *token)},
		},
	)
}
