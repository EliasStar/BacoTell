import type { Command, CommandLocale, CommandInteraction } from "../command.ts"

const peng: Command = {
    command: (loc: CommandLocale) => ({
        name: "peng",
        description: loc.cmds.peng.desc
    }),

    handler: (loc: CommandLocale) => ((inter: CommandInteraction) => {
        inter.reply(loc.cmds.peng.reply)
    })
}

export default peng