import { Locale } from "/locale.ts"

const es: Locale = {
    lang: {
        en: "Inglés",
        de: "Alemán",
        es: "Español"
    },
    cmds: {
        ping: {
            desc: "Obtén el ping entre el cliente BOT y Discord Gateway."
        },
        shoot: {
            cmd: "bang",
            desc: "El sonido de un disparo",
            reply: "¡Otro muerde el polvo!"
        },
        lang: {
            cmd: "idioma",
            desc: "Obtiene o cambia el idioma del BOT.",
            replies: {
                get: "Actualmente estoy hablando",
                set: "A partir de ahora responderé en español."
            },
            option: {
                name: "lengua",
                desc: "El idioma a usar"
            }
        }
    }
}

export default es