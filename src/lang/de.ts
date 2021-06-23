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
            replies: {
                get: "Ich spreche derzeit",
                set: "Von jetzt an spreche ich Deutsch."
            },
            option: {
                name: "sprache",
                desc: "Die zu verwendende Sprache"
            }
        },
        issues: {
            desc: "Erstellt Issues für Baco Tell auf GitHub.",
            buttons: {
                bug: "Fehler melden",
                feature: "Feature vorschlagen",
                question: "Frage stelle",
                other: "Anderen Issue erstellen"
            }
        },
        clear: {
            cmd: "lösche",
            desc: "Löscht mehrere Nachrichten im Kanal.",
            reply: "Fertig!",
            option: {
                name: "anzahl",
                desc: "Wie viele Nachrichten zu löschen sind"
            }
        },
        yeet: {
            desc: "Bewegt ein Mitglied schnell zwischen zwei Kanälen hin und her.",
            replies: {
                offline: "ist nicht verbunden!",
                yeet: "jetzt gehts los! Anschnallen, Aaabfaaahrt und Goou!"
            },
            options: {
                member: {
                    name: "nutzer",
                    desc: "Welche Person geyeeted wird"
                },
                amount: {
                    name: "anzahl",
                    desc: "Wie oft die Person geyeeted wird"
                }
            }
        }
    }
}

export default de