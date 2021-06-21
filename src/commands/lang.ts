import { Command, deployCommands, Interaction, Locale } from "/command.ts"
import { getLocaleFromGuild, setLocaleForGuild, loadLocale, LocaleIdentifier } from "/locale.ts"

const lang: Command = {
    command: (loc: Locale) => ({
        name: loc.cmds.lang.cmd,
        description: loc.cmds.lang.desc,
        options: [
            {
                name: loc.cmds.lang.option.name,
                description: loc.cmds.lang.option.desc,
                type: "STRING",
                choices: [
                    {
                        name: loc.lang.en,
                        value: "en",
                    },
                    {
                        name: loc.lang.de,
                        value: "de",
                    },
                    {
                        name: loc.lang.es,
                        value: "es",
                    },
                ],
            },
        ],
    }),

    handler: (loc: Locale) => (async (inter: Interaction) => {
        if (inter.guild == null) return
        await inter.defer()

        const guild = inter.guild

        if (inter.options.length === 0) {
            const locale = getLocaleFromGuild(guild.id)
            await inter.reply(`${loc.cmds.lang.reply.get} ${loc.lang[locale]}.`)
        } else {
            const newLocale = inter.options[0].value as LocaleIdentifier

            setLocaleForGuild(guild.id, newLocale)
            const newLoc = await loadLocale(newLocale)

            await deployCommands(guild)
            await inter.reply(newLoc.cmds.lang.reply.set)
        }
    }),
}

export default lang
