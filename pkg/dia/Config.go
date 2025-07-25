package dia

import (
	"errors"
	"os/user"
	"strings"

	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/tkanos/gonfig"
)

const (
	ArthswapExchange                   = "Arthswap"
	DiffusionExchange                  = "Diffusion"
	OmniDexExchange                    = "OmniDex"
	NetswapExchange                    = "Netswap"
	TethysExchange                     = "Tethys"
	HermesExchange                     = "Hermes"
	AnyswapExchange                    = "Anyswap"
	BalancerExchange                   = "Balancer"
	BalancerV2Exchange                 = "BalancerV2"
	BalancerV2ExchangeArbitrum         = "BalancerV2-Arbitrum"
	BalancerV2ExchangePolygon          = "BalancerV2-Polygon"
	BalancerV3Exchange                 = "BalancerV3"
	BeetsExchange                      = "Beets"
	KrakenExchange                     = "Kraken"
	BitfinexExchange                   = "Bitfinex"
	BitforexExchange                   = "Bitforex"
	BinanceExchange                    = "Binance"
	Binance2Exchange                   = "Binance2"
	Binance3Exchange                   = "Binance3"
	BinanceExchangeUS                  = "BinanceUS"
	BitstampExchange                   = "Bitstamp"
	CryptoDotComExchange               = "Crypto.com"
	FTXExchange                        = "FTX"
	Opyn                               = "OPYN"
	Premia                             = "Premia"
	BitBayExchange                     = "BitBay"
	CoinBaseExchange                   = "CoinBase"
	HitBTCExchange                     = "HitBTC"
	HuckleberryExchange                = "Huckleberry"
	TraderJoeExchange                  = "TraderJoe"
	TraderJoeExchangeV2_1              = "TraderJoeV2.1"
	TraderJoeExchangeV2_1Arbitrum      = "TraderJoeV2.1-Arbitrum"
	TraderJoeExchangeV2_1Avalanche     = "TraderJoeV2.1-Avalanche"
	TraderJoeExchangeV2_1BNB           = "TraderJoeV2.1-BNB"
	TraderJoeExchangeV2_2Avalanche     = "TraderJoeV2.2-Avalanche"
	PangolinExchange                   = "Pangolin"
	PlatypusExchange                   = "PlatypusFinance"
	SimexExchange                      = "Simex"
	OKExExchange                       = "OKEx"
	HuobiExchange                      = "Huobi"
	LBankExchange                      = "LBank"
	GateIOExchange                     = "GateIO"
	ZBExchange                         = "ZB"
	QuoineExchange                     = "Quoine"
	UnknownExchange                    = "Unknown"
	BlockSizeSeconds                   = 120
	FilterKing                         = "MAIR120"
	BancorExchange                     = "Bancor"
	UniswapExchange                    = "Uniswap"
	UniswapExchangeBase                = "Uniswap-Base"
	UniswapExchangeV3                  = "UniswapV3"
	UniswapExchangeV3Polygon           = "UniswapV3-polygon"
	UniswapExchangeV3Arbitrum          = "UniswapV3-Arbitrum"
	UniswapExchangeV3Base              = "UniswapV3-Base"
	UniswapExchangeV3Celo              = "UniswapV3-Celo"
	UniswapExchangeV4                  = "UniswapV4"
	LoopringExchange                   = "Loopring"
	ShadowV2Exchange                   = "ShadowV2"
	ShadowV3Exchange                   = "ShadowV3"
	CamelotExchange                    = "Camelot"
	CamelotExchangeV3                  = "CamelotV3"
	CurveFIExchange                    = "Curvefi"
	CurveFIExchangeFantom              = "Curvefi-Fantom"
	CurveFIExchangeMoonbeam            = "Curvefi-Moonbeam"
	CurveFIExchangePolygon             = "Curvefi-Polygon"
	CurveFIExchangeArbitrum            = "Curvefi-Arbitrum"
	MakerExchange                      = "Maker"
	KuCoinExchange                     = "KuCoin"
	SushiSwapExchange                  = "SushiSwap"
	SushiSwapExchangeArbitrum          = "SushiSwap-arbitrum"
	SushiSwapExchangePolygon           = "SushiSwap-polygon"
	SushiSwapExchangeFantom            = "SushiSwap-fantom"
	PanCakeSwap                        = "PanCakeSwap"
	PanCakeSwapExchangeV3              = "PanCakeSwapV3"
	ApeswapExchange                    = "Apeswap"
	BiswapExchange                     = "Biswap"
	DforceExchange                     = "Dforce"
	ZeroxExchange                      = "0x"
	KyberExchange                      = "Kyber"
	BitMartExchange                    = "BitMart"
	BitMaxExchange                     = "Bitmax"
	MEXCExchange                       = "MEXC"
	BKEXExchange                       = "BKEX"
	BKEX2Exchange                      = "BKEX2"
	CREX24Exchange                     = "CREX24"
	STEXExchange                       = "STEX"
	Deribit                            = "Deribit"
	DfynNetwork                        = "DFYN"
	UbeswapExchange                    = "Ubeswap"
	SpookyswapExchange                 = "Spookyswap"
	SpiritswapExchange                 = "Spiritswap"
	QuickswapExchange                  = "Quickswap"
	SerumExchange                      = "Serum"
	OrcaExchange                       = "Orca"
	SolarbeamExchange                  = "Solarbeam"
	TrisolarisExchange                 = "Trisolaris"
	ByBitExchange                      = "ByBit"
	BitMexExchange                     = "BitMex"
	MultiChain                         = "MultiChain"
	StellaswapExchange                 = "Stellaswap"
	WanswapExchange                    = "Wanswap"
	OsmosisExchange                    = "Osmosis"
	ZenlinkswapExchange                = "Zenlink"
	ZenlinkswapExchangeBifrostPolkadot = "Zenlink-bifrost-polkadot"
	VelodromeExchange                  = "Velodrome"
	VelodromeSlipstreamExchange        = "VelodromeSlipstream"
	MaverickExchange                   = "Maverick"
	MaverickExchangeZKSync             = "Maverick-zksync"
	MaverickExchangeBNB                = "Maverick-bnb"
	PearlfiExchangeTestnet             = "Pearlfi-Testnet"
	PearlfiExchange                    = "Pearlfi"
	PearlfiStableswapExchange          = "PearlfiStableswap"
	RamsesV1Exchange                   = "RamsesV1"
	RamsesV2Exchange                   = "RamsesV2"
	NileV1Exchange                     = "NileV1"
	NileV2Exchange                     = "NileV2"
	ThenaExchange                      = "Thena"
	ThenaV3Exchange                    = "ThenaV3"
	AyinExchange                       = "Ayin"
	BitflowExchange                    = "Bitflow"
	BitgetExchange                     = "Bitget"
	VelarExchange                      = "Velar"
	AerodromeSlipstreamExchange        = "AerodromeSlipstream"
	AerodromeV1Exchange                = "AerodromeV1"
	BifrostExchange                    = "Bifrost"
	HydrationExchange                  = "Hydration"
	VelodromeExchangeSwellchain        = "Velodrome-Swellchain"
	CoinExExchange                     = "CoinEx"
	// FinageForex        = "FinageForex"
)

const (
	SCRAPER_TYPE_ASSETCOLLECTOR   = "assetcollector"
	SCRAPER_TYPE_LIQUIDITYSCRAPER = "liquidityscraper"
	INDEX_TYPE_BLOCKNUMBER        = "blocknumber"
	INDEX_TYPE_INDEX              = "index"
)

type ConfigApi struct {
	ApiKey    string
	SecretKey string
}

type ConfigConnector struct {
	Coins []ExchangePair
}

func GetConfig(exchange string) (*ConfigApi, error) {
	if utils.Getenv("USE_ENV", "false") == "true" {
		return GetConfigFromEnv(exchange)
	}
	var configApi ConfigApi
	usr, _ := user.Current()
	dir := usr.HomeDir
	configFileApi := dir + "/config/secrets/api_" + strings.ToLower(exchange)
	err := gonfig.GetConf(configFileApi, &configApi)
	return &configApi, err
}

func GetConfigFromEnv(exchange string) (*ConfigApi, error) {
	if utils.Getenv("USE_ENV", "false") != "true" {
		return nil, errors.New("use of config by env without env activation ")
	}

	configApi := ConfigApi{
		ApiKey:    utils.Getenv("API_"+strings.ToUpper(exchange)+"_APIKEY", ""),
		SecretKey: utils.Getenv("API_"+strings.ToUpper(exchange)+"_SECRETKEY", ""),
	}
	return &configApi, nil
}
