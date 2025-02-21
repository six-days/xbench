package contracts

import (
	"fmt"

	"github.com/xuperchain/xuper-sdk-go/v2/account"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
)

type Contract interface {
	Deploy(from *account.Account, name string, code []byte, args map[string]string, opts ...xuper.RequestOption) (*xuper.Transaction, error)
	Invoke(from *account.Account, name, method string, args map[string]string, opts ...xuper.RequestOption) (*xuper.Transaction, error)
	Query(from *account.Account, name, method string, args map[string]string, opts ...xuper.RequestOption) (*xuper.Transaction, error)
}

type ContractConfig struct {
	// 合约地址
	ContractAccount string
	// 合约路径
	CodePath string

	// 模块名：wasm/native/evm
	ModuleName string
	// 合约名
	ContractName string
	// invoke方法名
	MethodInvokeName string
	// query方法名
	MethodQueryName string
	// 方法类型
	MethodType string

	// 自定义参数
	Args map[string]string
}

type NewContract func(config *ContractConfig, client *xuper.XClient) (Contract, error)

// 注册合约
var contracts = make(map[string]NewContract, 8)

func RegisterContract(name string, contract NewContract) {
	contracts[name] = contract
}

func GetContract(config *ContractConfig, client *xuper.XClient) (Contract, error) {
	if newContract, ok := contracts[config.ContractName]; ok {
		return newContract(config, client)
	}

	return nil, fmt.Errorf("contract not exist")
}
