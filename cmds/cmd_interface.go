package cmds

import (
	"github.com/EverlongProject/i18n4go/common"
)

type CommandInterface interface {
	common.PrinterInterface
	Options() common.Options
	Run() error
}
