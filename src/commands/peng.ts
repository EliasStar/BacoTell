import { Command, Interaction } from "../command.ts"

const peng: Command = {
    enabled: true,

    cmd: {
        name: "peng",
        description: ""
    },

    execute(interaction: Interaction) {
    }
}

export default peng