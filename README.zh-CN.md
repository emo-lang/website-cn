# Emo 中文官网 (website-cn)

[Emo](https://github.com/emo-lang/emo) 的中文宣传官网 —— 语法简洁、显式定义、符合直觉。
基于 [Airway](https://github.com/daqing/airway)(Gin + templ)构建,
宣传页全部由服务端渲染,不依赖任何客户端框架。英文版 README 见
[README.md](README.md),两份内容一一对应。

## 页面

| 路由 | 视图 | 内容 |
|---|---|---|
| `/` | `app/views/home` | 落地页:Hero、五大编译目标、特性网格、并发 / 接口 / EmoUI / EmoOS 板块 |
| `/features` | `app/views/features` | 逐节讲解语言的设计决定 |
| `/tour` | `app/views/tour` | 八步语言漫游,每步配注释代码 |
| `/quickstart` | `app/views/quickstart` | 构建编译器、第一个程序、原生构建、包管理、示例 |

支撑组件:

- `app/views/layouts` —— 页面骨架(`Base`)、品牌设计系统(`SiteStyles`)、共享的 `Nav`/`Footer`
- `app/views/emocode` —— 一个小型 Emo 语法高亮器,以及各页通用的终端风格代码块
- `app/assets/public` —— 随仓库提交的静态文件(logo),由 `assets.PublicHandler` 在 `/public/` 下提供

品牌色与英文站一致:Yellow `#EFBF6B` · Dark `#202022` ·
Grey `#F9FAFC` · Blue `#7678ED`,通过 `prefers-color-scheme` 支持明暗两色。

## 开发

`.env` 在脚手架阶段创建 —— 设置 `AIRWAY_ENV`(如 `local`)和 `LISTEN`
地址(`host:port`,如 `:1900`)。宣传页不需要数据库;只有用到 models
时才需要配置 `DSN`。

```bash
go generate ./...   # 编辑 .templ 文件后重新生成 *_templ.go
go run .            # 启动 HTTP 服务
go test ./...       # 运行测试
```

本项目踩过的 templ 语法坑(解析报错信息很隐晦,编辑视图时留意):

- 元素文本不能以 `if ` / `for ` / `switch ` 开头 —— templ 会把它当成控制流语句,换种措辞即可。
- 文本中的字面 `{` 和 `${` 必须写成 HTML 实体(`&#123;`、`&#125;`)—— 它们在 templ 里是表达式定界符。
- 文本里的换行会被折叠;预格式化的命令块要用 `<br/>` 换行。

## 静态导出

`airway static:build` 会把四个页面(在 `export.go` 中注册)连同随仓库提交的前端 bundle 一起渲染到 `dist/`。
