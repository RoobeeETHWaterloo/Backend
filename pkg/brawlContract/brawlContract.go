// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package brawlContract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CryptoBrawlABI is the input ABI used to generate the binding from.
const CryptoBrawlABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"challengesList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"temporaryAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fights\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"player1CharID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"player2CharID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"player1GeneralAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"player2GeneralAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"player1TempAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"player2TempAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stepNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastStepBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tempAddress\",\"type\":\"address\"}],\"name\":\"searchFight\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"playerAction1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"playerAction2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oponentAction1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oponentAction2\",\"type\":\"uint256\"}],\"name\":\"calculateDamage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_fightsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fightID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stepNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"player1Action1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"player1Action2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"player2Action1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"player2Action2\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"player1Salt\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"player2Salt\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"player1Signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"player2Signature\",\"type\":\"bytes\"}],\"name\":\"actionSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"chars\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"fightsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"fullHp\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"damage\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"fightId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"currentHP\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"lastFihgtBlockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fightID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stepNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"playerAction1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"playerAction2\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"playerSalt\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"checkAction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"LookingForAFight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player2\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fightId\",\"type\":\"uint256\"}],\"name\":\"FightCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fightId\",\"type\":\"uint256\"}],\"name\":\"FightFinished\",\"type\":\"event\"}]"

// CryptoBrawlFuncSigs maps the 4-byte function signature to its string representation.
var CryptoBrawlFuncSigs = map[string]string{
	"89c4c824": "_fightsCount()",
	"9cfd4490": "actionSet(uint256,uint256,uint256,uint256,uint256,uint256,string,string,bytes,bytes)",
	"78226fbf": "calculateDamage(uint256,uint256,uint256,uint256)",
	"014654ca": "challengesList(uint8)",
	"e7e8fa6e": "chars(bytes32)",
	"f84a7c5b": "checkAction(uint256,uint256,uint256,uint256,string,bytes)",
	"4315e723": "fights(uint256)",
	"52965009": "searchFight(address,uint256,address)",
	"4189a0c8": "temporaryAddresses(address)",
}

// CryptoBrawlBin is the compiled bytecode used for deploying new contracts.
var CryptoBrawlBin = "0x61018060405260016080819052600060a081905260c0819052602d60e0819052600f610100526101208290526101408190526101608290526008805460ff1990811690941790556009829055600a829055600b80548416821761ff001916610f00179055600c829055600d805490931617909155600e5534801561008257600080fd5b50611608806100926000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806378226fbf1161006657806378226fbf146101b457806389c4c824146101f95780639cfd449014610201578063e7e8fa6e14610459578063f84a7c5b146104c257610093565b8063014654ca146100985780634189a0c8146100ca5780634315e7231461010c578063529650091461017c575b600080fd5b6100b8600480360360208110156100ae57600080fd5b503560ff16610603565b60408051918252519081900360200190f35b6100f0600480360360208110156100e057600080fd5b50356001600160a01b0316610615565b604080516001600160a01b039092168252519081900360200190f35b6101296004803603602081101561012257600080fd5b5035610630565b60408051998a5260208a01989098526001600160a01b03968716898901529486166060890152928516608088015290841660a087015260c086015260e08501521661010083015251908190036101200190f35b6101b26004803603606081101561019257600080fd5b506001600160a01b0381358116916020810135916040909101351661068d565b005b6101e3600480360360808110156101ca57600080fd5b5080359060208101359060408101359060600135610904565b6040805160ff9092168252519081900360200190f35b6100b861099d565b6101b2600480360361014081101561021857600080fd5b81359160208101359160408201359160608101359160808201359160a08101359181019060e0810160c0820135600160201b81111561025657600080fd5b82018360208201111561026857600080fd5b803590602001918460018302840111600160201b8311171561028957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156102db57600080fd5b8201836020820111156102ed57600080fd5b803590602001918460018302840111600160201b8311171561030e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561036057600080fd5b82018360208201111561037257600080fd5b803590602001918460018302840111600160201b8311171561039357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156103e557600080fd5b8201836020820111156103f757600080fd5b803590602001918460018302840111600160201b8311171561041857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506109a3945050505050565b6104766004803603602081101561046f57600080fd5b5035611015565b6040805160ff998a1681526020810198909852878101969096529387166060870152918616608086015260a085015290931660c083015260e08201929092529051908190036101000190f35b6100f0600480360360c08110156104d857600080fd5b81359160208101359160408201359160608101359181019060a081016080820135600160201b81111561050a57600080fd5b82018360208201111561051c57600080fd5b803590602001918460018302840111600160201b8311171561053d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561058f57600080fd5b8201836020820111156105a157600080fd5b803590602001918460018302840111600160201b831117156105c257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611060945050505050565b60026020526000908152604090205481565b6003602052600090815260409020546001600160a01b031681565b600160208190526000918252604090912080549181015460028201546003830154600484015460058501546006860154600787015460089097015495966001600160a01b0395861696948616959384169492841693919290911689565b336000908152600360205260409020546001600160a01b0316156106f8576040805162461bcd60e51b815260206004820152601860248201527f616c726561647920686176652061206368616c6c656e67650000000000000000604482015290519081900360640190fd5b604080516bffffffffffffffffffffffff19606086901b1660208083019190915260348083018690528351808403909101815260549092018352815191810191909120600081815291829052919020600401541561075557600080fd5b6000818152602081815260408083205460ff168352600290915290205481141561077e57600080fd5b60008181526020819052604090205460ff1661081a576000818152602081905260409020600854815460ff1990811660ff9283161783556009546001840155600a546002840155600b805460038501805484169185169190911780825591546101009081900485160261ff0019909216919091179055600c546004840155600d546005840180549092169216919091179055600e546006909101555b60008181526004602090815260408083208054336001600160a01b031991821681179092558185526005808552838620879055868652858552838620600380820154928201805460ff94851660ff199091161790559054938752855283862080549092166001600160a01b03891617909155168084526002909252909120546108f35760ff811660008181526002602090815260409182902085905581513381529081019290925280517f17bf03596c4b865d7612f3b7e06d07a29bb027f883d5bef05f5aacb6ba6db8e39281900390910190a16108fd565b6108fd818361111a565b5050505050565b60008085600114806109165750846001145b15610936578360041415801561092d575082600414155b15610936576001015b85600214806109455750846002145b15610965578360051415801561095c575082600514155b15610965576001015b85600314806109745750846003145b15610994578360061415801561098b575082600614155b15610994576001015b95945050505050565b60075481565b336000908152600560209081526040808320548352908290529020600401548a146109cd57600080fd5b8688141580156109dd5750848614155b610a24576040805162461bcd60e51b8152602060048201526013602482015272756e617661696c61626c6520616374696f6e7360681b604482015290519081900360640190fd5b610a2c611587565b5060008a81526001602081815260409283902083516101208101855281548152928101549183019190915260028101546001600160a01b039081169383019390935260038101548316606083015260048101548316608083018190526005820154841660a0840152600682015460c0840152600782015460e084015260089091015490921661010082015290610ac68c8c8c8c8a89611060565b6001600160a01b0316148015610afe57508060a001516001600160a01b0316610af38c8c8a8a8988611060565b6001600160a01b0316145b610b0757600080fd5b8051600090815260208190526040812060030154610100900460ff16610b2f8b8b8b8b610904565b6020848101516000908152908190526040812060030154919092029250610100900460ff16610b608a8a8e8e610904565b60008f81526001602081815260408084204360078201556006018054909301909255875183528290529020600501549102915060ff8083169116111580610bc35750602080840151600090815290819052604090206005015460ff808416911611155b15610fb857825160009081526020819052604090206005015460ff808316911611610d7e576000808460000151815260200190815260200160002060030160009054906101000a900460ff166000808560000151815260200190815260200160002060050160006101000a81548160ff021916908360ff1602179055506000808460200151815260200190815260200160002060030160009054906101000a900460ff166000808560200151815260200190815260200160002060050160006101000a81548160ff021916908360ff1602179055508260600151600160008f815260200190815260200160002060080160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600160008085602001518152602001908152602001600020600201600082825401925050819055507fe91033fa9bf87e86271b8575b191f39d0ee30a536e966baca002d3f77e7bf7bc600160008f815260200190815260200160002060080160009054906101000a90046001600160a01b03168e60405180836001600160a01b03166001600160a01b031681526020018281526020019250505060405180910390a1610f10565b6000808460000151815260200190815260200160002060030160009054906101000a900460ff166000808560000151815260200190815260200160002060050160006101000a81548160ff021916908360ff1602179055506000808460200151815260200190815260200160002060030160009054906101000a900460ff166000808560200151815260200190815260200160002060050160006101000a81548160ff021916908360ff1602179055508260400151600160008f815260200190815260200160002060080160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600160008085600001518152602001908152602001600020600201600082825401925050819055507fe91033fa9bf87e86271b8575b191f39d0ee30a536e966baca002d3f77e7bf7bc600160008f815260200190815260200160002060080160009054906101000a90046001600160a01b03168e60405180836001600160a01b03166001600160a01b031681526020018281526020019250505060405180910390a15b8251600090815260208181526040808320600190810180548201905582870180518552828520820180549092019091558651845281842043600691820181905582518652838620909101558651845281842060049081018590559051845281842001839055808601516001600160a01b039081168452600390925280832080546001600160a01b03199081169091556060870151909216835290912080549091169055611006565b82516000908152602081815260408083206005908101805460ff1980821660ff9283168990038316179092559388015185529190932090920180549283169282168590039091169190911790555b50505050505050505050505050565b600060208190529081526040902080546001820154600283015460038401546004850154600586015460069096015460ff958616969495939483851694610100909404841693169088565b60008087878787876040516020018086815260200185815260200184815260200183815260200182805190602001908083835b602083106110b25780518252601f199092019160209182019101611093565b6001836020036101000a0380198251168184511680821785525050505050509050019550505050505060405160208183030381529060405280519060200120905060006110fe82611463565b9050600061110c82866114b4565b9a9950505050505050505050565b60ff821660009081526002602052604081208054919055600780546001019055611142611587565b6040518061012001604052808381526020018481526020016004600085815260200190815260200160002060009054906101000a90046001600160a01b03166001600160a01b03168152602001336001600160a01b03168152602001600360006004600087815260200190815260200160002060009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03166001600160a01b0316815260200160036000336001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03166001600160a01b031681526020016001815260200143815260200160006001600160a01b0316815250905080600160006007548152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060608201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060808201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060a08201518160050160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c0820151816006015560e082015181600701556101008201518160080160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555090505060075460008084815260200190815260200160002060040181905550600754600080858152602001908152602001600020600401819055507ff36f655ac60ebec92e72542503ad92977efc077913818ec39235b90baf49da856004600084815260200190815260200160002060009054906101000a90046001600160a01b03163360075460405180846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b03168152602001828152602001935050505060405180910390a150505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b60008060008084516041146114cf5760009350505050611581565b50505060208201516040830151606084015160001a601b8110156114f157601b015b8060ff16601b1415801561150957508060ff16601c14155b1561151a5760009350505050611581565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa158015611571573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b6040805161012081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101919091529056fea265627a7a7231582040177ee2b9a724bf9e17de64ecbf260ef21423bacbf17f532a7fc5a0ef0db59764736f6c634300050b0032"

// DeployCryptoBrawl deploys a new Ethereum contract, binding an instance of CryptoBrawl to it.
func DeployCryptoBrawl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CryptoBrawl, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoBrawlABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CryptoBrawlBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CryptoBrawl{CryptoBrawlCaller: CryptoBrawlCaller{contract: contract}, CryptoBrawlTransactor: CryptoBrawlTransactor{contract: contract}, CryptoBrawlFilterer: CryptoBrawlFilterer{contract: contract}}, nil
}

// CryptoBrawl is an auto generated Go binding around an Ethereum contract.
type CryptoBrawl struct {
	CryptoBrawlCaller     // Read-only binding to the contract
	CryptoBrawlTransactor // Write-only binding to the contract
	CryptoBrawlFilterer   // Log filterer for contract events
}

// CryptoBrawlCaller is an auto generated read-only Go binding around an Ethereum contract.
type CryptoBrawlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoBrawlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CryptoBrawlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoBrawlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CryptoBrawlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoBrawlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CryptoBrawlSession struct {
	Contract     *CryptoBrawl      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CryptoBrawlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CryptoBrawlCallerSession struct {
	Contract *CryptoBrawlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CryptoBrawlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CryptoBrawlTransactorSession struct {
	Contract     *CryptoBrawlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CryptoBrawlRaw is an auto generated low-level Go binding around an Ethereum contract.
type CryptoBrawlRaw struct {
	Contract *CryptoBrawl // Generic contract binding to access the raw methods on
}

// CryptoBrawlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CryptoBrawlCallerRaw struct {
	Contract *CryptoBrawlCaller // Generic read-only contract binding to access the raw methods on
}

// CryptoBrawlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CryptoBrawlTransactorRaw struct {
	Contract *CryptoBrawlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCryptoBrawl creates a new instance of CryptoBrawl, bound to a specific deployed contract.
func NewCryptoBrawl(address common.Address, backend bind.ContractBackend) (*CryptoBrawl, error) {
	contract, err := bindCryptoBrawl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CryptoBrawl{CryptoBrawlCaller: CryptoBrawlCaller{contract: contract}, CryptoBrawlTransactor: CryptoBrawlTransactor{contract: contract}, CryptoBrawlFilterer: CryptoBrawlFilterer{contract: contract}}, nil
}

// NewCryptoBrawlCaller creates a new read-only instance of CryptoBrawl, bound to a specific deployed contract.
func NewCryptoBrawlCaller(address common.Address, caller bind.ContractCaller) (*CryptoBrawlCaller, error) {
	contract, err := bindCryptoBrawl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoBrawlCaller{contract: contract}, nil
}

// NewCryptoBrawlTransactor creates a new write-only instance of CryptoBrawl, bound to a specific deployed contract.
func NewCryptoBrawlTransactor(address common.Address, transactor bind.ContractTransactor) (*CryptoBrawlTransactor, error) {
	contract, err := bindCryptoBrawl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoBrawlTransactor{contract: contract}, nil
}

// NewCryptoBrawlFilterer creates a new log filterer instance of CryptoBrawl, bound to a specific deployed contract.
func NewCryptoBrawlFilterer(address common.Address, filterer bind.ContractFilterer) (*CryptoBrawlFilterer, error) {
	contract, err := bindCryptoBrawl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CryptoBrawlFilterer{contract: contract}, nil
}

// bindCryptoBrawl binds a generic wrapper to an already deployed contract.
func bindCryptoBrawl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoBrawlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoBrawl *CryptoBrawlRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CryptoBrawl.Contract.CryptoBrawlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoBrawl *CryptoBrawlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoBrawl.Contract.CryptoBrawlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoBrawl *CryptoBrawlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoBrawl.Contract.CryptoBrawlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoBrawl *CryptoBrawlCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CryptoBrawl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoBrawl *CryptoBrawlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoBrawl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoBrawl *CryptoBrawlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoBrawl.Contract.contract.Transact(opts, method, params...)
}

// FightsCount is a free data retrieval call binding the contract method 0x89c4c824.
//
// Solidity: function _fightsCount() constant returns(uint256)
func (_CryptoBrawl *CryptoBrawlCaller) FightsCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CryptoBrawl.contract.Call(opts, out, "_fightsCount")
	return *ret0, err
}

// FightsCount is a free data retrieval call binding the contract method 0x89c4c824.
//
// Solidity: function _fightsCount() constant returns(uint256)
func (_CryptoBrawl *CryptoBrawlSession) FightsCount() (*big.Int, error) {
	return _CryptoBrawl.Contract.FightsCount(&_CryptoBrawl.CallOpts)
}

// FightsCount is a free data retrieval call binding the contract method 0x89c4c824.
//
// Solidity: function _fightsCount() constant returns(uint256)
func (_CryptoBrawl *CryptoBrawlCallerSession) FightsCount() (*big.Int, error) {
	return _CryptoBrawl.Contract.FightsCount(&_CryptoBrawl.CallOpts)
}

// CalculateDamage is a free data retrieval call binding the contract method 0x78226fbf.
//
// Solidity: function calculateDamage(uint256 playerAction1, uint256 playerAction2, uint256 oponentAction1, uint256 oponentAction2) constant returns(uint8)
func (_CryptoBrawl *CryptoBrawlCaller) CalculateDamage(opts *bind.CallOpts, playerAction1 *big.Int, playerAction2 *big.Int, oponentAction1 *big.Int, oponentAction2 *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _CryptoBrawl.contract.Call(opts, out, "calculateDamage", playerAction1, playerAction2, oponentAction1, oponentAction2)
	return *ret0, err
}

// CalculateDamage is a free data retrieval call binding the contract method 0x78226fbf.
//
// Solidity: function calculateDamage(uint256 playerAction1, uint256 playerAction2, uint256 oponentAction1, uint256 oponentAction2) constant returns(uint8)
func (_CryptoBrawl *CryptoBrawlSession) CalculateDamage(playerAction1 *big.Int, playerAction2 *big.Int, oponentAction1 *big.Int, oponentAction2 *big.Int) (uint8, error) {
	return _CryptoBrawl.Contract.CalculateDamage(&_CryptoBrawl.CallOpts, playerAction1, playerAction2, oponentAction1, oponentAction2)
}

// CalculateDamage is a free data retrieval call binding the contract method 0x78226fbf.
//
// Solidity: function calculateDamage(uint256 playerAction1, uint256 playerAction2, uint256 oponentAction1, uint256 oponentAction2) constant returns(uint8)
func (_CryptoBrawl *CryptoBrawlCallerSession) CalculateDamage(playerAction1 *big.Int, playerAction2 *big.Int, oponentAction1 *big.Int, oponentAction2 *big.Int) (uint8, error) {
	return _CryptoBrawl.Contract.CalculateDamage(&_CryptoBrawl.CallOpts, playerAction1, playerAction2, oponentAction1, oponentAction2)
}

// ChallengesList is a free data retrieval call binding the contract method 0x014654ca.
//
// Solidity: function challengesList(uint8 ) constant returns(bytes32)
func (_CryptoBrawl *CryptoBrawlCaller) ChallengesList(opts *bind.CallOpts, arg0 uint8) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _CryptoBrawl.contract.Call(opts, out, "challengesList", arg0)
	return *ret0, err
}

// ChallengesList is a free data retrieval call binding the contract method 0x014654ca.
//
// Solidity: function challengesList(uint8 ) constant returns(bytes32)
func (_CryptoBrawl *CryptoBrawlSession) ChallengesList(arg0 uint8) ([32]byte, error) {
	return _CryptoBrawl.Contract.ChallengesList(&_CryptoBrawl.CallOpts, arg0)
}

// ChallengesList is a free data retrieval call binding the contract method 0x014654ca.
//
// Solidity: function challengesList(uint8 ) constant returns(bytes32)
func (_CryptoBrawl *CryptoBrawlCallerSession) ChallengesList(arg0 uint8) ([32]byte, error) {
	return _CryptoBrawl.Contract.ChallengesList(&_CryptoBrawl.CallOpts, arg0)
}

// Chars is a free data retrieval call binding the contract method 0xe7e8fa6e.
//
// Solidity: function chars(bytes32 ) constant returns(uint8 level, uint256 fightsCount, uint256 winsCount, uint8 fullHp, uint8 damage, uint256 fightId, uint8 currentHP, uint256 lastFihgtBlockNumber)
func (_CryptoBrawl *CryptoBrawlCaller) Chars(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Level                uint8
	FightsCount          *big.Int
	WinsCount            *big.Int
	FullHp               uint8
	Damage               uint8
	FightId              *big.Int
	CurrentHP            uint8
	LastFihgtBlockNumber *big.Int
}, error) {
	ret := new(struct {
		Level                uint8
		FightsCount          *big.Int
		WinsCount            *big.Int
		FullHp               uint8
		Damage               uint8
		FightId              *big.Int
		CurrentHP            uint8
		LastFihgtBlockNumber *big.Int
	})
	out := ret
	err := _CryptoBrawl.contract.Call(opts, out, "chars", arg0)
	return *ret, err
}

// Chars is a free data retrieval call binding the contract method 0xe7e8fa6e.
//
// Solidity: function chars(bytes32 ) constant returns(uint8 level, uint256 fightsCount, uint256 winsCount, uint8 fullHp, uint8 damage, uint256 fightId, uint8 currentHP, uint256 lastFihgtBlockNumber)
func (_CryptoBrawl *CryptoBrawlSession) Chars(arg0 [32]byte) (struct {
	Level                uint8
	FightsCount          *big.Int
	WinsCount            *big.Int
	FullHp               uint8
	Damage               uint8
	FightId              *big.Int
	CurrentHP            uint8
	LastFihgtBlockNumber *big.Int
}, error) {
	return _CryptoBrawl.Contract.Chars(&_CryptoBrawl.CallOpts, arg0)
}

// Chars is a free data retrieval call binding the contract method 0xe7e8fa6e.
//
// Solidity: function chars(bytes32 ) constant returns(uint8 level, uint256 fightsCount, uint256 winsCount, uint8 fullHp, uint8 damage, uint256 fightId, uint8 currentHP, uint256 lastFihgtBlockNumber)
func (_CryptoBrawl *CryptoBrawlCallerSession) Chars(arg0 [32]byte) (struct {
	Level                uint8
	FightsCount          *big.Int
	WinsCount            *big.Int
	FullHp               uint8
	Damage               uint8
	FightId              *big.Int
	CurrentHP            uint8
	LastFihgtBlockNumber *big.Int
}, error) {
	return _CryptoBrawl.Contract.Chars(&_CryptoBrawl.CallOpts, arg0)
}

// CheckAction is a free data retrieval call binding the contract method 0xf84a7c5b.
//
// Solidity: function checkAction(uint256 fightID, uint256 stepNum, uint256 playerAction1, uint256 playerAction2, string playerSalt, bytes signature) constant returns(address)
func (_CryptoBrawl *CryptoBrawlCaller) CheckAction(opts *bind.CallOpts, fightID *big.Int, stepNum *big.Int, playerAction1 *big.Int, playerAction2 *big.Int, playerSalt string, signature []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoBrawl.contract.Call(opts, out, "checkAction", fightID, stepNum, playerAction1, playerAction2, playerSalt, signature)
	return *ret0, err
}

// CheckAction is a free data retrieval call binding the contract method 0xf84a7c5b.
//
// Solidity: function checkAction(uint256 fightID, uint256 stepNum, uint256 playerAction1, uint256 playerAction2, string playerSalt, bytes signature) constant returns(address)
func (_CryptoBrawl *CryptoBrawlSession) CheckAction(fightID *big.Int, stepNum *big.Int, playerAction1 *big.Int, playerAction2 *big.Int, playerSalt string, signature []byte) (common.Address, error) {
	return _CryptoBrawl.Contract.CheckAction(&_CryptoBrawl.CallOpts, fightID, stepNum, playerAction1, playerAction2, playerSalt, signature)
}

// CheckAction is a free data retrieval call binding the contract method 0xf84a7c5b.
//
// Solidity: function checkAction(uint256 fightID, uint256 stepNum, uint256 playerAction1, uint256 playerAction2, string playerSalt, bytes signature) constant returns(address)
func (_CryptoBrawl *CryptoBrawlCallerSession) CheckAction(fightID *big.Int, stepNum *big.Int, playerAction1 *big.Int, playerAction2 *big.Int, playerSalt string, signature []byte) (common.Address, error) {
	return _CryptoBrawl.Contract.CheckAction(&_CryptoBrawl.CallOpts, fightID, stepNum, playerAction1, playerAction2, playerSalt, signature)
}

// Fights is a free data retrieval call binding the contract method 0x4315e723.
//
// Solidity: function fights(uint256 ) constant returns(bytes32 player1CharID, bytes32 player2CharID, address player1GeneralAddress, address player2GeneralAddress, address player1TempAddress, address player2TempAddress, uint256 stepNum, uint256 lastStepBlock, address winner)
func (_CryptoBrawl *CryptoBrawlCaller) Fights(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Player1CharID         [32]byte
	Player2CharID         [32]byte
	Player1GeneralAddress common.Address
	Player2GeneralAddress common.Address
	Player1TempAddress    common.Address
	Player2TempAddress    common.Address
	StepNum               *big.Int
	LastStepBlock         *big.Int
	Winner                common.Address
}, error) {
	ret := new(struct {
		Player1CharID         [32]byte
		Player2CharID         [32]byte
		Player1GeneralAddress common.Address
		Player2GeneralAddress common.Address
		Player1TempAddress    common.Address
		Player2TempAddress    common.Address
		StepNum               *big.Int
		LastStepBlock         *big.Int
		Winner                common.Address
	})
	out := ret
	err := _CryptoBrawl.contract.Call(opts, out, "fights", arg0)
	return *ret, err
}

// Fights is a free data retrieval call binding the contract method 0x4315e723.
//
// Solidity: function fights(uint256 ) constant returns(bytes32 player1CharID, bytes32 player2CharID, address player1GeneralAddress, address player2GeneralAddress, address player1TempAddress, address player2TempAddress, uint256 stepNum, uint256 lastStepBlock, address winner)
func (_CryptoBrawl *CryptoBrawlSession) Fights(arg0 *big.Int) (struct {
	Player1CharID         [32]byte
	Player2CharID         [32]byte
	Player1GeneralAddress common.Address
	Player2GeneralAddress common.Address
	Player1TempAddress    common.Address
	Player2TempAddress    common.Address
	StepNum               *big.Int
	LastStepBlock         *big.Int
	Winner                common.Address
}, error) {
	return _CryptoBrawl.Contract.Fights(&_CryptoBrawl.CallOpts, arg0)
}

// Fights is a free data retrieval call binding the contract method 0x4315e723.
//
// Solidity: function fights(uint256 ) constant returns(bytes32 player1CharID, bytes32 player2CharID, address player1GeneralAddress, address player2GeneralAddress, address player1TempAddress, address player2TempAddress, uint256 stepNum, uint256 lastStepBlock, address winner)
func (_CryptoBrawl *CryptoBrawlCallerSession) Fights(arg0 *big.Int) (struct {
	Player1CharID         [32]byte
	Player2CharID         [32]byte
	Player1GeneralAddress common.Address
	Player2GeneralAddress common.Address
	Player1TempAddress    common.Address
	Player2TempAddress    common.Address
	StepNum               *big.Int
	LastStepBlock         *big.Int
	Winner                common.Address
}, error) {
	return _CryptoBrawl.Contract.Fights(&_CryptoBrawl.CallOpts, arg0)
}

// TemporaryAddresses is a free data retrieval call binding the contract method 0x4189a0c8.
//
// Solidity: function temporaryAddresses(address ) constant returns(address)
func (_CryptoBrawl *CryptoBrawlCaller) TemporaryAddresses(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoBrawl.contract.Call(opts, out, "temporaryAddresses", arg0)
	return *ret0, err
}

// TemporaryAddresses is a free data retrieval call binding the contract method 0x4189a0c8.
//
// Solidity: function temporaryAddresses(address ) constant returns(address)
func (_CryptoBrawl *CryptoBrawlSession) TemporaryAddresses(arg0 common.Address) (common.Address, error) {
	return _CryptoBrawl.Contract.TemporaryAddresses(&_CryptoBrawl.CallOpts, arg0)
}

// TemporaryAddresses is a free data retrieval call binding the contract method 0x4189a0c8.
//
// Solidity: function temporaryAddresses(address ) constant returns(address)
func (_CryptoBrawl *CryptoBrawlCallerSession) TemporaryAddresses(arg0 common.Address) (common.Address, error) {
	return _CryptoBrawl.Contract.TemporaryAddresses(&_CryptoBrawl.CallOpts, arg0)
}

// ActionSet is a paid mutator transaction binding the contract method 0x9cfd4490.
//
// Solidity: function actionSet(uint256 fightID, uint256 stepNum, uint256 player1Action1, uint256 player1Action2, uint256 player2Action1, uint256 player2Action2, string player1Salt, string player2Salt, bytes player1Signature, bytes player2Signature) returns()
func (_CryptoBrawl *CryptoBrawlTransactor) ActionSet(opts *bind.TransactOpts, fightID *big.Int, stepNum *big.Int, player1Action1 *big.Int, player1Action2 *big.Int, player2Action1 *big.Int, player2Action2 *big.Int, player1Salt string, player2Salt string, player1Signature []byte, player2Signature []byte) (*types.Transaction, error) {
	return _CryptoBrawl.contract.Transact(opts, "actionSet", fightID, stepNum, player1Action1, player1Action2, player2Action1, player2Action2, player1Salt, player2Salt, player1Signature, player2Signature)
}

// ActionSet is a paid mutator transaction binding the contract method 0x9cfd4490.
//
// Solidity: function actionSet(uint256 fightID, uint256 stepNum, uint256 player1Action1, uint256 player1Action2, uint256 player2Action1, uint256 player2Action2, string player1Salt, string player2Salt, bytes player1Signature, bytes player2Signature) returns()
func (_CryptoBrawl *CryptoBrawlSession) ActionSet(fightID *big.Int, stepNum *big.Int, player1Action1 *big.Int, player1Action2 *big.Int, player2Action1 *big.Int, player2Action2 *big.Int, player1Salt string, player2Salt string, player1Signature []byte, player2Signature []byte) (*types.Transaction, error) {
	return _CryptoBrawl.Contract.ActionSet(&_CryptoBrawl.TransactOpts, fightID, stepNum, player1Action1, player1Action2, player2Action1, player2Action2, player1Salt, player2Salt, player1Signature, player2Signature)
}

// ActionSet is a paid mutator transaction binding the contract method 0x9cfd4490.
//
// Solidity: function actionSet(uint256 fightID, uint256 stepNum, uint256 player1Action1, uint256 player1Action2, uint256 player2Action1, uint256 player2Action2, string player1Salt, string player2Salt, bytes player1Signature, bytes player2Signature) returns()
func (_CryptoBrawl *CryptoBrawlTransactorSession) ActionSet(fightID *big.Int, stepNum *big.Int, player1Action1 *big.Int, player1Action2 *big.Int, player2Action1 *big.Int, player2Action2 *big.Int, player1Salt string, player2Salt string, player1Signature []byte, player2Signature []byte) (*types.Transaction, error) {
	return _CryptoBrawl.Contract.ActionSet(&_CryptoBrawl.TransactOpts, fightID, stepNum, player1Action1, player1Action2, player2Action1, player2Action2, player1Salt, player2Salt, player1Signature, player2Signature)
}

// SearchFight is a paid mutator transaction binding the contract method 0x52965009.
//
// Solidity: function searchFight(address ERC721, uint256 tokenID, address tempAddress) returns()
func (_CryptoBrawl *CryptoBrawlTransactor) SearchFight(opts *bind.TransactOpts, ERC721 common.Address, tokenID *big.Int, tempAddress common.Address) (*types.Transaction, error) {
	return _CryptoBrawl.contract.Transact(opts, "searchFight", ERC721, tokenID, tempAddress)
}

// SearchFight is a paid mutator transaction binding the contract method 0x52965009.
//
// Solidity: function searchFight(address ERC721, uint256 tokenID, address tempAddress) returns()
func (_CryptoBrawl *CryptoBrawlSession) SearchFight(ERC721 common.Address, tokenID *big.Int, tempAddress common.Address) (*types.Transaction, error) {
	return _CryptoBrawl.Contract.SearchFight(&_CryptoBrawl.TransactOpts, ERC721, tokenID, tempAddress)
}

// SearchFight is a paid mutator transaction binding the contract method 0x52965009.
//
// Solidity: function searchFight(address ERC721, uint256 tokenID, address tempAddress) returns()
func (_CryptoBrawl *CryptoBrawlTransactorSession) SearchFight(ERC721 common.Address, tokenID *big.Int, tempAddress common.Address) (*types.Transaction, error) {
	return _CryptoBrawl.Contract.SearchFight(&_CryptoBrawl.TransactOpts, ERC721, tokenID, tempAddress)
}

// CryptoBrawlFightCreatedIterator is returned from FilterFightCreated and is used to iterate over the raw logs and unpacked data for FightCreated events raised by the CryptoBrawl contract.
type CryptoBrawlFightCreatedIterator struct {
	Event *CryptoBrawlFightCreated // Event containing the contract specifics and raw log

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
func (it *CryptoBrawlFightCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoBrawlFightCreated)
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
		it.Event = new(CryptoBrawlFightCreated)
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
func (it *CryptoBrawlFightCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoBrawlFightCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoBrawlFightCreated represents a FightCreated event raised by the CryptoBrawl contract.
type CryptoBrawlFightCreated struct {
	Player1 common.Address
	Player2 common.Address
	FightId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFightCreated is a free log retrieval operation binding the contract event 0xf36f655ac60ebec92e72542503ad92977efc077913818ec39235b90baf49da85.
//
// Solidity: event FightCreated(address player1, address player2, uint256 fightId)
func (_CryptoBrawl *CryptoBrawlFilterer) FilterFightCreated(opts *bind.FilterOpts) (*CryptoBrawlFightCreatedIterator, error) {

	logs, sub, err := _CryptoBrawl.contract.FilterLogs(opts, "FightCreated")
	if err != nil {
		return nil, err
	}
	return &CryptoBrawlFightCreatedIterator{contract: _CryptoBrawl.contract, event: "FightCreated", logs: logs, sub: sub}, nil
}

// WatchFightCreated is a free log subscription operation binding the contract event 0xf36f655ac60ebec92e72542503ad92977efc077913818ec39235b90baf49da85.
//
// Solidity: event FightCreated(address player1, address player2, uint256 fightId)
func (_CryptoBrawl *CryptoBrawlFilterer) WatchFightCreated(opts *bind.WatchOpts, sink chan<- *CryptoBrawlFightCreated) (event.Subscription, error) {

	logs, sub, err := _CryptoBrawl.contract.WatchLogs(opts, "FightCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoBrawlFightCreated)
				if err := _CryptoBrawl.contract.UnpackLog(event, "FightCreated", log); err != nil {
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

// ParseFightCreated is a log parse operation binding the contract event 0xf36f655ac60ebec92e72542503ad92977efc077913818ec39235b90baf49da85.
//
// Solidity: event FightCreated(address player1, address player2, uint256 fightId)
func (_CryptoBrawl *CryptoBrawlFilterer) ParseFightCreated(log types.Log) (*CryptoBrawlFightCreated, error) {
	event := new(CryptoBrawlFightCreated)
	if err := _CryptoBrawl.contract.UnpackLog(event, "FightCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CryptoBrawlFightFinishedIterator is returned from FilterFightFinished and is used to iterate over the raw logs and unpacked data for FightFinished events raised by the CryptoBrawl contract.
type CryptoBrawlFightFinishedIterator struct {
	Event *CryptoBrawlFightFinished // Event containing the contract specifics and raw log

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
func (it *CryptoBrawlFightFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoBrawlFightFinished)
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
		it.Event = new(CryptoBrawlFightFinished)
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
func (it *CryptoBrawlFightFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoBrawlFightFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoBrawlFightFinished represents a FightFinished event raised by the CryptoBrawl contract.
type CryptoBrawlFightFinished struct {
	Winner  common.Address
	FightId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFightFinished is a free log retrieval operation binding the contract event 0xe91033fa9bf87e86271b8575b191f39d0ee30a536e966baca002d3f77e7bf7bc.
//
// Solidity: event FightFinished(address winner, uint256 fightId)
func (_CryptoBrawl *CryptoBrawlFilterer) FilterFightFinished(opts *bind.FilterOpts) (*CryptoBrawlFightFinishedIterator, error) {

	logs, sub, err := _CryptoBrawl.contract.FilterLogs(opts, "FightFinished")
	if err != nil {
		return nil, err
	}
	return &CryptoBrawlFightFinishedIterator{contract: _CryptoBrawl.contract, event: "FightFinished", logs: logs, sub: sub}, nil
}

// WatchFightFinished is a free log subscription operation binding the contract event 0xe91033fa9bf87e86271b8575b191f39d0ee30a536e966baca002d3f77e7bf7bc.
//
// Solidity: event FightFinished(address winner, uint256 fightId)
func (_CryptoBrawl *CryptoBrawlFilterer) WatchFightFinished(opts *bind.WatchOpts, sink chan<- *CryptoBrawlFightFinished) (event.Subscription, error) {

	logs, sub, err := _CryptoBrawl.contract.WatchLogs(opts, "FightFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoBrawlFightFinished)
				if err := _CryptoBrawl.contract.UnpackLog(event, "FightFinished", log); err != nil {
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

// ParseFightFinished is a log parse operation binding the contract event 0xe91033fa9bf87e86271b8575b191f39d0ee30a536e966baca002d3f77e7bf7bc.
//
// Solidity: event FightFinished(address winner, uint256 fightId)
func (_CryptoBrawl *CryptoBrawlFilterer) ParseFightFinished(log types.Log) (*CryptoBrawlFightFinished, error) {
	event := new(CryptoBrawlFightFinished)
	if err := _CryptoBrawl.contract.UnpackLog(event, "FightFinished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CryptoBrawlLookingForAFightIterator is returned from FilterLookingForAFight and is used to iterate over the raw logs and unpacked data for LookingForAFight events raised by the CryptoBrawl contract.
type CryptoBrawlLookingForAFightIterator struct {
	Event *CryptoBrawlLookingForAFight // Event containing the contract specifics and raw log

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
func (it *CryptoBrawlLookingForAFightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoBrawlLookingForAFight)
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
		it.Event = new(CryptoBrawlLookingForAFight)
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
func (it *CryptoBrawlLookingForAFightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoBrawlLookingForAFightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoBrawlLookingForAFight represents a LookingForAFight event raised by the CryptoBrawl contract.
type CryptoBrawlLookingForAFight struct {
	Player common.Address
	Level  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLookingForAFight is a free log retrieval operation binding the contract event 0x17bf03596c4b865d7612f3b7e06d07a29bb027f883d5bef05f5aacb6ba6db8e3.
//
// Solidity: event LookingForAFight(address player, uint256 level)
func (_CryptoBrawl *CryptoBrawlFilterer) FilterLookingForAFight(opts *bind.FilterOpts) (*CryptoBrawlLookingForAFightIterator, error) {

	logs, sub, err := _CryptoBrawl.contract.FilterLogs(opts, "LookingForAFight")
	if err != nil {
		return nil, err
	}
	return &CryptoBrawlLookingForAFightIterator{contract: _CryptoBrawl.contract, event: "LookingForAFight", logs: logs, sub: sub}, nil
}

// WatchLookingForAFight is a free log subscription operation binding the contract event 0x17bf03596c4b865d7612f3b7e06d07a29bb027f883d5bef05f5aacb6ba6db8e3.
//
// Solidity: event LookingForAFight(address player, uint256 level)
func (_CryptoBrawl *CryptoBrawlFilterer) WatchLookingForAFight(opts *bind.WatchOpts, sink chan<- *CryptoBrawlLookingForAFight) (event.Subscription, error) {

	logs, sub, err := _CryptoBrawl.contract.WatchLogs(opts, "LookingForAFight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoBrawlLookingForAFight)
				if err := _CryptoBrawl.contract.UnpackLog(event, "LookingForAFight", log); err != nil {
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

// ParseLookingForAFight is a log parse operation binding the contract event 0x17bf03596c4b865d7612f3b7e06d07a29bb027f883d5bef05f5aacb6ba6db8e3.
//
// Solidity: event LookingForAFight(address player, uint256 level)
func (_CryptoBrawl *CryptoBrawlFilterer) ParseLookingForAFight(log types.Log) (*CryptoBrawlLookingForAFight, error) {
	event := new(CryptoBrawlLookingForAFight)
	if err := _CryptoBrawl.contract.UnpackLog(event, "LookingForAFight", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SignatureVerificationABI is the input ABI used to generate the binding from.
const SignatureVerificationABI = "[]"

// SignatureVerificationBin is the compiled bytecode used for deploying new contracts.
var SignatureVerificationBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a7231582074e9b98d94a1a4f41cd60ab2aa8b7f999ee2faca334708a8ccbc1b2393acb77164736f6c634300050b0032"

// DeploySignatureVerification deploys a new Ethereum contract, binding an instance of SignatureVerification to it.
func DeploySignatureVerification(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SignatureVerification, error) {
	parsed, err := abi.JSON(strings.NewReader(SignatureVerificationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SignatureVerificationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignatureVerification{SignatureVerificationCaller: SignatureVerificationCaller{contract: contract}, SignatureVerificationTransactor: SignatureVerificationTransactor{contract: contract}, SignatureVerificationFilterer: SignatureVerificationFilterer{contract: contract}}, nil
}

// SignatureVerification is an auto generated Go binding around an Ethereum contract.
type SignatureVerification struct {
	SignatureVerificationCaller     // Read-only binding to the contract
	SignatureVerificationTransactor // Write-only binding to the contract
	SignatureVerificationFilterer   // Log filterer for contract events
}

// SignatureVerificationCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignatureVerificationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureVerificationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignatureVerificationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureVerificationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignatureVerificationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureVerificationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignatureVerificationSession struct {
	Contract     *SignatureVerification // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SignatureVerificationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignatureVerificationCallerSession struct {
	Contract *SignatureVerificationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// SignatureVerificationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignatureVerificationTransactorSession struct {
	Contract     *SignatureVerificationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SignatureVerificationRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignatureVerificationRaw struct {
	Contract *SignatureVerification // Generic contract binding to access the raw methods on
}

// SignatureVerificationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignatureVerificationCallerRaw struct {
	Contract *SignatureVerificationCaller // Generic read-only contract binding to access the raw methods on
}

// SignatureVerificationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignatureVerificationTransactorRaw struct {
	Contract *SignatureVerificationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignatureVerification creates a new instance of SignatureVerification, bound to a specific deployed contract.
func NewSignatureVerification(address common.Address, backend bind.ContractBackend) (*SignatureVerification, error) {
	contract, err := bindSignatureVerification(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignatureVerification{SignatureVerificationCaller: SignatureVerificationCaller{contract: contract}, SignatureVerificationTransactor: SignatureVerificationTransactor{contract: contract}, SignatureVerificationFilterer: SignatureVerificationFilterer{contract: contract}}, nil
}

// NewSignatureVerificationCaller creates a new read-only instance of SignatureVerification, bound to a specific deployed contract.
func NewSignatureVerificationCaller(address common.Address, caller bind.ContractCaller) (*SignatureVerificationCaller, error) {
	contract, err := bindSignatureVerification(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureVerificationCaller{contract: contract}, nil
}

// NewSignatureVerificationTransactor creates a new write-only instance of SignatureVerification, bound to a specific deployed contract.
func NewSignatureVerificationTransactor(address common.Address, transactor bind.ContractTransactor) (*SignatureVerificationTransactor, error) {
	contract, err := bindSignatureVerification(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureVerificationTransactor{contract: contract}, nil
}

// NewSignatureVerificationFilterer creates a new log filterer instance of SignatureVerification, bound to a specific deployed contract.
func NewSignatureVerificationFilterer(address common.Address, filterer bind.ContractFilterer) (*SignatureVerificationFilterer, error) {
	contract, err := bindSignatureVerification(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignatureVerificationFilterer{contract: contract}, nil
}

// bindSignatureVerification binds a generic wrapper to an already deployed contract.
func bindSignatureVerification(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignatureVerificationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureVerification *SignatureVerificationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SignatureVerification.Contract.SignatureVerificationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureVerification *SignatureVerificationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureVerification.Contract.SignatureVerificationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureVerification *SignatureVerificationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureVerification.Contract.SignatureVerificationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureVerification *SignatureVerificationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SignatureVerification.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureVerification *SignatureVerificationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureVerification.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureVerification *SignatureVerificationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureVerification.Contract.contract.Transact(opts, method, params...)
}
