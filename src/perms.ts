import { path } from "./deps.ts"

const envPerms = await Deno.permissions.request({ name: "env" })
if (envPerms.state === "denied") {
    console.error("no env access")
    Deno.exit()
}

const netPerms = await Deno.permissions.request({ name: "net" })
if (netPerms.state === "denied") {
    console.error("no net access")
    Deno.exit()
}

export const rootDir = path.dirname(path.fromFileUrl(import.meta.url))
const readPerms = await Deno.permissions.request({ name: "read", path: rootDir })
if (readPerms.state === "denied") {
    console.error("no read access")
    Deno.exit()
}

const token = Deno.env.get("BOT_TOKEN")
if (token == null) {
    console.error("no BOT_TOKEN env var defined")
    Deno.exit()
}

export default token