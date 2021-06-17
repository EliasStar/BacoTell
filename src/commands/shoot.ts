import { Command, Locale, Interaction } from "/command.ts"

const peng: Command = {
    command: (loc: Locale) => ({
        name: loc.cmds.shoot.cmd,
        description: loc.cmds.shoot.desc
    }),

    handler: (loc: Locale) => ((inter: Interaction) => {
        inter.reply(loc.cmds.shoot.reply)
    })
}

export default peng