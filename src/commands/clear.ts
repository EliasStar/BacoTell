import { GuildTextChannel } from "https://raw.githubusercontent.com/EliasStar/Harmony/main/mod.ts"
import { Command, Locale, Interaction } from "/command.ts"

const clear: Command = {
    command: (loc: Locale) => ({
        name: loc.cmds.clear.cmd,
        description: loc.cmds.clear.desc,
        options: [
            {
                name: loc.cmds.clear.option.name,
                description: loc.cmds.clear.option.desc,
                type: "INTEGER",
                required: true
            }
        ]
    }),

    handler: (loc: Locale) => (async (inter: Interaction) => {
        if (!(inter.channel instanceof GuildTextChannel)) return
        if (inter.options.length !== 1) return

        const count = inter.options[0].value as number
        await inter.channel.bulkDelete(count)

        await inter.reply({
            content: loc.cmds.clear.reply,
            ephemeral: true
        })
    })
}

export default clear