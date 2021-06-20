export default {
    lang: {
        en: "English",
        de: "German",
        es: "Spanish"
    },
    cmds: {
        ping: {
            desc: "Gets the ping between BOT client and Discord Gateway.",
            reply: {
                results: "Results",
                latency: "Latency"
            }
        },
        shoot: {
            cmd: "bang",
            desc: "The sound of a gunshot",
            reply: "Another one bites the dust!"
        },
        lang: {
            cmd: "language",
            desc: "Gets or changes the language the BOT uses.",
            reply: {
                get: "I'm currently speaking",
                set: "I'll now answer in English."
            },
            option: {
                name: "locale",
                desc: "The language to change to"
            }
        },
        issue: {
            cmd: "issue",
            desc: "Report a bug or request a feature.",
        }
    }
}