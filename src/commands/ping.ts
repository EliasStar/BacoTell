import type { Command, CommandLocale, CommandInteraction } from "../command.ts"

const ping: Command = {
    command: (loc: CommandLocale) => ({
        name: "ping",
        description: loc.cmds.ping.desc
    }),

    handler: () => ((inter: CommandInteraction) => {
        inter.reply("Pong")
    })
}

export default ping