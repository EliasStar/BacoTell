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
            replies: {
                get: "I'm currently speaking",
                set: "I'll now answer in English."
            },
            option: {
                name: "locale",
                desc: "The language to change to"
            }
        },
        issues: {
            desc: "Creates Issues for Baco Tell on GitHub.",
            buttons: {
                bug: "Report a bug",
                feature: "Request a feature",
                question: "Ask a question",
                other: "Create other issue"
            }
        },
        clear: {
            cmd: "delete",
            desc: "Deletes multiple messages in the channel.",
            reply: "Done!",
            option: {
                name: "amount",
                desc: "How many messages to be deleted"
            }
        },
        yeet: {
            desc: "Repeatedly moves a member between two channels in quick succession.",
            replies: {
                offline: "is not connected!",
                yeet: "glhf!"
            },
            options: {
                member: {
                    name: "member",
                    desc: "The person who gets yeeted"
                },
                amount: {
                    name: "amount",
                    desc: "How often the person gets yeeted"
                }
            }
        }
    }
}