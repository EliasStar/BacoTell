import { Command, Interaction } from "../command.ts"

const bell: Command = {
    enabled: true,

    cmd: {
        name: "bell",
        description: ""
    },

    execute(interaction: Interaction) {
    }
}

export default bell