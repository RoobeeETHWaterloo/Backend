package main

import (
	"../../config"
	"../../pkg/brawlContract"
	"../../pkg/discord"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis"
	"math/big"
	"os"
	"os/signal"
	"time"
)

var (
	bot         *discord.Bot
	redisClient *redis.Client
	err         error
)

func init() {
	bot, err = discord.NewDiscord(config.DiscordAPIToken)
	if err != nil {
		panic(err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

type GameStat struct {
	Player1WinRate float64
	Player2WinRate float64
	MovesCount     int64
	Winner         common.Address
	Player1Addr    common.Address
	Player2Addr    common.Address
}

func getStat(fightId *big.Int, caller *brawlContract.CryptoBrawl) (*GameStat, error) {
	data, err := caller.Fights(&bind.CallOpts{}, fightId)
	if err != nil {
		return nil, err
	}

	data1, err := caller.Chars(&bind.CallOpts{}, data.Player1CharID)
	if err != nil {
		return nil, err
	}

	stat := &GameStat{}
	stat.Player1WinRate = float64(data1.WinsCount.Int64()) / float64(data1.FightsCount.Int64())

	data1, err = caller.Chars(&bind.CallOpts{}, data.Player2CharID)
	if err != nil {
		return nil, err
	}
	stat.Player2WinRate = float64(data1.WinsCount.Int64()) / float64(data1.FightsCount.Int64())
	stat.MovesCount = data.StepNum.Int64()
	stat.Player1Addr = data.Player1GeneralAddress
	stat.Player2Addr = data.Player2GeneralAddress
	return stat, nil
}

func discordIdGetByAddress(address common.Address) string {
	cmd := redisClient.Get(address.String())

	if cmd.Val() == "" {
		return address.String()
	} else {
		return "<@" + cmd.Val() + ">"
	}
}

func handleSkale() {
	c, err := ethclient.Dial(config.Skale.RpcUrl)
	if err != nil {
		fmt.Println("dial", err)
		return
	}

	fmt.Println(`Started handling Skale Events`)
	addr := common.HexToAddress(config.Skale.ContractAddr)
	contract, err := brawlContract.NewCryptoBrawl(addr, c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lastFightCount := int64(0)
	zeroFight := "0x0000000000000000000000000000000000000000"
	openGames := []int64{}

	check := func() {
		fightCount, err := contract.FightsCount(nil)
		if err != nil {
			fmt.Println(`Get Fight Count error`, err)
			return
		}

		if lastFightCount == 0 {
			// пропускаем уже созданные, либо в редисе проверяем отправленные и сверяем
			// надо записывать в редис изменения
			lastFightCount = fightCount.Int64() - 1
		}

		if lastFightCount != fightCount.Int64() {
			for fightId := lastFightCount + 1; fightId <= fightCount.Int64(); fightId++ {
				stat, err := getStat(big.NewInt(fightId), contract)
				if err != nil {
					fmt.Println("err", err, "fightId", fightId)
					continue
				}

				if stat.Winner.String() == zeroFight {
					// game started
					notifyGameStart(fightId, discordIdGetByAddress(stat.Player1Addr), stat.Player1WinRate, discordIdGetByAddress(stat.Player2Addr), stat.Player2WinRate)
					openGames = append(openGames, fightId)
				} else {
					// game end
					notifyGameEnd(fightId, discordIdGetByAddress(stat.Player1Addr), stat.Player1WinRate, discordIdGetByAddress(stat.Player2Addr), discordIdGetByAddress(stat.Winner), stat.Player2WinRate, stat.MovesCount)
				}
			}
			// появились новые игры
			lastFightCount = fightCount.Int64()
		} else {
			for _, fightId := range openGames {
				stat, err := getStat(big.NewInt(fightId), contract)
				if err != nil {
					fmt.Println("err", err, "fightId", fightId)
					continue
				}
				fmt.Println("checking open game", fightId)

				if stat.Winner.String() != zeroFight {
					// game end (mean status changed)
					notifyGameEnd(fightId, discordIdGetByAddress(stat.Player1Addr), stat.Player1WinRate, discordIdGetByAddress(stat.Player2Addr), discordIdGetByAddress(stat.Winner), stat.Player2WinRate, stat.MovesCount)
				}
			}
		}
	}
	for {
		check()
		time.Sleep(time.Second * 10)
	}
}

func main() {
	err = bot.Session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	go notifierHandler()
	go handleSkale()

	<-interrupt
	fmt.Println(`Interrupting app`)
	bot.Session.Close()
}

func userOpenChannel(channelId string, userId string) {
	err := bot.Session.ChannelPermissionSet(channelId, userId, "member", 3072, 0)
	if err != nil {
		fmt.Println("cannot open permission for channel ", err)
	}
}

func userCloseChannel(channelId string, userId string) {
	err := bot.Session.ChannelPermissionSet(channelId, userId, "member", 0, 0)
	if err != nil {
		fmt.Println("cannot open permission for channel ", err)
	}
}

// обработчик оповещений
func notifierHandler() {
	// отправляем пользователю, что ему нужно сходить
	/*bot.Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		fmt.Println(m.Author.ID, m.ChannelID, "mess", m.Content)
	})
	*/
}

// start closed notify
func notifyGameStart(fightId int64, playerLeft string, playerLeftWr float64, playerRight string, playerRightWr float64) {
	msg, err := bot.Session.ChannelMessageSend(config.DiscordGameNotifyChannelId, fmt.Sprintf(`Match №%d Started, %s WR %f against %s WR %f!`, fightId, playerLeft, playerLeftWr, playerRight, playerRightWr))
	if err != nil {
		fmt.Printf("Cannot send msg, err %v\n", err)
	} else {
		fmt.Printf("Message id %s, succesfully sent\n", msg.ID)
	}
}

// end notify
func notifyGameEnd(fightId int64, playerLeft string, playerLeftWr float64, playerRight, winner string, playerRightWr float64, movesCount int64) {
	var plStatus, prStatus = "winner", "looser"
	if playerLeft != winner {
		plStatus, prStatus = "looser", "winner"
	}
	msg, err := bot.Session.ChannelMessageSend(config.DiscordGameNotifyChannelId, fmt.Sprintf(`Match №%d End, %s(%s) WR %f against %s(%s) WR %f! Moves count: %d`, fightId, playerLeft, plStatus, playerLeftWr, playerRight, prStatus, playerRightWr, movesCount))
	if err != nil {
		fmt.Printf("Cannot send msg, err %v\n", err)
	} else {
		fmt.Printf("Message id %s, succesfully sent\n", msg.ID)
	}
}
