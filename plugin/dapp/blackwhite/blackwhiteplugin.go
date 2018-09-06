package blackwhite

import (
	"github.com/spf13/cobra"
	"gitlab.33.cn/chain33/chain33/executor/drivers"
	"gitlab.33.cn/chain33/chain33/plugin/dapp/blackwhite/commands"
	"gitlab.33.cn/chain33/chain33/plugin/dapp/blackwhite/executor"
	gt "gitlab.33.cn/chain33/chain33/plugin/dapp/blackwhite/types"
	"gitlab.33.cn/chain33/chain33/pluginmanager/manager"
	"gitlab.33.cn/chain33/chain33/pluginmanager/plugin"
	"gitlab.33.cn/chain33/chain33/types"
)

var gblackwhitePlugin *blackwhitePlugin

func init() {
	types.AllowUserExec = append(types.AllowUserExec, gt.ExecerBlackwhite)
	gblackwhitePlugin = &blackwhitePlugin{}
	manager.RegisterPlugin(gblackwhitePlugin)
}

type blackwhitePlugin struct {
	plugin.PluginBase
}

func (p *blackwhitePlugin) GetPackageName() string {
	return "gitlab.33.cn.blackwhite"
}

func (p *blackwhitePlugin) GetExecutorName() string {
	return executor.GetName()
}

func (p *blackwhitePlugin) InitExecutor() {
	// TODO: 这里应该将初始化的内容统一放在一个初始化的地方
	executor.InitTypes()
	executor.SetReciptPrefix()

	drivers.Register(executor.GetName(), executor.NewBlackwhite, types.ForkV25BlackWhite)
}

func (p *blackwhitePlugin) AddCustomCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(commands.BlackwhiteCmd())
}