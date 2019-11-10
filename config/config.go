package config

type Channel struct {
	From      int64
	To        int64
	ChannelId string
}

type Ethereum struct {
	OwnerPrivKey string
	ContractAddr string
	RpcUrl       string
}

var (
	DiscordAPIToken            string // discord api token
	DiscordGameNotifyChannelId string // notify channel id
	DiscordLevelChannel        []Channel
	Skale                      Ethereum
)
