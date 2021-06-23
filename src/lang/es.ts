import { Locale } from "/locale.ts"

const es: Locale = {
    lang: {
        en: "Inglés",
        de: "Alemán",
        es: "Español"
    },
    cmds: {
        ping: {
            desc: "Obtén el ping entre el cliente BOT y Discord Gateway.",
            reply: {
                results: "Resultados",
                latency: "Latencia"
            }
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
        },
        issues: {
            desc: "Crear Issues para Baco Tell en GitHub.",
            buttons: {
                bug: "Informar de un error",
                feature: "Solicitar una función",
                question: "Haz una pregunta",
                other: "Crear otro Issue"
            }
        },
        clear: {
            cmd: "borras",
            desc: "Elimina varios mensajes en el canal.",
            reply: "¡Hecho!",
            option: {
                name: "cantidad",
                desc: "Cuántos mensajes eliminar"
            }
        },
        yeet: {
            desc: "Mueve repetidamente un miembro entre dos canales en rápida sucesión.",
            replies: {
                offline: "no está conectado.",
                yeet: "aquí vamos. ¡Abróchate el cinturón, baja y listo!"
            },
            options: {
                member: {
                    name: "miembro",
                    desc: "La persona que será yeeted"
                },
                amount: {
                    name: "cantidad",
                    desc: "El número de veces que la persona será yeeted"
                }
            }
        }
    }
}

export default es