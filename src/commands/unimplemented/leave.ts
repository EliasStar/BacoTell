import Command from "../typings/command";
import { Message } from "discord.js";

const Leave: Command = class {
    static featureName = "leave";
    static commandNames = ["leave", "stop"];
    static args = 0;
    static timeout = 5000;

    static description = "leave.description";

    static execute(msg: Message, mention: boolean, commandName: string) {
        if (msg.guild?.me?.voice.connection) {
            msg.guild?.me?.voice.connection.disconnect();
        }
    }
}

export default Leave;