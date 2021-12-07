package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	// 注：以下插件均可通过前面加 // 注释，注释后停用并不加载插件

	// 词库类
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_atri"      // ATRI词库
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_chat"      // 基础词库
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_qingyunke" // 青云客

	// 实用类
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_b14"          // base16384加解密
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_github"       // 搜索GitHub仓库
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_manager"      // 群管
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_nbnhhsh"      // 拼音首字母缩写释义工具
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_runcode"      // 在线运行代码
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_sleep_manage" // 统计睡眠时间
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_translation"  // 翻译

	// 娱乐类
	// _ "github.com/FloatTech/ZeroBot-Plugin/plugin_wtf"    // 鬼东西
	_ "github.com/FloatTech/ZeroBot-Plugin-Gif"              // 制图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_ai_false"  // 服务器监控
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_choose"    // 选择困难症帮手
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_fortune"   // 运势
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_hs"        // 炉石
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_minecraft" // MCSManager
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_moyu"      // 摸鱼
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_music"     // 点歌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_omikuji"   // 浅草寺求签
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_reborn"    // 投胎
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_shindan"   // 测定

	// b站相关
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_bilibili" // 查询b站用户信息
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_diana"    // 嘉心糖发病

	// 二次元图片
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_acgimage"      // 随机图片与AI点评
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_aiwife"        // 随机老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_image_finder"  // 关键字搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_lolicon"       // lolicon 随机图片
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_nativesetu"    // 本地涩图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_nativewife"    // 本地老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_saucenao"      // 以图搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_setutime"      // 来份涩图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_tracemoe"      // 搜番
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin_vtb_quotation" // vtb语录

	// 以下为内置依赖，勿动
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var (
	contents = []string{
		"* OneBot + go-cqhttp + ZeroBot + Golang",
		"* Version 1.2.1g - 2021-12-3 23:30:58 +0800 CST",
		"* Copyright © 2020 - 2021 FloatTech. All Rights Reserved.",
		"* Project: https://github.com/FloatTech/ZeroBot-Plugin",
	}
	banner = strings.Join(contents, "\n")
	qqs    []string
)

func init() {
	arg := os.Args
	if len(arg) > 1 {
		for _, a := range arg {
			i, err := strconv.ParseUint(a, 10, 64)
			if err == nil {
				qqs = append(qqs, strconv.FormatUint(i, 10))
			}
		}
	}
}

func printBanner() {
	fmt.Print(
		"\n======================[ZeroBot-Plugin]======================",
		"\n", banner, "\n",
		"============================================================\n",
	)
}

func init() {
	driver.NewFuncallClient("zbp", func(f *driver.FCClient) {
		printBanner()
		// 帮助
		zero.OnFullMatchGroup([]string{"/help", ".help", "菜单"}, zero.OnlyToMe).SetBlock(true).FirstPriority().
			Handle(func(ctx *zero.Ctx) {
				ctx.SendChain(message.Text(banner))
			})
		zero.Run(
			zero.Config{
				NickName:      []string{"椛椛", "ATRI", "atri", "亚托莉", "アトリ"},
				CommandPrefix: "/",
				// SuperUsers 某些功能需要主人权限，可通过以下两种方式修改
				// "12345678", "87654321"：通过代码写死的方式添加主人账号
				// flag.Args()：通过命令行参数的方式添加主人账号，无需修改下方任何代码
				SuperUsers: append([]string{"12345678", "87654321"}, qqs...),
				Driver:     []zero.Driver{f},
			},
		)
	})
}
