import { harmony } from "./deps.ts"

export interface Command {
    enabled: boolean
    cmd: harmony.SlashCommandPartial
    execute: harmony.SlashCommandHandlerCallback
}

export class Interaction extends harmony.SlashCommandInteraction { }