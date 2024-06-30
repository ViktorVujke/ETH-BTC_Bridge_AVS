// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIncredibleSquaringServiceManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractIncredibleSquaringServiceManagerMetaData contains all meta data concerning the ContractIncredibleSquaringServiceManager contract.
var ContractIncredibleSquaringServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_incredibleSquaringTaskManager\",\"type\":\"address\",\"internalType\":\"contractIIncredibleSquaringTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"incredibleSquaringTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIncredibleSquaringTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162001b5738038062001b5783398101604081905262000035916200015a565b6001600160a01b03808516608052600060a081905281851660c05290831660e05284908484620000646200007f565b505050506001600160a01b03166101005250620001c2915050565b600054610100900460ff1615620000ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200013f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015757600080fd5b50565b600080600080608085870312156200017157600080fd5b84516200017e8162000141565b6020860151909450620001918162000141565b6040860151909350620001a48162000141565b6060860151909250620001b78162000141565b939692955090935050565b60805160a05160c05160e051610100516118e062000277600039600081816101460152610930015260008181610650015281816107ab0152818161084201528181610c6d01528181610df00152610e8f01526000818161045b0152818161050a0152818161058a015281816109fd01528181610a9301528181610bab0152610d4b01526000818161030a01526103e801526000818161010c01528181610a5101528181610ae70152610b7001526118e06000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638da5cb5b116100715780638da5cb5b146101685780639926ee7d14610179578063a364f4da1461018c578063a98fb3551461019f578063e481af9d146101b2578063f2fde38b146101ba57600080fd5b80631b445516146100b957806333cfb7b7146100ce57806338c8ee64146100f75780636b3aa72e1461010a578063715018a61461013957806377ef731d14610141575b600080fd5b6100cc6100c7366004611183565b6101cd565b005b6100e16100dc36600461120c565b610455565b6040516100ee9190611230565b60405180910390f35b6100cc61010536600461120c565b610925565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516100ee919061127d565b6100cc6109de565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b031661012c565b6100cc610187366004611344565b6109f2565b6100cc61019a36600461120c565b610a88565b6100cc6101ad3660046113ee565b610b51565b6100e1610ba5565b6100cc6101c836600461120c565b610f6e565b6101d5610fe4565b60005b818110156103d0578282828181106101f2576101f261143e565b90506020028101906102049190611454565b61021590604081019060200161120c565b6001600160a01b03166323b872dd33308686868181106102375761023761143e565b90506020028101906102499190611454565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af11580156102a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c49190611484565b508282828181106102d7576102d761143e565b90506020028101906102e99190611454565b6102fa90604081019060200161120c565b6001600160a01b031663095ea7b37f000000000000000000000000000000000000000000000000000000000000000085858581811061033b5761033b61143e565b905060200281019061034d9190611454565b604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152013560248201526044016020604051808303816000875af115801561039b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103bf9190611484565b506103c9816114bc565b90506101d8565b50604051630da22a8b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631b4455169061041f9085908590600401611566565b600060405180830381600087803b15801561043957600080fd5b505af115801561044d573d6000803e3d6000fd5b505050505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166313542a4e846040518263ffffffff1660e01b81526004016104a5919061127d565b602060405180830381865afa1580156104c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e69190611673565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa158015610551573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610575919061168c565b90506001600160c01b038116158061060f57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060a91906116b5565b60ff16155b1561062b57505060408051600081526020810190915292915050565b600061063f826001600160c01b031661103e565b90506000805b8251811015610715577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f584838151811061068f5761068f61143e565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156106d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f79190611673565b61070190836116d8565b91508061070d816114bc565b915050610645565b506000816001600160401b0381111561073057610730611291565b604051908082528060200260200182016040528015610759578160200160208202803683370190505b5090506000805b845181101561091857600085828151811061077d5761077d61143e565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa1580156107f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108169190611673565b905060005b81811015610902576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610890573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b491906116f0565b600001518686815181106108ca576108ca61143e565b6001600160a01b0390921660209283029190910190910152846108ec816114bc565b95505080806108fa906114bc565b91505061081b565b5050508080610910906114bc565b915050610760565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109db5760405162461bcd60e51b815260206004820152604a60248201527f6f6e6c79496e6372656469626c655371756172696e675461736b4d616e61676560448201527f723a206e6f742066726f6d206372656469626c65207371756172696e6720746160648201526939b59036b0b730b3b2b960b11b608482015260a4015b60405180910390fd5b50565b6109e6610fe4565b6109f06000611100565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a3a5760405162461bcd60e51b81526004016109d29061174e565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061041f9085908590600401611813565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ad05760405162461bcd60e51b81526004016109d29061174e565b6040516351b27a6d60e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90610b1c90849060040161127d565b600060405180830381600087803b158015610b3657600080fd5b505af1158015610b4a573d6000803e3d6000fd5b5050505050565b610b59610fe4565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610b1c90849060040161185e565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c2b91906116b5565b60ff16905080610c4957505060408051600081526020810190915290565b6000805b82811015610cfe57604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610cbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce09190611673565b610cea90836116d8565b915080610cf6816114bc565b915050610c4d565b506000816001600160401b03811115610d1957610d19611291565b604051908082528060200260200182016040528015610d42578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610da7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dcb91906116b5565b60ff16811015610f6457604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610e3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e639190611673565b905060005b81811015610f4f576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610edd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0191906116f0565b60000151858581518110610f1757610f1761143e565b6001600160a01b039092166020928302919091019091015283610f39816114bc565b9450508080610f47906114bc565b915050610e68565b50508080610f5c906114bc565b915050610d49565b5090949350505050565b610f76610fe4565b6001600160a01b038116610fdb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109d2565b6109db81611100565b6033546001600160a01b031633146109f05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109d2565b606060008061104c84611152565b61ffff166001600160401b0381111561106757611067611291565b6040519080825280601f01601f191660200182016040528015611091576020820181803683370190505b5090506000805b8251821080156110a9575061010081105b15610f64576001811b9350858416156110f0578060f81b8383815181106110d2576110d261143e565b60200101906001600160f81b031916908160001a9053508160010191505b6110f9816114bc565b9050611098565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b821561117d57611167600184611871565b909216918061117581611888565b915050611156565b92915050565b6000806020838503121561119657600080fd5b82356001600160401b03808211156111ad57600080fd5b818501915085601f8301126111c157600080fd5b8135818111156111d057600080fd5b8660208260051b85010111156111e557600080fd5b60209290920196919550909350505050565b6001600160a01b03811681146109db57600080fd5b60006020828403121561121e57600080fd5b8135611229816111f7565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156112715783516001600160a01b03168352928401929184019160010161124c565b50909695505050505050565b6001600160a01b0391909116815260200190565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156112c9576112c9611291565b60405290565b60006001600160401b03808411156112e9576112e9611291565b604051601f8501601f19908116603f0116810190828211818310171561131157611311611291565b8160405280935085815286868601111561132a57600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561135757600080fd5b8235611362816111f7565b915060208301356001600160401b038082111561137e57600080fd5b908401906060828703121561139257600080fd5b61139a6112a7565b8235828111156113a957600080fd5b83019150601f820187136113bc57600080fd5b6113cb878335602085016112cf565b815260208301356020820152604083013560408201528093505050509250929050565b60006020828403121561140057600080fd5b81356001600160401b0381111561141657600080fd5b8201601f8101841361142757600080fd5b611436848235602084016112cf565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235609e1983360301811261146a57600080fd5b9190910192915050565b803561147f816111f7565b919050565b60006020828403121561149657600080fd5b8151801515811461122957600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156114d0576114d06114a6565b5060010190565b6001600160601b03811681146109db57600080fd5b8183526000602080850194508260005b8581101561154757813561150f816111f7565b6001600160a01b0316875281830135611527816114d7565b6001600160601b03168784015260409687019691909101906001016114fc565b509495945050505050565b803563ffffffff8116811461147f57600080fd5b60208082528181018390526000906040808401600586901b8501820187855b8881101561166557878303603f190184528135368b9003609e190181126115ab57600080fd5b8a0160a0813536839003601e190181126115c457600080fd5b820180356001600160401b038111156115dc57600080fd5b8060061b36038413156115ee57600080fd5b828752611600838801828c85016114ec565b9250505061160f888301611474565b6001600160a01b03168886015281870135878601526060611631818401611552565b63ffffffff16908601526080611648838201611552565b63ffffffff16950194909452509285019290850190600101611585565b509098975050505050505050565b60006020828403121561168557600080fd5b5051919050565b60006020828403121561169e57600080fd5b81516001600160c01b038116811461122957600080fd5b6000602082840312156116c757600080fd5b815160ff8116811461122957600080fd5b600082198211156116eb576116eb6114a6565b500190565b60006040828403121561170257600080fd5b604051604081018181106001600160401b038211171561172457611724611291565b6040528251611732816111f7565b81526020830151611742816114d7565b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b818110156117ec576020818501810151868301820152016117d0565b818111156117fe576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b038316815260406020820152600082516060604084015261183d60a08401826117c6565b90506020840151606084015260408401516080840152809150509392505050565b60208152600061122960208301846117c6565b600082821015611883576118836114a6565b500390565b600061ffff808316818114156118a0576118a06114a6565b600101939250505056fea2646970667358221220a1ee19651d4f0d5cfb71125a1ee64ccaf6828121ddf6da1f9ebec467d64e868f64736f6c634300080c0033",
}

// ContractIncredibleSquaringServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIncredibleSquaringServiceManagerMetaData.ABI instead.
var ContractIncredibleSquaringServiceManagerABI = ContractIncredibleSquaringServiceManagerMetaData.ABI

// ContractIncredibleSquaringServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIncredibleSquaringServiceManagerMetaData.Bin instead.
var ContractIncredibleSquaringServiceManagerBin = ContractIncredibleSquaringServiceManagerMetaData.Bin

// DeployContractIncredibleSquaringServiceManager deploys a new Ethereum contract, binding an instance of ContractIncredibleSquaringServiceManager to it.
func DeployContractIncredibleSquaringServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _incredibleSquaringTaskManager common.Address) (common.Address, *types.Transaction, *ContractIncredibleSquaringServiceManager, error) {
	parsed, err := ContractIncredibleSquaringServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIncredibleSquaringServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _incredibleSquaringTaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIncredibleSquaringServiceManager{ContractIncredibleSquaringServiceManagerCaller: ContractIncredibleSquaringServiceManagerCaller{contract: contract}, ContractIncredibleSquaringServiceManagerTransactor: ContractIncredibleSquaringServiceManagerTransactor{contract: contract}, ContractIncredibleSquaringServiceManagerFilterer: ContractIncredibleSquaringServiceManagerFilterer{contract: contract}}, nil
}

// ContractIncredibleSquaringServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractIncredibleSquaringServiceManager struct {
	ContractIncredibleSquaringServiceManagerCaller     // Read-only binding to the contract
	ContractIncredibleSquaringServiceManagerTransactor // Write-only binding to the contract
	ContractIncredibleSquaringServiceManagerFilterer   // Log filterer for contract events
}

// ContractIncredibleSquaringServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIncredibleSquaringServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIncredibleSquaringServiceManagerSession struct {
	Contract     *ContractIncredibleSquaringServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                             // Call options to use throughout this session
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ContractIncredibleSquaringServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIncredibleSquaringServiceManagerCallerSession struct {
	Contract *ContractIncredibleSquaringServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                   // Call options to use throughout this session
}

// ContractIncredibleSquaringServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIncredibleSquaringServiceManagerTransactorSession struct {
	Contract     *ContractIncredibleSquaringServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                   // Transaction auth options to use throughout this session
}

// ContractIncredibleSquaringServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIncredibleSquaringServiceManagerRaw struct {
	Contract *ContractIncredibleSquaringServiceManager // Generic contract binding to access the raw methods on
}

// ContractIncredibleSquaringServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringServiceManagerCallerRaw struct {
	Contract *ContractIncredibleSquaringServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIncredibleSquaringServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringServiceManagerTransactorRaw struct {
	Contract *ContractIncredibleSquaringServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIncredibleSquaringServiceManager creates a new instance of ContractIncredibleSquaringServiceManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringServiceManager(address common.Address, backend bind.ContractBackend) (*ContractIncredibleSquaringServiceManager, error) {
	contract, err := bindContractIncredibleSquaringServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManager{ContractIncredibleSquaringServiceManagerCaller: ContractIncredibleSquaringServiceManagerCaller{contract: contract}, ContractIncredibleSquaringServiceManagerTransactor: ContractIncredibleSquaringServiceManagerTransactor{contract: contract}, ContractIncredibleSquaringServiceManagerFilterer: ContractIncredibleSquaringServiceManagerFilterer{contract: contract}}, nil
}

// NewContractIncredibleSquaringServiceManagerCaller creates a new read-only instance of ContractIncredibleSquaringServiceManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIncredibleSquaringServiceManagerCaller, error) {
	contract, err := bindContractIncredibleSquaringServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManagerCaller{contract: contract}, nil
}

// NewContractIncredibleSquaringServiceManagerTransactor creates a new write-only instance of ContractIncredibleSquaringServiceManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIncredibleSquaringServiceManagerTransactor, error) {
	contract, err := bindContractIncredibleSquaringServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManagerTransactor{contract: contract}, nil
}

// NewContractIncredibleSquaringServiceManagerFilterer creates a new log filterer instance of ContractIncredibleSquaringServiceManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIncredibleSquaringServiceManagerFilterer, error) {
	contract, err := bindContractIncredibleSquaringServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManagerFilterer{contract: contract}, nil
}

// bindContractIncredibleSquaringServiceManager binds a generic wrapper to an already deployed contract.
func bindContractIncredibleSquaringServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIncredibleSquaringServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSquaringServiceManager.Contract.ContractIncredibleSquaringServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.ContractIncredibleSquaringServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.ContractIncredibleSquaringServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSquaringServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.AvsDirectory(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.AvsDirectory(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractIncredibleSquaringServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractIncredibleSquaringServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.GetRestakeableStrategies(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.GetRestakeableStrategies(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// IncredibleSquaringTaskManager is a free data retrieval call binding the contract method 0x77ef731d.
//
// Solidity: function incredibleSquaringTaskManager() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) IncredibleSquaringTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "incredibleSquaringTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IncredibleSquaringTaskManager is a free data retrieval call binding the contract method 0x77ef731d.
//
// Solidity: function incredibleSquaringTaskManager() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) IncredibleSquaringTaskManager() (common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.IncredibleSquaringTaskManager(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// IncredibleSquaringTaskManager is a free data retrieval call binding the contract method 0x77ef731d.
//
// Solidity: function incredibleSquaringTaskManager() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) IncredibleSquaringTaskManager() (common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.IncredibleSquaringTaskManager(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.Owner(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.Owner(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractIncredibleSquaringServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractIncredibleSquaringServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.FreezeOperator(&_ContractIncredibleSquaringServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.FreezeOperator(&_ContractIncredibleSquaringServiceManager.TransactOpts, operatorAddr)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.PayForRange(&_ContractIncredibleSquaringServiceManager.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.PayForRange(&_ContractIncredibleSquaringServiceManager.TransactOpts, rangePayments)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.RegisterOperatorToAVS(&_ContractIncredibleSquaringServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.RegisterOperatorToAVS(&_ContractIncredibleSquaringServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.RenounceOwnership(&_ContractIncredibleSquaringServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.RenounceOwnership(&_ContractIncredibleSquaringServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.TransferOwnership(&_ContractIncredibleSquaringServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.TransferOwnership(&_ContractIncredibleSquaringServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.UpdateAVSMetadataURI(&_ContractIncredibleSquaringServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.UpdateAVSMetadataURI(&_ContractIncredibleSquaringServiceManager.TransactOpts, _metadataURI)
}

// ContractIncredibleSquaringServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractIncredibleSquaringServiceManager contract.
type ContractIncredibleSquaringServiceManagerInitializedIterator struct {
	Event *ContractIncredibleSquaringServiceManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringServiceManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringServiceManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringServiceManagerInitialized represents a Initialized event raised by the ContractIncredibleSquaringServiceManager contract.
type ContractIncredibleSquaringServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractIncredibleSquaringServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManagerInitializedIterator{contract: _ContractIncredibleSquaringServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringServiceManagerInitialized)
				if err := _ContractIncredibleSquaringServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractIncredibleSquaringServiceManagerInitialized, error) {
	event := new(ContractIncredibleSquaringServiceManagerInitialized)
	if err := _ContractIncredibleSquaringServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractIncredibleSquaringServiceManager contract.
type ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator struct {
	Event *ContractIncredibleSquaringServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringServiceManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringServiceManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractIncredibleSquaringServiceManager contract.
type ContractIncredibleSquaringServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator{contract: _ContractIncredibleSquaringServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringServiceManagerOwnershipTransferred)
				if err := _ContractIncredibleSquaringServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractIncredibleSquaringServiceManagerOwnershipTransferred, error) {
	event := new(ContractIncredibleSquaringServiceManagerOwnershipTransferred)
	if err := _ContractIncredibleSquaringServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
