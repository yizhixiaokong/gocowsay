# gocowsay

[![asciicast](https://asciinema.org/a/8Ys75reg1PUkGEoc1PQXYtPcr.svg)](https://asciinema.org/a/8Ys75reg1PUkGEoc1PQXYtPcr)

## 简介

gocowsay 是一个基于go实现的命令行工具，可以在终端中显示带有文字气泡的动物字符画。本项目参考[Building a CLI command with Go: cowsay](https://flaviocopes.com/go-tutorial-cowsay/)实现，并进行了以下增强：

- 优化了对中文字符宽度的适应
- 增加了多种小动物的字符画

## 安装

```bash
go install github.com/yizhixiaokong/gocowsay@latest
```

## 使用方法

基本用法：
```bash
echo "为什么要演奏春日影？！！" | gocowsay
```

选择动物：
```bash
echo "BanG Dream! It s MyGO\!\!\!\!\!" | gocowsay -f gopher
```

查看可用动物数量：
```bash
gocowsay -c
```

查看帮助信息：
```bash
gocowsay -h # 可显示可用的动物种类
```

## 致谢

字符画来源：

- [https://github.com/jesslam948/gophersay](https://github.com/jesslam948/gophersay)
- [Building a CLI command with Go: cowsay](https://flaviocopes.com/go-tutorial-cowsay/)

## 更新日志

查看[CHANGELOG.md](./CHANGELOG.md)了解项目的所有变更。