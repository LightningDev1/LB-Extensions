package lightningbot

import (
	"github.com/LightningDev1/dgc"
	"github.com/LightningDev1/discordgo"
)

// Version is the current version of the bot
const Version = ""

// ClearConsole clears the console. Works for the native, web and UI console. Example usage:
//
//	lightningbot.ClearConsole()
func ClearConsole() {}

// Notification sends a toast notification to the user. Example usage:
//
//	lightningbot.Notification("My notification")
func Notification(message string) {}

// LogError logs an error to the console. Example usage:
//
//	lightningbot.LogError("My error")
func LogError(message string) {}

// Config is the configuration of LightningBot.
type Config struct {
	Account struct {
		LightningBot struct {
			Username string
			Password string
		}
		Discord struct {
			Token string
		}
	}
	Notifications struct {
		Enabled bool
	}
	Snipers struct {
		Blacklist       []string
		ServerBlacklist []string
		GiveawaySniper  struct {
			Sleep      int
			Enabled    bool
			CustomBots []string
		}
		NitroSniper      bool
		PrivnoteSniper   bool
		MultiTokenSniper bool
		Tokens           []string
	}
	Webhooks struct {
		NitroSniper      string
		GiveawaySniper   string
		SelfbotDetection string
		StaffDetection   string
		Notifiers        string
		PingDetection    string
		Misc             string
		Sessions         string
	}
	Events struct {
		BanNotifier           bool
		DMNotifier            bool
		DMEditNotifier        bool
		DMDeleteNotifier      bool
		DMTypingNotifier      bool
		MessageDeleteNotifier bool
		DiscordStaffDetection bool
		GhostPingDetection    bool
		SelfbotDetection      bool
		TicketNotifier        bool
		SessionConnected      bool
		SessionDisconnected   bool
	}
	RPC struct {
		Enabled    bool
		ClientID   string
		State      string
		Details    string
		Time       bool
		LargeImage string
		LargeText  string
		SmallImage string
		SmallText  string
		Invitable  bool
		Buttons    []struct {
			Label string
			URL   string
		}
	}
	Activity struct {
		Enabled        bool
		ActivityType   string
		Type           string
		ClientID       string
		Name           string
		State          string
		Details        string
		Invitable      bool
		Time           bool
		StreamingURL   string
		LargeImage     string
		LargeImageText string
		SmallImage     string
		SmallImageText string
		Button1Label   string
		Button1URL     string
		Button2Label   string
		Button2URL     string
	}
	UI struct {
		Theme               string
		HideToSystemTray    bool
		RunOnStartup        bool
		RunSelfbotOnStartup bool
	}
	Selfbot struct {
		CommandPrefix        string
		CommandMode          string
		Title                string
		Footer               string
		DeleteResponses      bool
		DeleteResponsesDelay int
		DeleteCommands       bool
		MobileFriendly       bool
		Client               string
	}
	Protections struct {
		AntiPhishing      bool
		AntiFriendRequest bool
		AntiMassDM        bool
		ServerProtections struct {
			AntiInvite       bool
			AntiSpam         bool
			AntiUpper        bool
			AntiDelete       bool
			ProtectedServers []string
		}
	}
	AFK struct {
		Enabled         bool
		Delay           int
		RandomDelay     bool
		OnDM            bool
		OnPing          bool
		MustBeInactive  bool
		InactiveTimeout int
		Reply           bool
		Typing          bool
		Message         string
		AI              bool
	}
	Sharing struct {
		SharedUsers            []string
		SharedCategories       []string
		SharedCommands         []string
		ShareAll               bool
		BlockDangerousCommands bool
	}
}

// Save saves the current configuration of LightningBot. Example usage:
//
//		config, err := lightningbot.GetConfig()
//
//	 err = config.Save()
func (c Config) Save() error { return nil }

// GetConfig gets the current configuration of LightningBot. Example usage:
//
//	config, err := lightningbot.GetConfig()
//
//	fmt.Println(config)
func GetConfig() (Config, error) { return Config{}, nil }

// ConsoleMessage is a utility to show information in the console.
type ConsoleMessage interface {
	// SetTitle sets the title of the message
	SetTitle(title string) ConsoleMessage

	// SetColor sets the color of the message
	// hexColor should be in the format #RRGGBB
	SetColor(hexColor string) ConsoleMessage

	// SetWebhookURL sets the webhook URL to send the message to
	SetWebhookURL(webhookURL string) ConsoleMessage

	// AddInfo adds a key/value pair to the message
	AddInfo(key, value string) ConsoleMessage

	// AddInfoConditional adds a key/value pair to the message if the condition is true
	AddInfoConditional(key, value string, condition bool) ConsoleMessage

	// Show shows the message in the console
	Show()
}

// NewConsoleMessage creates a new console message. Example usage:
//
//	lightningbot.NewConsoleMessage().
//		SetTitle("My Title").
//		SetColor("#FF0000").
//		SetWebhookURL("https://discord.com/api/webhooks/...").
//		AddInfo("Key", "Value").
//		AddInfoConditional("Key", value, value != "").
//		Show()
func NewConsoleMessage() ConsoleMessage { return nil }

// Embed is a utility to send embeds to Discord.
type Embed interface {
	// SetTitle sets the title of the embed
	SetTitle(title string) Embed

	// SetDescription sets the description of the embed
	SetDescription(description string) Embed

	// SetSubtext sets the subtext of the embed
	SetSubtext(subtext string) Embed

	// AddField adds a field to the embed
	AddField(name, value string) Embed

	// SetFooter sets the footer of the embed
	SetFooter(footer string) Embed

	// SetImage sets the image of the embed
	SetImage(imageURL string) Embed

	// SetAuthor sets the author of the embed
	SetAuthor(author string) Embed

	// Send sends the embed to Discord
	Send(ctx *dgc.Ctx) []*discordgo.Message

	// SendWebhook sends the embed to a webhook
	SendWebhook(webhookURL string)
}

// NewEmbed creates a new embed. This is the main way LightningBot sends messages in Discord. Example usage:
//
//	lightningbot.NewEmbed().
//		SetTitle("My Title").
//		SetDescription("My Description").
//		SetSubtext("My Subtext").
//		AddField("My Field", "My Value").
//		SetFooter("My Footer").
//		SetImage("https://example.com/image.png").
//		SetAuthor("My Author").
//		Send(ctx)
func NewEmbed() Embed { return nil }
