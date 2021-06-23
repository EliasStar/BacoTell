import { Command, Locale, Interaction } from "/command.ts"
import { ChannelTypes, Member, VoiceChannel } from "harmony"

const yeet: Command = {
    command: (loc: Locale) => ({
        name: "yeet",
        description: loc.cmds.yeet.desc,
        options: [
            {
                name: loc.cmds.yeet.options.member.name,
                description: loc.cmds.yeet.options.member.desc,
                type: "USER",
                required: true
            },
            {
                name: loc.cmds.yeet.options.amount.name,
                description: loc.cmds.yeet.options.amount.desc,
                type: "INTEGER"
            }
        ]
    }),

    handler: (loc: Locale) => (async (inter: Interaction) => {
        const member = inter.option<{ member: Member }>(loc.cmds.yeet.options.member.name).member

        if (member == null || inter.guild == null) return
        await inter.defer()

        const channel = (await inter.guild.voiceStates.get(member.id))?.channel
        if (channel == null) {
            await inter.reply(member.displayName + " " + loc.cmds.yeet.replies.offline)
            return
        }

        const yeet = (await inter.guild.createChannel({ name: "YEET", type: ChannelTypes.GUILD_VOICE })) as VoiceChannel
        const amount = inter.option<number>(loc.cmds.yeet.options.amount.name) ?? 10

        await inter.reply(member.displayName + " " + loc.cmds.yeet.replies.yeet)
        for (let i = 0; i < amount; i++) {
            await member.moveVoiceChannel(yeet)
            await member.moveVoiceChannel(channel)
        }

        await inter.guild.channels.delete(yeet.id)
    })
}

export default yeet