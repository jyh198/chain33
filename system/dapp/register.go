package dapp

//store package store the world - state data
import (
	log "github.com/inconshreveable/log15"
	"gitlab.33.cn/chain33/chain33/common/address"
	clog "gitlab.33.cn/chain33/chain33/common/log"
	"gitlab.33.cn/chain33/chain33/types"
)

var elog = log.New("module", "execs")

func SetLogLevel(level string) {
	clog.SetLogLevel(level)
}

func DisableLog() {
	elog.SetHandler(log.DiscardHandler())
}

type DriverCreate func() Driver

type driverWithHeight struct {
	create DriverCreate
	height int64
}

var (
	execDrivers        = make(map[string]*driverWithHeight)
	execAddressNameMap = make(map[string]string)
	registedExecDriver = make(map[string]*driverWithHeight)
)

func Register(name string, create DriverCreate, height int64) {
	if create == nil {
		panic("Execute: Register driver is nil")
	}
	if _, dup := registedExecDriver[name]; dup {
		panic("Execute: Register called twice for driver " + name)
	}
	driverWithHeight := &driverWithHeight{
		create: create,
		height: height,
	}
	registedExecDriver[name] = driverWithHeight
	registerAddress(name)
	execDrivers[ExecAddress(name)] = driverWithHeight
}

func LoadDriver(name string, height int64) (driver Driver, err error) {
	// user.evm.xxxx 的交易，使用evm执行器
	//   user.p.evm
	name = string(types.GetRealExecName([]byte(name)))
	c, ok := registedExecDriver[name]
	if !ok {
		return nil, types.ErrUnknowDriver
	}
	if height >= c.height || height == -1 {
		return c.create(), nil
	}
	return nil, types.ErrUnknowDriver
}

func LoadDriverAllow(tx *types.Transaction, index int, height int64) (driver Driver) {
	exec, err := LoadDriver(string(tx.Execer), height)
	if err == nil {
		err = exec.Allow(tx, index)
	}
	if err != nil {
		exec, err = LoadDriver("none", height)
		if err != nil {
			panic(err)
		}
	} else {
		exec.SetName(string(types.GetRealExecName(tx.Execer)))
		exec.SetCurrentExecName(string(tx.Execer))
	}
	return exec
}

func IsDriverAddress(addr string, height int64) bool {
	c, ok := execDrivers[addr]
	if !ok {
		return false
	}
	if height >= c.height || height == -1 {
		return true
	}
	return false
}

func registerAddress(name string) {
	if len(name) == 0 {
		panic("empty name string")
	}
	addr := ExecAddress(name)
	execAddressNameMap[name] = addr
}

func ExecAddress(name string) string {
	if addr, ok := execAddressNameMap[name]; ok {
		return addr
	}
	return address.ExecAddress(name)
}