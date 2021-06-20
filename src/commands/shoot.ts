import { Command, Locale, Interaction } from "/command.ts"

const shoot: Command = {
    command: (loc: Locale) => ({
        name: loc.cmds.shoot.cmd,
        description: loc.cmds.shoot.desc
    }),

    handler: (loc: Locale) => (async (inter: Interaction) => {
        await inter.reply(loc.cmds.shoot.reply)
    })
}

export default shoot