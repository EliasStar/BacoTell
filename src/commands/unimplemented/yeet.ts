import { Command, Locale, Interaction } from "/command.ts"

const yeet: Command = {
    command: (loc: Locale) => ({
        name: "issuetracker",
        description: loc.cmds.issues.desc
    }),

    handler: (loc: Locale) => (async (inter: Interaction) => {
        if (args.length < 1) {
            evt.getChannel().sendMessage(Messages.NOT_ENOUGH_ARGUMENTS).queue()
            return false
        } else if (args.length > 1) {
            evt.getChannel().sendMessage(Messages.TOO_MUCH_ARGUMENTS).queue()
            return false
        }

        Member member = null

        for (Member curMember : evt.getGuild().getMembers()) {
            if (curMember.getEffectiveName().equals(args[0]) && curMember.getOnlineStatus() == OnlineStatus.ONLINE) {
                member = curMember
                break
            }
        }

        if (member == null) {
            evt.getChannel().sendMessage(Messages.MEMBER_NOT_FOUND).queue()
            return false
        }

        String catName = Messages.COMMAND_MOVE_CATEGORY
        String pingName = Messages.COMMAND_MOVE_CHANNEL_PING
        String pongName = Messages.COMMAND_MOVE_CHANNEL_PONG

        evt.getGuild().getController().createCategory(catName).complete()
        List < Category > cat = evt.getGuild().getCategoriesByName(catName, false)

        evt.getGuild().getController().createVoiceChannel(pingName).setUserlimit(1).setParent(cat.get(0)).setBitrate(8000).complete()
        evt.getGuild().getController().createVoiceChannel(pongName).setUserlimit(1).setParent(cat.get(0)).setBitrate(8000).complete()

        List < VoiceChannel > ping = evt.getGuild().getVoiceChannelsByName(pingName, false)
        List < VoiceChannel > pong = evt.getGuild().getVoiceChannelsByName(pongName, false)

        VoiceChannel memberChannel = member.getVoiceState().getChannel()

        evt.getMessage().getChannel().sendMessage(Messages.COMMAND_MOVE).queue()

        // 4 for no timeout and 9 for 2 rounds
        for (int i = 0; i < 4; i++) {
            evt.getGuild().getController().moveVoiceMember(member, ping.get(0)).complete()
            evt.getGuild().getController().moveVoiceMember(member, pong.get(0)).complete()
        }

        evt.getGuild().getController().moveVoiceMember(member, memberChannel).complete()

        pong.get(0).delete().complete()
        ping.get(0).delete().complete()
        cat.get(0).delete().complete()

        return true

        await inter.reply({
            content: "Baco Tell Issue Tracker"
        })
    })
}

export default yeet