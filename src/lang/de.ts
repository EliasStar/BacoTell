import { Locale } from "/locale.ts"

const de: Locale = {
    lang: {
        en: "Englisch",
        de: "Deutsch",
        es: "Spanisch"
    },
    cmds: {
        ping: {
            desc: "Ruft den Ping zwischen dem BOT-Client und Discord Gateway ab."
        },
        shoot: {
            cmd: "peng",
            desc: "Ping-Befehl mit Tippfehler",
            reply: "Noch einer beißt ins Gras!"
        },
        lang: {
            cmd: "sprache",
            desc: "Ruft die Sprache des BOT ab oder ändert sie.",
            option: {
                name: "sprache",
                desc: "Die zu verwendende Sprache"
            }
        }
    }
}

export default de