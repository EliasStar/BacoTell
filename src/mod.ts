import token from "/perms.ts"
import { loadCommands, syncCommands } from "/command.ts"
import { Client, GatewayIntents } from "harmony"

const client = new Client()

client.once("ready", async () => {
    console.log(`logged in as ${client.user?.tag}`)

    console.log("loading local commands")
    const localCommands = await loadCommands([])

    console.log("loading remote commands")
    const remoteCommands = (await client.slash.commands.guild("620996650269278240")).array()

    console.log("syncing commands")
    await syncCommands("620996650269278240", localCommands, remoteCommands)
})

client.connect(token, [GatewayIntents.GUILDS])

export default client