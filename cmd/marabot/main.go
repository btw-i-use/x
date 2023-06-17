package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "modernc.org/sqlite"
	"tailscale.com/hostinfo"
	"within.website/ln"
	"within.website/ln/opname"
	"within.website/x/internal"
	"within.website/x/web/revolt"
)

var (
	dbFile                = flag.String("db-file", "marabot.db", "Path to the database file")
	discordToken          = flag.String("discord-token", "", "Discord bot token")
	revoltToken           = flag.String("revolt-token", "", "Revolt bot token")
	revoltAPIServer       = flag.String("revolt-api-server", "https://api.revolt.chat", "API server for Revolt")
	revoltWebsocketServer = flag.String("revolt-ws-server", "wss://ws.revolt.chat", "Websocket server for Revolt")
	tsAuthkey             = flag.String("ts-authkey", "", "Tailscale authkey")
	tsHostname            = flag.String("ts-hostname", "", "Tailscale hostname")

	//go:embed schema.sql
	dbSchema string
)

const (
	furryholeDiscord = "192289762302754817"
	furryholeRevolt  = "01H2VRKJFPYPEAE438B6JRFSCP"
)

func main() {
	internal.HandleStartup()

	hostinfo.SetApp("within.website/x/cmd/marabot")

	ctx, cancel := context.WithCancel(opname.With(context.Background(), "marabot"))
	defer cancel()

	ln.Log(ctx, ln.Action("starting up"))

	db, err := sql.Open("sqlite", *dbFile)
	if err != nil {
		ln.FatalErr(ctx, err, ln.Action("opening sqlite database"))
	}
	defer db.Close()

	if _, err := db.ExecContext(ctx, dbSchema); err != nil {
		ln.FatalErr(ctx, err, ln.Action("running database schema"))
	}

	ircmsgs := make(chan string)

	go NewIRCBot(ctx, db, ircmsgs)

	// Init a new client.
	client := revolt.NewWithEndpoint(*revoltToken, *revoltAPIServer, *revoltWebsocketServer)

	mr := &MaraRevolt{
		cli: client,
		db:  db,
	}

	client.Connect(ctx, mr)

	dg, err := discordgo.New("Bot " + *discordToken)
	if err != nil {
		ln.FatalErr(ctx, err, ln.Action("creating discord client"))
	}

	dg.AddHandler(mr.DiscordMessageCreate)
	dg.AddHandler(mr.DiscordMessageEdit)

	if err := dg.Open(); err != nil {
		ln.FatalErr(ctx, err, ln.Action("opening discord client"))
	}
	defer dg.Close()

	if err := importDiscordData(ctx, db, dg); err != nil {
		ln.Error(ctx, err)
	}

	// Wait for close.
	sc := make(chan os.Signal, 1)

	signal.Notify(
		sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)
	for {
		select {
		case <-ctx.Done():
			return
		case <-sc:
			ln.Log(ctx, ln.Info("shutting down"))
			cancel()
			time.Sleep(150 * time.Millisecond)
			return
		}
	}
}

func (mr *MaraRevolt) DiscordMessageEdit(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if _, err := mr.db.Exec("UPDATE discord_messages SET content = ?, edited_at = ? WHERE id = ?", m.Content, time.Now().Format(time.RFC3339), m.ID); err != nil {
		ln.Error(context.Background(), err)
	}
}

func (mr *MaraRevolt) DiscordMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if err := mr.discordMessageCreate(s, m); err != nil {
		ln.Error(context.Background(), err)
	}
}

func (mr *MaraRevolt) discordMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if _, err := mr.db.Exec(`INSERT INTO discord_users (id, username, avatar_url, accent_color)
VALUES (?, ?, ?, ?)
ON CONFLICT(id)
DO UPDATE SET username = EXCLUDED.username, avatar_url = EXCLUDED.avatar_url, accent_color = EXCLUDED.accent_color`, m.Author.ID, m.Author.Username, m.Author.AvatarURL(""), m.Author.AccentColor); err != nil {
		return err
	}

	if _, err := mr.db.Exec(`INSERT INTO discord_messages (id, guild_id, channel_id, author_id, content, created_at, edited_at, webhook_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, m.ID, m.GuildID, m.ChannelID, m.Author.ID, m.Content, m.Timestamp.Format(time.RFC3339), m.EditedTimestamp, m.WebhookID); err != nil {
		return err
	}

	if m.WebhookID != "" {
		if _, err := mr.db.Exec("INSERT INTO discord_webhook_message_info (id, name, avatar_url) VALUES (?, ?, ?)", m.ID, m.Author.Username, m.Author.AvatarURL("")); err != nil {
			return err
		}
	}

	for _, att := range m.Attachments {
		if _, err := mr.db.Exec(`INSERT INTO discord_attachments (id, message_id, url, proxy_url, filename, content_type, width, height, size) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`, att.ID, m.ID, att.URL, att.ProxyURL, att.Filename, att.ContentType, att.Width, att.Height, att.Size); err != nil {
			return err
		}
	}

	ch, err := s.Channel(m.ChannelID)
	if err != nil {
		return err
	}

	if _, err := mr.db.Exec("INSERT INTO discord_channels (id, guild_id, name, topic, nsfw) VALUES (?, ?, ?, ?, ?) ON CONFLICT(id) DO UPDATE SET name = EXCLUDED.name, topic = EXCLUDED.topic, nsfw = EXCLUDED.nsfw", ch.ID, ch.GuildID, ch.Name, ch.Topic, ch.NSFW); err != nil {
		return err
	}

	for _, emoji := range m.GetCustomEmojis() {
		if _, err := mr.db.Exec("INSERT INTO discord_emoji (id, guild_id, name, url) VALUES (?, ?, ?, ?) ON CONFLICT(id) DO UPDATE SET name = EXCLUDED.name, url = EXCLUDED.url", emoji.ID, furryholeDiscord, emoji.Name, fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.webp?size=240&quality=lossless", emoji.ID)); err != nil {
			return err
		}
	}

	return nil
}

type MaraRevolt struct {
	cli *revolt.Client
	db  *sql.DB
	revolt.NullHandler
}

func (mr *MaraRevolt) MessageCreate(ctx context.Context, msg *revolt.Message) error {
	if msg.Content == "!ping" {
		sendMsg := &revolt.SendMessage{
			Masquerade: &revolt.Masquerade{
				Name:      "cipra",
				AvatarURL: "https://cdn.xeiaso.net/avatar/" + internal.Hash("cipra", "yasomi"),
			},
		}
		sendMsg.SetContent("🏓 Pong!")

		if _, err := mr.cli.MessageReply(ctx, msg.ChannelId, msg.ID, true, sendMsg); err != nil {
			return err
		}
	}
	return nil
}
