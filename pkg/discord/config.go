package discord

var (
	DefaultConfig = Config{
		ApplicationID:       "",
		PublicKey:           "",
		BotToken:            "",
		WebhookURL:          "",
		HokieWorldChannelID: "",
		DalhaegaGuildID:     "",
		AdminUserID:         "98434604616122368",
	}
)

type Config struct {
	ApplicationID       string `mapstructure:"application_id"`
	PublicKey           string `mapstructure:"public_key"`
	BotToken            string `mapstructure:"bot_token"`
	WebhookURL          string `mapstructure:"webhook_url"`
	HokieWorldChannelID string `mapstructure:"hokie_world_channel_id"`
	DalhaegaGuildID     string `mapstructure:"dalhaega_guild_id"`

	// UserIDs
	AdminUserID string `mapstructure:"admin_user_id"`
}

// bot invite URL
// https://discord.com/oauth2/authorize?client_id=1055796640973336641&permissions=8&scope=bot
