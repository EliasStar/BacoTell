import { Locale } from "/locale.ts"

const de: Locale = {
    lang: {
        en: "Englisch",
        de: "Deutsch",
        es: "Spanisch"
    },
    cmds: {
        ping: {
            desc: "Ruft den Ping zwischen dem BOT-Client und Discord Gateway ab.",
            reply: {
                results: "Ergebnisse",
                latency: "Latenz"
            }
        },
        shoot: {
            cmd: "peng",
            desc: "Ping-Befehl mit Tippfehler",
            reply: "Noch einer beißt ins Gras!"
        },
        lang: {
            cmd: "sprache",
            desc: "Ruft die Sprache des BOT ab oder ändert sie.",
            reply: {
                get: "Ich spreche derzeit",
                set: "Von jetzt an spreche ich Deutsch."
            },
            option: {
                name: "sprache",
                desc: "Die zu verwendende Sprache"
            }
        },
        issues: {
            desc: "Erstelle Issues für Baco Tell auf GitHub.",
            button: {
                bug: "Fehler melden",
                feature: "Feature vorschlagen",
                question: "Frage stelle",
                other: "Anderen Issue erstellen"
            }
        }
    }
}

export default de