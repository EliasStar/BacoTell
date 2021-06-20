import { Command, Locale, Interaction } from "/command.ts"
import { Embed, InteractionMessageOptions, MessageComponentData } from "harmony"

const issue: Command = {
    command: (loc: Locale) => ({
        name: loc.cmds.issue.cmd,
        description: loc.cmds.issue.desc
    }),

    handler: (loc: Locale) => (async (inter: Interaction) => {
        const embed = new Embed()
            .setColor("random")
            .setAuthor("Elias*", "https://avatars.githubusercontent.com/u/31409841")
            .setTitle("Baco Tell Issue Tracker")
            .setFooter("GitHub", "https://avatars0.githubusercontent.com/u/9919");

        const component: MessageComponentData = {
            type: 1,
            components: [
                { type: 2, style: 5, label: "Report a bug", url: "https://github.com/EliasStar/BacoTell/issues/new" },
                { type: 2, style: 5, label: "Request a feature", url: "https://github.com/EliasStar/BacoTell/issues/new" }
            ]
        }


        await inter.reply({
            content: "This is a message with components",
            components: [component],
            embeds: [embed]
        })
    })
}

export default issue