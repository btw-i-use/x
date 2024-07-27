// Package jufra lets Mimi have conversations with users.
//
// "jufra" means "utterance" in Lojban.
package jufra

import (
	"context"
	"flag"
	"log/slog"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
	"within.website/x/cmd/mimi/internal"
	"within.website/x/web/ollama"
	"within.website/x/web/ollama/llamaguard"
	"within.website/x/web/openai/chatgpt"
)

var (
	chatChannels    = flag.String("jufra-chat-channels", "217096701771513856,1266740925137289287", "comma-separated list of channels to allow chat in")
	llamaGuardModel = flag.String("jufra-llama-guard-model", "xe/llamaguard3", "ollama model tag for llama guard")
	mimiModel       = flag.String("jufra-mimi-model", "xe/mimi:llama3.1", "ollama model tag for mimi")
)

type Module struct {
	sess   *discordgo.Session
	cli    chatgpt.Client
	ollama *ollama.Client

	convHistory map[string][]ollama.Message
	lock        sync.Mutex
}

func New(sess *discordgo.Session) *Module {
	result := &Module{
		sess:        sess,
		cli:         chatgpt.NewClient("").WithBaseURL(internal.OllamaHost()),
		ollama:      internal.OllamaClient(),
		convHistory: make(map[string][]ollama.Message),
	}

	sess.AddHandler(result.messageCreate)

	if _, err := sess.ApplicationCommandCreate("1119055490882732105", "", &discordgo.ApplicationCommand{
		Name:                     "clearconv",
		Type:                     discordgo.ChatApplicationCommand,
		Description:              "Clear the conversation history for the current channel",
		DefaultMemberPermissions: &[]int64{discordgo.PermissionSendMessages}[0],
	}); err != nil {
		slog.Error("error creating clearconv command", "err", err)
	}

	sess.AddHandler(result.clearConv)

	return result
}

func (m *Module) clearConv(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != "clearconv" {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.convHistory, i.ChannelID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "conversation history cleared",
		},
	})
}

func (m *Module) messageCreate(s *discordgo.Session, mc *discordgo.MessageCreate) {
	if mc.Author.Bot {
		return
	}

	if !strings.Contains(*chatChannels, mc.ChannelID) {
		return
	}

	if mc.Content == "" {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	conv := m.convHistory[mc.ChannelID]

	conv = append(conv, ollama.Message{
		Role:    "user",
		Content: mc.Content,
	})

	slog.Info("message count", "len", len(conv))

	lgResp, err := m.llamaGuardCheck(context.Background(), "user", conv)
	if err != nil {
		slog.Error("error checking message", "err", err, "message_id", mc.ID, "channel_id", mc.ChannelID)
		s.ChannelMessageSend(mc.ChannelID, "error checking message")
		return
	}

	if !lgResp.IsSafe {
		msg, err := m.llamaGuardComplain(context.Background(), "user", lgResp)
		if err != nil {
			slog.Error("error generating response", "err", err, "message_id", mc.ID, "channel_id", mc.ChannelID)
			s.ChannelMessageSend(mc.ChannelID, "error generating response")
			return
		}

		s.ChannelMessageSend(mc.ChannelID, msg)
		return
	}

	cr := &ollama.CompleteRequest{
		Model:    *mimiModel,
		Messages: conv,
	}

	resp, err := m.ollama.Chat(context.Background(), cr)
	if err != nil {
		slog.Error("error chatting", "err", err, "message_id", mc.ID, "channel_id", mc.ChannelID)
		s.ChannelMessageSend(mc.ChannelID, "error chatting")
		return
	}

	conv = append(conv, resp.Message)

	lgResp, err = m.llamaGuardCheck(context.Background(), "mimi", conv)
	if err != nil {
		slog.Error("error checking message", "err", err, "message_id", mc.ID, "channel_id", mc.ChannelID)
		s.ChannelMessageSend(mc.ChannelID, "error checking message")
		return
	}

	if !lgResp.IsSafe {
		slog.Error("rule violation detected", "message_id", mc.ID, "channel_id", mc.ChannelID, "categories", lgResp.ViolationCategories, "message", resp.Message.Content)
		msg, err := m.llamaGuardComplain(context.Background(), "assistant", lgResp)
		if err != nil {
			slog.Error("error generating response", "err", err, "message_id", mc.ID, "channel_id", mc.ChannelID)
			s.ChannelMessageSend(mc.ChannelID, "error generating response")
			return
		}

		s.ChannelMessageSend(mc.ChannelID, msg)
		return
	}

	s.ChannelMessageSend(mc.ChannelID, resp.Message.Content)

	m.convHistory[mc.ChannelID] = conv
}

func (m *Module) llamaGuardCheck(ctx context.Context, role string, messages []ollama.Message) (*llamaguard.Response, error) {
	return llamaguard.Check(ctx, m.ollama, role, *llamaGuardModel, messages)
}

func (m *Module) llamaGuardComplain(ctx context.Context, from string, lgResp *llamaguard.Response) (string, error) {
	var sb strings.Builder
	sb.WriteString("⚠️ Rule violation detected from ")
	sb.WriteString(from)
	sb.WriteString(":\n")
	for _, cat := range lgResp.ViolationCategories {
		sb.WriteString("- ")
		sb.WriteString(cat.String())
		sb.WriteString("\n")
	}

	return sb.String(), nil
}
