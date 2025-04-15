package utils

import (
	"fmt"
	"runtime"
)

// 全局变量定义
//
// @Description 这些变量用于存储版本相关的信息，如版本号、Git 分支、标签、提交信息等
var (
	// version 表示当前程序的版本号
	version string
	// gitBranch 表示当前程序所基于的 Git 分支
	gitBranch string
	// gitTag 表示当前程序所基于的 Git 标签
	gitTag string
	// gitCommit 表示当前程序所基于的 Git 提交哈希
	gitCommit string
	// gitTreeState 表示当前 Git 仓库的状态
	gitTreeState string
	// buildDate 表示当前程序的构建日期
	buildDate string
)

// VersionInfo 结构体用于存储版本信息
//
// @Description 包含了程序的版本号、Git 相关信息、构建日期、Go 版本等
// @Tags VersionInfo
type VersionInfo struct {
	// Version 表示当前程序的版本号
	Version string `json:"version"`
	// GitBranch 表示当前程序所基于的 Git 分支
	GitBranch string `json:"gitBranch"`
	// GitTag 表示当前程序所基于的 Git 标签
	GitTag string `json:"gitTag"`
	// GitCommit 表示当前程序所基于的 Git 提交哈希
	GitCommit string `json:"gitCommit"`
	// GitTreeState 表示当前 Git 仓库的状态
	GitTreeState string `json:"gitTreeState"`
	// BuildDate 表示当前程序的构建日期
	BuildDate string `json:"buildDate"`
	// GoVersion 表示编译当前程序所使用的 Go 版本
	GoVersion string `json:"goVersion"`
	// Compiler 表示编译当前程序所使用的编译器
	Compiler string `json:"compiler"`
	// Platform 表示当前程序所运行的平台
	Platform string `json:"platform"`
}

// String 方法用于将 VersionInfo 结构体转换为字符串
//
// @Summary 将 VersionInfo 结构体转换为字符串
// @Description 返回 VersionInfo 结构体中的 Git 提交哈希作为字符串表示
// @Tags VersionInfo
// @Success 200 {string} string "成功返回 Git 提交哈希字符串"
// @Router /versionInfoToString [get]
func (info VersionInfo) String() string {
	return info.GitCommit
}

// GetVersion 函数用于获取当前程序的版本信息
//
// @Summary 获取当前程序的版本信息
// @Description 返回一个包含版本号、Git 信息、构建日期、Go 版本等的 VersionInfo 结构体
// @Tags VersionInfo
// @Router /getVersion [get]
func GetVersion() VersionInfo {
	return VersionInfo{
		Version:      version,
		GitBranch:    gitBranch,
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
