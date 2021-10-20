中文 | [English](./README-US.md)

## :memo: 介绍

`reserver` 是一款为静态网站预览或开发设计的具有实时重新加载功能的本地服务器。

> 其主要运用场景为：
> - 单页应用的预览（例如Vue编译之后的项目不可以直接通过`file://`协议进行访问）
> - 具有`ajax`请求的页面（因浏览器安全限制，默认禁止`file://`协议进行`ajax`请求）
> - 静态页面开发（无需手动刷新，实时预览页面加速开发效率）

## :tada: 使用

到 [Releases ](https://github.com/Licoy/reserver/releases) 中选择自己操作系统的最新发行版，使用方法：

- 将此应用放置在您的项目的根目录。
- 将本程序添加至全局，在您的项目中打开终端启动 `reserver [options]` 即可。

## :wrench: 配置

```text
  -p, --port        监听端口 (默认值: 8080)
  -r, --root        根目录
  -H, --host        主机绑定记录 (default: 0.0.0.0)
      --no-browser  不自动打开浏览器
      --no-watch    不观察文件改变
      --browser     指定打开的浏览器
  -P, --path        默认打开的链接路径
      --hide-log    隐藏观察文件改变的日志
  -w, --wait        在重新加载之前等待指定的时间 (默认值: 100ms)
  -i, --ignore      忽略观察文件改变的路径（允许多个）,例如：-i /a -i /b
  -v, --version     查看当前版本
  -h, --help        显示帮助信息
```

## :label: 反馈
- 若您在使用过程中遇到问题，可以直接提交 [issue](https://github.com/Licoy/reserver/issues/new) 进行讨论或反馈。

## :page_facing_up: 许可证

[MIT](./LICENSE)