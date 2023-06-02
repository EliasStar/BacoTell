package proxy

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
)

type submitProxy struct {
	interactionProxy
}

var _ common.SubmitProxy = submitProxy{}

func NewSubmitProxy(session *discordgo.Session, interaction *discordgo.Interaction) common.SubmitProxy {
	return submitProxy{
		interactionProxy: interactionProxy{
			session:     session,
			interaction: interaction,
		},
	}
}
