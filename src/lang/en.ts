export default {
    lang: {
        en: "English",
        de: "German",
        es: "Spanish"
    },
    cmds: {
        ping: {
            desc: "Gets the ping between BOT client and Discord Gateway."
        },
        shoot: {
            cmd: "bang",
            desc: "The sound of a gunshot",
            reply: "Another one bites the dust!"
        },
        lang: {
            cmd: "language",
            desc: "Gets or changes the language the BOT uses.",
            replies: {
                get: "I'm currently speaking",
                set: "I'll now answer in English."
            },
            option: {
                name: "locale",
                desc: "The language to change to"
            }
        }
    }
}