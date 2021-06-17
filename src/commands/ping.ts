import { Command, Locale, Interaction } from "/command.ts"

const ping: Command = {
    command: (loc: Locale) => ({
        name: "ping",
        description: loc.cmds.ping.desc
    }),

    handler: () => ((inter: Interaction) => {
        inter.reply("Pong")
    })
}

export default ping