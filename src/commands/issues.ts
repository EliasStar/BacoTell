import { Command, Locale, Interaction, Component } from "/command.ts"

const issues: Command = {
    command: (loc: Locale) => ({
        name: "issuetracker",
        description: loc.cmds.issues.desc
    }),

    handler: (loc: Locale) => (async (inter: Interaction) => {
        const component: Component = {
            type: 1,
            components: [
                {
                    type: 2,
                    style: 5,
                    label: loc.cmds.issues.buttons.bug,
                    url: "https://github.com/EliasStar/BacoTell/issues/new?labels=bug%2C+new&template=bug_report.md&title=%5BBug%5D+%3Ctitle%3E"
                },
                {
                    type: 2,
                    style: 5,
                    label: loc.cmds.issues.buttons.feature,
                    url: "https://github.com/EliasStar/BacoTell/issues/new?labels=feature%2C+new&template=feature_request.md&title=%5BFeature%5D+%3Ctitle%3E"
                },
                {
                    type: 2,
                    style: 5,
                    label: loc.cmds.issues.buttons.question,
                    url: "https://github.com/EliasStar/BacoTell/issues/new?labels=question%2C+new&template=question.md&title=%5BQuestion%5D+%3Ctitle%3E"
                },
                {
                    type: 2,
                    style: 5,
                    label: loc.cmds.issues.buttons.other,
                    url: "https://github.com/EliasStar/BacoTell/issues/new?labels=new"
                }
            ]
        }

        await inter.reply({
            content: "Baco Tell Issue Tracker",
            ephemeral: true,
            components: [component]
        })
    })
}

export default issues