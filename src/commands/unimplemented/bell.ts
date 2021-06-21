import Command from "../typings/command";
import { Message, MessageMentions, TextChannel } from "discord.js";
import fs from "fs";
import { join } from "path";

const Bell: Command = class {
    static featureName = "bell";
    static commandNames = ["bell"];
    static args = { min: 0, max: 1 };

    static description = "bell.description";

    static checkArguments(cmd: string, args: string[]) {
        return !args[0] || MessageMentions.CHANNELS_PATTERN.test(args[0]);
    }

    static async execute(msg: Message, mention: boolean, commandName: string, args?: string[]) {
        if (msg.member?.voice.channel) {
            Bell.busy = true;
            const connection = await msg.member?.voice.channel.join();

            connection.on('disconnect', () => Bell.busy = false);

            let dispatcher = connection.play(fs.createReadStream(join(__dirname, '../media/bell.ogg')), { volume: 2, type: 'ogg/opus' });

            dispatcher.on('finish', () => connection.disconnect());

            if (args) {
                let matches = args[0].match(/^<#(\d+)>$/);

                if (!matches) return;

                let channel = msg.guild?.channels.resolve(matches[1]);

                if (channel && channel.type === "text") {
                    (<TextChannel>channel).send("Die 🔔 beendet den Unterricht!");
                }
            }
        }
    }
}

export default Bell;