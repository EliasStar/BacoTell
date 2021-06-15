import { harmony } from "./deps.ts"

export interface Command {
    enabled: boolean
    cmd: harmony.SlashCommandPartial
    execute: harmony.SlashCommandHandlerCallback
}

export class Interaction extends harmony.SlashCommandInteraction { }

export function areIdentical(cmd1: harmony.SlashCommandPartial, cmd2: harmony.SlashCommandPartial): boolean {
    let identical = cmd1.name === cmd2.name && cmd1.description === cmd2.description

    if (cmd1.defaultPermission != null && cmd2.defaultPermission != null) {
        identical &&= cmd1.defaultPermission === cmd2.defaultPermission
    } else if (cmd1.defaultPermission != null) {
        identical &&= cmd1.defaultPermission
    } else if (cmd2.defaultPermission != null) {
        identical &&= cmd2.defaultPermission
    }

    identical &&= optionsIdentical(cmd1.options, cmd2.options)

    return identical
}

function optionsIdentical(opt1?: harmony.SlashCommandOption[], opt2?: harmony.SlashCommandOption[]): boolean {
    if (opt1 != null && opt2 != null) {
        if (opt1.length !== opt2.length) return false

        let identical = true

        for (let i = 0; i < opt1.length && identical; i++) {
            identical &&= opt1[i].name === opt2[i].name && opt1[i].description === opt2[i].description && opt1[i].type === opt2[i].type

            if (opt1[i].required != null && opt2[i].required != null) {
                identical &&= opt1[i].required === opt2[i].required
            } else if (opt1[i].required != null) {
                identical &&= !opt1[i].required
            } else if (opt2[i].required != null) {
                identical &&= !opt2[i].required
            }

            // Choice

            identical &&= optionsIdentical(opt1[i].options, opt2[i].options)
        }

        return identical
    } else if (opt1 != null) {
        return opt1.length === 0
    } else if (opt2 != null) {
        return opt2.length === 0
    }

    return true
}