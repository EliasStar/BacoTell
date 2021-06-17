import { Command, Locale, Interaction } from "/command.ts"

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
                        value: "en"
                    },
                    {
                        name: loc.lang.de,
                        value: "de"
                    },
                    {
                        name: loc.lang.es,
                        value: "es"
                    }
                ],


            }
        ]
    }),

    handler: (_loc: Locale) => ((_inter: Interaction) => {

    })
}

export default lang