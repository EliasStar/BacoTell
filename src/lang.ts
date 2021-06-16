export type LocaleIdentifier = "en" | "de" | "es"
export type Locale = {
    cmds: {
        ping: {
            desc: string
        }
        peng: {
            desc: string
            reply: string
        }
    }
}

export class Translator {

    private currentLang: Locale

    constructor(defaultLang: Locale) {
        this.currentLang = defaultLang
    }

    async setLang(lang: LocaleIdentifier) {
        try {
            this.currentLang = await import(`./lang/${lang}.ts`) as Locale
        } catch (e) {
            console.error(`cannot load locale with '${lang}' identifier: ${e}`)
        }
    }

    lang(): Locale {
        return this.currentLang
    }
}