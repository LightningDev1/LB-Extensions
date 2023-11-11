package lightningbot

// Version is the current version of the bot
const Version = ""

// ClearConsole clears the console. Works for the native, web and UI console. Example usage:
// 	lightningbot.ClearConsole()
func ClearConsole() {}

// Notification sends a toast notification to the user. Example usage:
// 	lightningbot.Notification("My notification")
func Notification(message string) {}

// LogError logs an error to the console. Example usage:
// 	lightningbot.LogError("My error")
func LogError(message string) {}

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
// 	lightningbot.NewConsoleMessage().
// 		SetTitle("My Title").
// 		SetColor("#FF0000").
// 		SetWebhookURL("https://discord.com/api/webhooks/...").
// 		AddInfo("Key", "Value").
// 		AddInfoConditional("Key", value, value != "").
// 		Show()
func NewConsoleMessage() ConsoleMessage {
	return nil
}
