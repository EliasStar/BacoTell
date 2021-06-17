import token from "/perms.ts"
import { loadCommands, updateCommands } from "/command.ts"
import { Client, GatewayIntents } from "harmony"

const client = new Client()

client.once("ready", async () => {
    console.log(`logged in as "${client.user!.tag}"`)

    const guild = await client.guilds.fetch("620996650269278240")

    console.log("loading commands")
    const localCommands = await loadCommands()

    console.log("updating commands")
    await updateCommands(localCommands, guild)
})

client.connect(token, [GatewayIntents.GUILDS])