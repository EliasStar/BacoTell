package proxy

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
)

// The default implementation for bacotell_common.SubmitProxy.
type submitProxy struct {
	interactionProxy
}

// submitProxy implements bacotell_common.SubmitProxy.
var _ common.SubmitProxy = submitProxy{}

// NewSubmitProxy returns a new SubmitProxy implementation.
func NewSubmitProxy(session *discordgo.Session, interaction *discordgo.Interaction) common.SubmitProxy {
	return submitProxy{
		interactionProxy: interactionProxy{
			session:     session,
			interaction: interaction,
		},
	}
}
