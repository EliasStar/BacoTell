import Command from "../typings/command";
import { Message } from "discord.js";
import fs from "fs";
import { join } from "path";

const Unz: Command = class {
    static featureName = "unz";
    static commandNames = ["unz", "unzunz"];
    static args = { min: 0, max: 1 };
    static timeout = 5000;
    static busy: boolean;

    static description = "unz.description";

    static checkArguments(cmd: string, args: string[]) {
        return cmd === "unzunz" && args.length === 0 || cmd === "unz" && args.length === 0 || cmd === "unz" && args[0] === "unz";
    }

    static async execute(msg: Message, mention: boolean, commandName: string, args?: string[]) {
        if (msg.member?.voice.channel) {
            Unz.busy = true;
            const connection = await msg.member?.voice.channel.join();

            connection.on('disconnect', () => Unz.busy = false);

            const file = commandName === "unzunz" || args?.length === 1 ? 'unzunz.ogg' : 'unz.ogg';

            const dispatcher = connection.play(fs.createReadStream(join(__dirname, "../media/", file)), { type: 'ogg/opus' });

            dispatcher.on('finish', () => connection.disconnect());

            msg.react("🇺");
            msg.react("🇳");
            msg.react("🇿");
        }
    }
}

export default Unz;