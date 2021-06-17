import { dirname, fromFileUrl } from "path"

export const rootDir = dirname(fromFileUrl(import.meta.url))
export const blacklistFile = "./cmd_blacklist.json"

const readRootPerms = await Deno.permissions.request({ name: "read", path: rootDir })
const readBlacklistPerms = await Deno.permissions.request({ name: "read", path: blacklistFile })
if (readRootPerms.state === "denied" || readBlacklistPerms.state === "denied") {
    console.error("no read access")
    Deno.exit()
}

const netPerms = await Deno.permissions.request({ name: "net" })
if (netPerms.state === "denied") {
    console.error("no net access")
    Deno.exit()
}

const envPerms = await Deno.permissions.request({ name: "env" })
if (envPerms.state === "denied") {
    console.error("no env access")
    Deno.exit()
}

const token = Deno.env.get("BOT_TOKEN")
if (token == null) {
    console.error("no BOT_TOKEN env var defined")
    Deno.exit()
}

export default token