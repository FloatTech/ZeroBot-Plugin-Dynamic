<div align="center">
  <img src=".github/yaya.jpg" width = "150" height = "150" alt="OneBot-YaYa"><br>
  <h1>ZeroBot-Plugin-Dynamic</h1>
  ZeroBot-Plugin 的模块化，不支持 Windows，需要 CGO<br><br>

  <img src="http://sayuri.fumiama.top/cmoe?name=ZeroBot-Plugin&theme=r34" />

[![YAYA](https://img.shields.io/badge/OneBot-YaYa-green.svg?style=social&logo=appveyor)](https://github.com/Yiwen-Chan/OneBot-YaYa)
[![GOCQ](https://img.shields.io/badge/OneBot-MiraiGo-green.svg?style=social&logo=appveyor)](https://github.com/Mrs4s/go-cqhttp)
[![OICQ](https://img.shields.io/badge/OneBot-OICQ-green.svg?style=social&logo=appveyor)](https://github.com/takayama-lily/node-onebot)
[![MIRAI](https://img.shields.io/badge/OneBot-Mirai-green.svg?style=social&logo=appveyor)](https://github.com/yyuueexxiinngg/onebot-kotlin)

[![Go Report Card](https://goreportcard.com/badge/github.com/FloatTech/ZeroBot-Plugin?style=flat-square&logo=go)](https://goreportcard.com/report/github.com/github.com/FloatTech/ZeroBot-Plugin)
[![Badge](https://img.shields.io/badge/onebot-v11-black?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAHAAAABwCAMAAADxPgR5AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAxQTFRF////29vbr6+vAAAAk1hCcwAAAAR0Uk5T////AEAqqfQAAAKcSURBVHja7NrbctswDATQXfD//zlpO7FlmwAWIOnOtNaTM5JwDMa8E+PNFz7g3waJ24fviyDPgfhz8fHP39cBcBL9KoJbQUxjA2iYqHL3FAnvzhL4GtVNUcoSZe6eSHizBcK5LL7dBr2AUZlev1ARRHCljzRALIEog6H3U6bCIyqIZdAT0eBuJYaGiJaHSjmkYIZd+qSGWAQnIaz2OArVnX6vrItQvbhZJtVGB5qX9wKqCMkb9W7aexfCO/rwQRBzsDIsYx4AOz0nhAtWu7bqkEQBO0Pr+Ftjt5fFCUEbm0Sbgdu8WSgJ5NgH2iu46R/o1UcBXJsFusWF/QUaz3RwJMEgngfaGGdSxJkE/Yg4lOBryBiMwvAhZrVMUUvwqU7F05b5WLaUIN4M4hRocQQRnEedgsn7TZB3UCpRrIJwQfqvGwsg18EnI2uSVNC8t+0QmMXogvbPg/xk+Mnw/6kW/rraUlvqgmFreAA09xW5t0AFlHrQZ3CsgvZm0FbHNKyBmheBKIF2cCA8A600aHPmFtRB1XvMsJAiza7LpPog0UJwccKdzw8rdf8MyN2ePYF896LC5hTzdZqxb6VNXInaupARLDNBWgI8spq4T0Qb5H4vWfPmHo8OyB1ito+AysNNz0oglj1U955sjUN9d41LnrX2D/u7eRwxyOaOpfyevCWbTgDEoilsOnu7zsKhjRCsnD/QzhdkYLBLXjiK4f3UWmcx2M7PO21CKVTH84638NTplt6JIQH0ZwCNuiWAfvuLhdrcOYPVO9eW3A67l7hZtgaY9GZo9AFc6cryjoeFBIWeU+npnk/nLE0OxCHL1eQsc1IciehjpJv5mqCsjeopaH6r15/MrxNnVhu7tmcslay2gO2Z1QfcfX0JMACG41/u0RrI9QAAAABJRU5ErkJggg==)](https://github.com/howmanybots/onebot)
[![Badge](https://img.shields.io/badge/zerobot-v1.4.1-black?style=flat-square&logo=go)](https://github.com/wdvxdr1123/ZeroBot)
[![License](https://img.shields.io/github/license/FloatTech/ZeroBot-Plugin.svg?style=flat-square&logo=gnu)](https://raw.githubusercontent.com/FloatTech/ZeroBot-Plugin/master/LICENSE)
[![qq group](https://img.shields.io/badge/group-1048452984-red?style=flat-square&logo=tencent-qq)](https://jq.qq.com/?_wv=1027&k=QMb7x1mM)

本项目符合 [OneBot](https://github.com/howmanybots/onebot) 标准，可基于以下项目与机器人框架/平台进行交互
| 项目地址 | 平台 | 核心作者 |
| --- | --- | --- |
| [Mrs4s/go-cqhttp](https://github.com/Mrs4s/go-cqhttp) | [MiraiGo](https://github.com/Mrs4s/MiraiGo) | Mrs4s |
| [yyuueexxiinngg/cqhttp-mirai](https://github.com/yyuueexxiinngg/cqhttp-mirai) | [Mirai](https://github.com/mamoe/mirai) | yyuueexxiinngg |
| [takayama-lily/onebot](https://github.com/takayama-lily/onebot) | [OICQ](https://github.com/takayama-lily/oicq) | takayama |

</div>

## 命令行参数
> `[]`代表是可选参数
```bash
zerobot [-h] [-t token] [-u url] [-n nickname] [-p prefix] [-d|w] [-g 监听地址:端口] [qq1 qq2 qq3 ...] [&]
```
- **-h**: 显示帮助
- **-t token**: 设置`AccessToken`，默认为空
- **-u url**: 设置`Url`，默认为`ws://127.0.0.1:6700`
- **-n nickname**: 设置默认昵称，默认为`椛椛`
- **-p prefix**: 设置命令前缀，默认为`/`
- **-d|w**: 开启 debug | warning 级别及以上日志输出
- **-g 监听地址:端口**: 在 http://监听地址:端口 上开启 [webgui](https://github.com/FloatTech/bot-manager)
- **qqs**: superusers 的 qq 号
- **&**: 驻留在后台，必须放在最后，仅`Linux`下有效

## 功能
> 更多插件详见[ZeroBot-Plugin](https://github.com/FloatTech/ZeroBot-Plugin)

- **web管理** `github.com/FloatTech/zbpctrl/web`
    - 开启后可执行文件大约增加 5M ，默认注释不开启。如需开启请自行编辑`main.go`取消注释
    - 需要配合 [webgui](https://github.com/FloatTech/bot-manager) 使用
- **动态加载插件**
    - [x] /刷新插件
    - [x] /加载插件 service名
    - [x] /卸载插件 service名
    - 仅 Linux, FreeBSD, macOS 可用
    - 开启后主可执行文件大约增加 2M ，每个插件的`.so`文件约为 2 ~ 10 M ，如非必要不建议使用
    - 动态加载的插件需放置在`plugins/`下，命名为`service名.so`，编译模版详见[Plugin-Template](https://github.com/FloatTech/Plugin-Template)
- **插件控制**
    - [x] /启用 xxx (在发送的群/用户启用xxx)
    - [x] /禁用 xxx (在发送的群/用户禁用xxx)
    - [x] /全局启用 xxx
    - [x] /全局禁用 xxx
    - [x] /还原 xxx (在发送的群/用户还原xxx的开启状态到初始状态)
    - [x] /禁止 service qq1 qq2... (禁止 qqs 使用服务 service)
    - [x] /允许 service qq1 qq2... (重新允许 qqs 使用服务 service)
    - [x] /用法 xxx
    - [x] /服务列表
    - [x] /服务详情
    - [x] @Bot 插件冲突检测 (会在本群发送一条消息并在约 1s 后撤回以检测其它同类 bot 中已启用的插件并禁用)

## 使用方法

### 使用稳定版/测试版 (推荐)

可以前往[Release](https://github.com/FloatTech/ZeroBot-Plugin-Dynamic/releases)页面下载对应系统版本可执行文件，并将插件放到`plugins/`目录下。

### 本地直接运行

1. 下载安装最新 [Go](https://studygolang.com/dl) 环境
2. 下载本项目[压缩包](https://github.com/FloatTech/ZeroBot-Plugin/archive/master.zip)，本地解压
3. 编辑 main.go 文件，内容按需修改
4. 运行 OneBot 框架
5. `Windows`下双击 run.bat 文件，`Linux`下使用 run.sh 运行本插件

### 编译运行

#### 利用 Actions 在线编译

1. 点击右上角 Fork 本项目，并转跳到自己 Fork 的仓库
2. 点击仓库上方的 Actions 按钮，确认使用 Actions
3. 编辑 main.go 文件，内容按需修改
4. 前往 Release 页面发布一个 Release，`tag`形如`v1.2.3`，以触发稳定版编译流程
5. 点击 Actions 按钮，等待编译完成，回到 Release 页面下载编译好的文件
6. 运行 OneBot 框架，并同时运行本插件
7. 啾咪~

#### 本地编译/交叉编译

1. 下载安装最新 [Go](https://studygolang.com/dl) 环境
2. clone 并进入本项目，下载所需包

```bash
git clone --depth=1 https://github.com/FloatTech/ZeroBot-Plugin.git
cd ZeroBot-Plugin
go version
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=auto
go mod tidy
```

3. 编辑 main.go 文件，内容按需修改
4. 按照平台输入命令编译，下面举了一些例子

```bash
# 本机平台
go build -ldflags "-s -w" -o zerobot -trimpath
# x64 Linux 平台 如各种云服务器
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o zerobot -trimpath
# x64 Windows 平台 如大多数家用电脑
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o zerobot.exe -trimpath
# armv6 Linux 平台 如树莓派 zero W
GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=0 go build -ldflags "-s -w" -o zerobot -trimpath
# （由于引入了github.com/logoove/sqlite，本项不再可用）mips Linux 平台 如 路由器 wndr4300
GOOS=linux GOARCH=mips GOMIPS=softfloat CGO_ENABLED=0 go build -ldflags "-s -w" -o zerobot -trimpath
```

5. 运行 OneBot 框架，并同时运行本插件

## 特别感谢

- [ZeroBot](https://github.com/wdvxdr1123/ZeroBot)
- [ATRI](https://github.com/Kyomotoi/ATRI)
