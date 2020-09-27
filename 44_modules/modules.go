package main

// 必须在$GOPATH/src目录以外使用go mod
// 1. go mod init dic_name (只需要在项目的根目录执行go mod init，不需要在子目录执行）。
//		当modules 功能启用时，依赖包的存放位置变更为$GOPATH/pkg，允许同一个package多个版本并存，且多个项目可以共享缓存的 module。
//2. go list -m -u all  用来检查可以升级的package
//   go get -u need-upgrade-package 升级后会将新的依赖版本更新到go.mod
//    或者 go get -u 升级所有依赖