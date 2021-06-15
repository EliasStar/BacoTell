import { harmony } from "./deps.ts"
//import { Locale } from "./lang.ts";

export class Interaction extends harmony.SlashCommandInteraction { }
export interface Command {
    enabled: boolean
    cmd: harmony.SlashCommandPartial
    execute: harmony.SlashCommandHandlerCallback
    //setLang(local: Locale): void
}

export function areIdentical(cmd1: harmony.SlashCommandPartial, cmd2: harmony.SlashCommandPartial): boolean {
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

function compareOptions(opt1?: harmony.SlashCommandOption[], opt2?: harmony.SlashCommandOption[]): boolean {
    if (opt1 != null && opt2 != null) {
        if (opt1.length !== opt2.length) return false

        for (let i = 0; i < opt1.length; i++) {
            let identical = opt1[i].name === opt2[i].name && opt1[i].description === opt2[i].description && opt1[i].type === opt2[i].type

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

function compareChoices(cho1?: harmony.SlashCommandChoice[], cho2?: harmony.SlashCommandChoice[]): boolean {
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