import { Command, Interaction } from "../command.ts"

const lang: Command = {
    enabled: true,

    cmd: {
        name: "lang",
        description: ""
    },

    execute(interaction: Interaction) {
    }
}

export default lang