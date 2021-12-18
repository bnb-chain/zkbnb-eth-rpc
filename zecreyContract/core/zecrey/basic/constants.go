package basic

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"log"
	"math/big"
)

const (
	Bytes32Len                = 32
	OnChainOpsTreeDepth       = 5
	OnChainOpsTreeHelperDepth = 4

	GovernanceContractFilePrefix      = "contract_Governance_"
	AssetGovernanceContractFilePrefix = "contract_AssetGovernance_"
	VerifierContractFilePrefix        = "contract_Verifier_"
	ZecreyContractFilePrefix          = "contract_Zecrey_"

	ContractInfoFileName = "contract_addresses.json"

	SuperAddress = "0xE9b15a2D396B349ABF60e53ec66Bcf9af262D449"
	SuperSk      = "acbaa269bd7573ff12361be4b97201aef019776ea13384681d4e5ba6a88367d9"

	ERC20_REY_Name   = "Zecrey Token"
	ERC20_REY_Symbol = "REY"

	ERC20_ETH_Name   = "ETH Mock"
	ERC20_ETH_Symbol = "ETH"

	ERC20_MATIC_Name   = "MATIC Mock"
	ERC20_MATIC_Symbol = "MATIC"

	ERC20_NEAR_Name   = "NEAR Mock"
	ERC20_NEAR_Symbol = "NEAR"

	ERC20_AVAX_Name   = "AVAX Mock"
	ERC20_AVAX_Symbol = "AVAX"

	ERC20_BIT_Name   = "BIT Mock"
	ERC20_BIT_Symbol = "BIT"

	ERC20_USDT_Name   = "USDT Mock"
	ERC20_USDT_Symbol = "USDT"

	ERC20_USDC_Name   = "USDC Mock"
	ERC20_USDC_Symbol = "USDC"

	ERC20_DAI_Name   = "DAI Mock"
	ERC20_DAI_Symbol = "DAI"

	ZecreyMetaData =

	ERC20_BNB_Name   = "BNB Mock"
	ERC20_BNB_Symbol = "BNB"
)

type (
	ProviderClient = _rpc.ProviderClient
	AuthClient     = _rpc.AuthClient
)

var (
	AddressType, _      = abi.NewType("address", "", nil)
	Bytes32Type, _      = abi.NewType("bytes32", "", nil)
	Uint32Type, _       = abi.NewType("uint32", "", nil)
	Uint16Type, _       = abi.NewType("uint16", "", nil)
	Uint256Type, _      = abi.NewType("uint256", "", nil)
	MerkleHelperType, _ = abi.NewType("bool[4]", "", nil)

	InitialSupply, _ = new(big.Int).SetString("100000000000000000000000000", 10)
)

func init() {
	// set log info
	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds)
}

type ZecreyContracts struct {
	GovernanceAddr      string
	AssetGovernanceAddr string
	VerifierAddr        string
	ZecreyAddr          string
	ReyERC20Addr        string
	EthERC20Addr        string
	MaticERC20Addr      string
	NearERC20Addr       string
	AvaxERC20Addr       string
	BitERC20Addr        string
	UsdtERC20Addr       string
	UsdcERC20Addr       string
	DaiERC20Addr        string
	BnbERC20Addr        string
}
