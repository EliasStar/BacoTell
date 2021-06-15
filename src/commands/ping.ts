import { Command, Interaction } from "../command.ts"

const ping: Command = {
    enabled: true,

    cmd: {
        name: "ping",
        description: "Get ping between BOT client and Discord Gateway."
    },

    execute: (interaction: Interaction) => interaction.reply("Pong")
}

export default ping