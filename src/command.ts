import { blacklistFile, rootDir } from "/perms.ts"
import { getLocaleFromGuild, loadLocale, Locale } from "/locale.ts"
import { join } from "path"
import { Guild, SlashCommandChoice, SlashCommandHandlerCallback, SlashCommandInteraction, SlashCommandOption, SlashCommandOptionType, SlashCommandPartial } from "harmony"

export type { Locale } from "/locale.ts"
export type { SlashCommandInteraction as Interaction } from "harmony"
export interface Command {
    command(loc: Locale): SlashCommandPartial
    handler(loc: Locale): SlashCommandHandlerCallback
}

const loadedCommands: Command[] = []

export async function loadCommands() {
    console.log("loading command blacklist")
    const blacklist = JSON.parse(await Deno.readTextFile(blacklistFile))
    const blacklistedCmds = Array.isArray(blacklist) ? blacklist as string[] : []

    console.log("loading commands")
    loadedCommands.splice(0) // Remove all commands
    for await (const file of Deno.readDir(join(rootDir, "./commands/"))) {
        if (!file.isFile) continue

        const filename = file.name.slice(0, -3)

        console.log(`found "${filename}"`)
        if (blacklistedCmds.includes(filename)) {
            console.log(`"${filename}" is disabled`)
            continue
        }

        try {
            const cmd = (await import(`/commands/${file.name}`)).default as Command
            loadedCommands.push(cmd)
            console.log(`loaded "${filename}"`)
        } catch (error) {
            console.error(`loading error: ${error}`)
        }
    }
}

export async function deployCommands(guild: Guild) {
    const deployedCommands = await guild.commands.all()
    const locale = await loadLocale(getLocaleFromGuild(guild.id))

    console.log(`deploying commands to "${guild.name}" (${guild.id})`)
    for (const localCmd of loadedCommands) {
        const cmdPartial = localCmd.command(locale)
        const cmdHandler: SlashCommandHandlerCallback = inter => {
            console.log(`command "${inter.name}" used on "${guild.name}" (${guild.id})`)
            return localCmd.handler(locale)(inter)
        }

        let remoteCmd = deployedCommands.find(c => c.name === cmdPartial.name)
        if (remoteCmd == null) {
            console.log(`registering "${cmdPartial.name}"`)
            remoteCmd = await guild.commands.create(cmdPartial)
        } else {
            deployedCommands.delete(remoteCmd.id)
            if (!compareCommands(remoteCmd, cmdPartial)) {
                console.log(`updating "${cmdPartial.name}"`)
                await remoteCmd.edit(cmdPartial)
            }
        }

        console.log(`registering "${remoteCmd.name}" handler`)
        remoteCmd.handle(cmdHandler)
    }

    for (const [_, cmd] of deployedCommands) {
        console.log(`deregistering "${cmd.name}"`)
        await cmd.delete()
    }
}

function compareCommands(cmd1: SlashCommandPartial, cmd2: SlashCommandPartial): boolean {
    let identical = cmd1.name === cmd2.name && cmd1.description === cmd2.description

    if (cmd1.defaultPermission != null && cmd2.defaultPermission != null) {
        identical &&= cmd1.defaultPermission === cmd2.defaultPermission
    } else if (cmd1.defaultPermission != null) {
        identical &&= cmd1.defaultPermission
    } else if (cmd2.defaultPermission != null) {
        identical &&= cmd2.defaultPermission
    }

    identical &&= compareOptions(cmd1.options, cmd2.options)

    return identical
}

function compareOptions(opt1?: SlashCommandOption[], opt2?: SlashCommandOption[]): boolean {
    if (opt1 != null && opt2 != null) {
        if (opt1.length !== opt2.length) return false

        for (let i = 0; i < opt1.length; i++) {
            let identical = opt1[i].name === opt2[i].name && opt1[i].description === opt2[i].description

            const type1 = typeof opt1[i].type === "string" ? SlashCommandOptionType[opt1[i].type] : opt1[i].type
            const type2 = typeof opt2[i].type === "string" ? SlashCommandOptionType[opt2[i].type] : opt2[i].type
            identical &&= type1 === type2

            if (opt1[i].required != null && opt2[i].required != null) {
                identical &&= opt1[i].required === opt2[i].required
            } else if (opt1[i].required != null) {
                identical &&= !opt1[i].required
            } else if (opt2[i].required != null) {
                identical &&= !opt2[i].required
            }

            identical &&= compareChoices(opt1[i].choices, opt2[i].choices)
            identical &&= compareOptions(opt1[i].options, opt2[i].options)

            if (!identical) return false
        }
    } else if (opt1 != null) {
        return opt1.length === 0
    } else if (opt2 != null) {
        return opt2.length === 0
    }

    return true
}

function compareChoices(cho1?: SlashCommandChoice[], cho2?: SlashCommandChoice[]): boolean {
    if (cho1 != null && cho2 != null) {
        if (cho1.length !== cho2.length) return false

        for (let i = 0; i < cho1.length; i++) {
            if (cho1[i].name !== cho2[i].name || cho1[i].value !== cho2[i].value) return false
        }
    } else if (cho1 != null) {
        return cho1.length === 0
    } else if (cho2 != null) {
        return cho2.length === 0
    }

    return true
}