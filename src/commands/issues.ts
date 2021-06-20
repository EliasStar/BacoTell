import { Command, Locale, Interaction } from "/command.ts"
import { MessageComponentData } from "harmony"

const issues: Command = {
    command: (loc: Locale) => ({
        name: "issuetracker",
        description: loc.cmds.issues.desc
    }),

    handler: (loc: Locale) => (async (inter: Interaction) => {
        const component: MessageComponentData = {
            type: 1,
            components: [
                {
                    type: 2,
                    style: 5,
                    label: loc.cmds.issues.button.bug,
                    url: "https://github.com/EliasStar/BacoTell/issues/new?labels=bug%2C+new&template=bug_report.md&title=%5BBug%5D+%3Ctitle%3E"
                },
                {
                    type: 2,
                    style: 5,
                    label: loc.cmds.issues.button.feature,
                    url: "https://github.com/EliasStar/BacoTell/issues/new?labels=feature%2C+new&template=feature_request.md&title=%5BFeature%5D+%3Ctitle%3E"
                },
                {
                    type: 2,
                    style: 5,
                    label: loc.cmds.issues.button.question,
                    url: "https://github.com/EliasStar/BacoTell/issues/new?labels=question%2C+new&template=question.md&title=%5BQuestion%5D+%3Ctitle%3E"
                },
                {
                    type: 2,
                    style: 5,
                    label: loc.cmds.issues.button.other,
                    url: "https://github.com/EliasStar/BacoTell/issues/new?labels=new"
                }
            ]
        }

        await inter.reply({
            content: "Baco Tell Issue Tracker",
            components: [component]
        })
    })
}

export default issues