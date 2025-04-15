package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zzu-andrew/toolkit/config"
	"github.com/zzu-andrew/toolkit/pkg/tea"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"syscall"

	"github.com/zzu-andrew/toolkit/utils"
)

func init() {
	// 支持命令前缀匹配
	cobra.EnablePrefixMatching = true
}

// Service config a struct with stop chan
type Service struct {
	stop chan os.Signal
}

var svc = &Service{
	stop: make(chan os.Signal, 1),
}

func Start() error {
	// 创建日志记录器， 每个 100M, 两个备份，最多三个，备份日志最长保存30天，压缩备份日志
	log, err := utils.CreateProductZapLogger(zapcore.DebugLevel, 100, 2, 30, true)
	if err != nil {
		fmt.Println("Failed to create logger")
		return err
	}

	ctx := config.NewCtx(config.GetConfig(), log)
	// capture signals
	go func() {
		defer func(log *zap.Logger) {
			err := log.Sync()
			if err != nil {

			}
		}(log)

		select {
		case <-ctx.Context.Done():
			return
			// Quit all goroutines
			//ctx.Cancel()
		case <-svc.stop:
			// Quit all goroutines
			ctx.Cancel()
			return
		}
	}()
	// 配置文件路径
	rootCmd := &cobra.Command{
		Use:        "toolkit",
		Short:      "toolkit is a automated k8s tools",
		Long:       `Automated get the k8s infos`,
		SuggestFor: []string{"toolkit --config.file xxx.yaml"},

		Run: func(cmd *cobra.Command, args []string) {
			// 解析配置文件
			err = config.ParseConfig()
			if err != nil {
				log.Error("Failed to parse config", zap.Error(err))
			}

			tea.Tea()
			// Quit all goroutines
			ctx.Cancel()
			log.Info("shutting down...")
		},
	}

	// config --config.file data
	rootCmd.PersistentFlags().String("config.file", "xxx.yaml", "config file path")
	// Bind viper to the root command
	err = viper.BindPFlag("configFile", rootCmd.PersistentFlags().Lookup("config.file"))
	if err != nil {
		log.Error("Error binding flag", zap.Error(err))
		return err
	}
	viper.SetConfigType("yaml")

	log.Info("Starting kube-tools", zap.Any("git versions info", utils.GetVersion()))

	// Make help just show the usage
	rootCmd.SetHelpTemplate(`{{.UsageString}}`)
	// Capture signals
	signal.Notify(svc.stop, syscall.SIGINT, syscall.SIGTERM)
	return rootCmd.Execute()
}
func main() {
	if err := Start(); err != nil {
		fmt.Println("Failed to start proxy")
		return
	}
}
