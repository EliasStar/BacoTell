import token from "/perms.ts"
import { loadCommands, deployCommands } from "/command.ts"
import { Client, GatewayIntents } from "harmony"

await loadCommands()

const client = new Client()

client.once("ready", () => console.log(`logged in as "${client.user!.tag}"`))
client.on("guildCreate", guild => deployCommands(guild))
client.on("guildLoaded", guild => deployCommands(guild))

client.connect(token, [GatewayIntents.GUILDS, GatewayIntents.GUILD_VOICE_STATES])