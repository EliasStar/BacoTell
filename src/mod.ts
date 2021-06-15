import token, { rootDir } from "./perms.ts"

import { harmony, path } from "./deps.ts"
import { areIdentical, Command } from "./command.ts";

const client = new harmony.Client()

client.once("ready", async () => {
    console.log(`logged in as ${client.user?.tag}`)

    console.log("loading local commands")
    const localCommands = []

    for await (const dir of Deno.readDir(path.join(rootDir, "./commands/"))) {
        if (!dir.isFile) continue

        console.log(`found ${dir.name}`)
        try {
            const cmd = (await import(`./commands/${dir.name}`)).default as Command

            if (cmd.enabled) {
                localCommands.push(cmd)
                console.log(`loaded ${dir.name}`)
            } else {
                console.log(`${dir.name} is disabled`)
            }
        } catch (error) {
            console.error(`loading ${dir.name} error: ${error}`)
        }
    }

    console.log("loading remote commands")
    const remoteCommands = (await client.slash.commands.guild("620996650269278240")).array()

    console.log("syncing commands")
    for (const localCmd of localCommands) {
        const index = remoteCommands.findIndex(c => c.name === localCmd.cmd.name)

        if (index === -1) {
            console.log(`registering ${localCmd.cmd.name}`)
            const cmd = await client.slash.commands.create(localCmd.cmd, "620996650269278240")

            console.log(`registering ${cmd.name} handler`)
            cmd.handle(localCmd.execute)
        } else {
            const cmd = remoteCommands.splice(index, 1)[0]
            if (!areIdentical(cmd, localCmd.cmd)) {
                console.log(`updating ${localCmd.cmd.name}`)
                await cmd.edit(localCmd.cmd)
            }

            console.log(`registering ${cmd.name} handler`)
            cmd.handle(localCmd.execute)
        }
    }

    for (const cmd of remoteCommands) {
        console.log(`deregistering ${cmd.name}`)
        await cmd.delete()
    }
})

client.connect(token, harmony.Intents.None)