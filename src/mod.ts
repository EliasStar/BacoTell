import token, { rootDir } from "./perms.ts"

import { harmony, path } from "./deps.ts"
import { Command } from "./command.ts";

const client = new harmony.Client()

client.once("ready", async () => {
    console.log(`logged in as ${client.user?.tag}`)

    console.log("loading local commands")
    const localCmds = []

    for await (const dir of Deno.readDir(path.join(rootDir, "./commands/"))) {
        if (!dir.isFile) continue

        console.log(`found ${dir.name}`)
        try {
            const cmd = (await import(`./commands/${dir.name}`)).default as Command

            if (cmd.enabled) {
                localCmds.push(cmd)
                console.log(`loaded ${dir.name}`)
            } else {
                console.log(`${dir.name} is disabled`)
            }
        } catch (error) {
            console.error(`loading ${dir.name} error: ${error}`)
        }
    }

    console.log("loading remote commands")
    const remoteCmds = (await client.slash.commands.guild("620996650269278240")).array()

    console.log("syncing commands")

    for (const cmd of remoteCmds) {
        await cmd.delete()
    }

    for (const cmd of localCmds) {
        const c = await client.slash.commands.create(cmd.cmd, "620996650269278240")
        c.handle(cmd.execute)
    }
})

client.connect(token, harmony.Intents.None)