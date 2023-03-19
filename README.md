<div align="center">
  <img src=".github/yaya.jpg" width = "150" height = "150" alt="OneBot-YaYa"><br>
  <h1>ZeroBot-Plugin-Dynamic</h1>
  ZeroBot-Plugin 的模块化，不支持 Windows，需要 CGO<br><br>

  <img src="http://counter.seku.su/cmoe?name=ZeroBot-Plugin&theme=r34" />

[![GOCQ](https://img.shields.io/badge/OneBot-MiraiGo-green.svg?style=social&logo=appveyor)](https://github.com/Mrs4s/go-cqhttp)
[![OICQ](https://img.shields.io/badge/OneBot-OICQ-green.svg?style=social&logo=appveyor)](https://github.com/takayama-lily/node-onebot)
[![MIRAI](https://img.shields.io/badge/OneBot-Mirai-green.svg?style=social&logo=appveyor)](https://github.com/yyuueexxiinngg/onebot-kotlin)

[![Go Report Card](https://goreportcard.com/badge/github.com/FloatTech/ZeroBot-Plugin?style=flat-square&logo=go)](https://goreportcard.com/report/github.com/github.com/FloatTech/ZeroBot-Plugin)
[![Badge](https://img.shields.io/badge/onebot-v11-black?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAHAAAABwCAMAAADxPgR5AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAxQTFRF////29vbr6+vAAAAk1hCcwAAAAR0Uk5T////AEAqqfQAAAKcSURBVHja7NrbctswDATQXfD//zlpO7FlmwAWIOnOtNaTM5JwDMa8E+PNFz7g3waJ24fviyDPgfhz8fHP39cBcBL9KoJbQUxjA2iYqHL3FAnvzhL4GtVNUcoSZe6eSHizBcK5LL7dBr2AUZlev1ARRHCljzRALIEog6H3U6bCIyqIZdAT0eBuJYaGiJaHSjmkYIZd+qSGWAQnIaz2OArVnX6vrItQvbhZJtVGB5qX9wKqCMkb9W7aexfCO/rwQRBzsDIsYx4AOz0nhAtWu7bqkEQBO0Pr+Ftjt5fFCUEbm0Sbgdu8WSgJ5NgH2iu46R/o1UcBXJsFusWF/QUaz3RwJMEgngfaGGdSxJkE/Yg4lOBryBiMwvAhZrVMUUvwqU7F05b5WLaUIN4M4hRocQQRnEedgsn7TZB3UCpRrIJwQfqvGwsg18EnI2uSVNC8t+0QmMXogvbPg/xk+Mnw/6kW/rraUlvqgmFreAA09xW5t0AFlHrQZ3CsgvZm0FbHNKyBmheBKIF2cCA8A600aHPmFtRB1XvMsJAiza7LpPog0UJwccKdzw8rdf8MyN2ePYF896LC5hTzdZqxb6VNXInaupARLDNBWgI8spq4T0Qb5H4vWfPmHo8OyB1ito+AysNNz0oglj1U955sjUN9d41LnrX2D/u7eRwxyOaOpfyevCWbTgDEoilsOnu7zsKhjRCsnD/QzhdkYLBLXjiK4f3UWmcx2M7PO21CKVTH84638NTplt6JIQH0ZwCNuiWAfvuLhdrcOYPVO9eW3A67l7hZtgaY9GZo9AFc6cryjoeFBIWeU+npnk/nLE0OxCHL1eQsc1IciehjpJv5mqCsjeopaH6r15/MrxNnVhu7tmcslay2gO2Z1QfcfX0JMACG41/u0RrI9QAAAABJRU5ErkJggg==)](https://github.com/howmanybots/onebot)
[![Badge](https://img.shields.io/badge/zerobot-v1.6.11-black?style=flat-square&logo=go)](https://github.com/wdvxdr1123/ZeroBot)
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
zerobot [-h] [-n nickname] [-t token] [-u url] [-p prefix] [-d|w] [-c|s config.json] [-l latency] [-r ringlen] [-x max process time] [qq1 qq2 qq3 ...] [&]
```
- **-h**: 显示帮助
- **-n nickname**: 设置默认昵称，默认为`椛椛`
- **-t token**: 设置`AccessToken`，默认为空
- **-u url**: 设置`Url`，默认为`ws://127.0.0.1:6700`
- **-p prefix**: 设置命令前缀，默认为`/`
- **-d|w**: 开启 debug | warning 级别及以上日志输出
- **-c config.json**: 从`config.json`加载`bot`配置
- **-s config.json**: 保存现在`bot`配置到`config.json`
- **-l latency**: 全局处理延时 (ms)
- **-r ringlen**: 接收消息环缓冲区大小，`0`为不设缓冲，并发处理
- **-x max process time**: 最大处理时间 (min)
- **qqs**: superusers 的 qq 号
- **&**: 驻留在后台，必须放在最后，仅`Linux`下有效

默认配置文件格式如下。当选择从配置文件加载时，将忽略相应命令行参数。
```json
{
    "zero": {
        "nickname": [
            "椛椛",
            "ATRI",
            "atri",
            "亚托莉",
            "アトリ"
        ],
        "command_prefix": "/",
        "super_users": [],
        "ring_len": 4096,
        "latency": 233000000,
        "max_process_time": 240000000000
    },
    "ws": [
        {
            "Url": "ws://127.0.0.1:6700",
            "AccessToken": ""
        }
    ],
    "wss": null
}
```

## 功能
> 更多插件详见[ZeroBot-Plugin](https://github.com/FloatTech/ZeroBot-Plugin)

- **动态加载插件**
    - [x] /刷新插件
    - [x] /加载插件 service名
    - [x] /卸载插件 service名
    - 仅 Linux, FreeBSD, macOS 可用
    - 开启后主可执行文件大约增加 2M ，每个插件的`.so`文件约为 2 ~ 10 M ，如非必要不建议使用
    - 动态加载的插件需放置在`plugins/`下，命名为`service名.so`，编译模版详见[Plugin-Template](https://github.com/FloatTech/Plugin-Template)
- **插件控制**
    - [x] /响应 (在发送的群/用户开始工作)
    - [x] /沉默 (在发送的群/用户停止工作)
    - [x] /全局响应 (在所有位置开始工作，无视单独的沉默)
    - [x] /全局沉默 (在所有本应沉默的位置停止工作，显式指定启用的位置不受影响)
    - [x] /启用 xxx (在发送的群/用户启用xxx)
    - [x] /禁用 xxx (在发送的群/用户禁用xxx)
    - [x] /此处启用所有插件
    - [x] /此处禁用所有插件
    - [x] /全局启用 xxx
    - [x] /全局禁用 xxx
    - [x] /还原 xxx (在发送的群/用户还原xxx的开启状态到初始状态)
    - 注：当全局未配置或与默认相同时，状态取决于单独配置，后备为默认配置；当全局与默认不同时，状态取决于全局配置，单独配置失效。
    - [x] /改变默认启用状态 xxx
    - [x] /禁止 service qq1 qq2... (禁止 qqs 使用服务 service)
    - [x] /允许 service qq1 qq2... (重新允许 qqs 使用服务 service)
    - [x] /封禁 qq1 qq2... (禁止 qqs 使用全部服务)
    - [x] /解封 qq1 qq2... (允许 qqs 使用全部服务)
    - [x] /用法 xxx
    - [x] /服务列表
    - [x] /设置服务列表显示行数 xx (默认值为 9, 该设置仅运行时有效, zbp 重启后重置)
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
