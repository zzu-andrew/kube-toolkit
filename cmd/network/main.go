package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os/signal"
	"syscall"

	"kube-tools/utils"
)

func Start() error {
	// 创建日志记录器， 每个 100M, 两个备份，最多三个，备份日志最长保存30天，压缩备份日志
	log, err := utils.CreateProductZapLogger(zapcore.DebugLevel, 100, 2, 30, true)
	if err != nil {
		fmt.Println("Failed to create logger")
		return err
	}

	log.Info("Starting proxy!", zap.String("Config info", "/v1/data"),
		zap.Any("git versions info", utils.GetVersion()))

	// capture signals
	go func() {
		defer func(log *zap.Logger) {
			err := log.Sync()
			if err != nil {

			}
		}(log)

		select {
		case <-ctx.Ctx.Done():
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
		Use:        "ysp_installer",
		Short:      "ysp_installer is a automated deployment tools",
		Long:       `Automated deployment tool for deploying ysp stand-alone and cluster`,
		SuggestFor: []string{"ysp_installer --config.file xxx.yaml"},

		Run: func(cmd *cobra.Command, args []string) {
			// Parse configuration file info

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

	rootCmd.SetUsageFunc(usageFunc)
	// Make help just show the usage
	rootCmd.SetHelpTemplate(`{{.UsageString}}`)
	// Capture signals
	signal.Notify(svc.stop, syscall.SIGINT, syscall.SIGTERM)
	return rootCmd.Execute()
}
func main() {
	// TODO: add your code here
}
