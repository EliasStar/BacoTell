import token from "./perms.ts"

import en from "./lang/en.ts"
import { harmony } from "./deps.ts"
import { loadCommands, syncCommands } from "./command.ts";

const client = new harmony.Client()

client.once("ready", async () => {
    console.log(`logged in as ${client.user?.tag}`)

    console.log("loading local commands")
    const localCommands = await loadCommands([])

    console.log("loading remote commands")
    const remoteCommands = (await client.slash.commands.guild("620996650269278240")).array()

    console.log("syncing commands")
    await syncCommands(localCommands, remoteCommands, en, "620996650269278240")
})

client.connect(token, [harmony.GatewayIntents.GUILDS])

export default client