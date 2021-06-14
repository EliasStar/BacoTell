import { Command, Interaction } from "../command.ts"

const ping: Command = {
    enabled: true,

    cmd: {
        name: "ping",
        description: "Test Command"
    },

    execute(interaction: Interaction) {
        interaction.reply("pong/")
        interaction.send("30ms")
    }
}

export default ping