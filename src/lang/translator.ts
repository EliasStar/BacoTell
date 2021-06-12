import en from "./en.json"
import de from "./de.json"
import es from "./es.json";

type EnglishKeys = keyof typeof en
type GermanKeys = keyof typeof de
type SpanishKeys = keyof typeof es

export type Locale = { [K in EnglishKeys & GermanKeys & SpanishKeys]: string }
export type TranslationKey = keyof Locale

export class Translator {

    private currentLang: Locale

    constructor(defaultLang: Locale) {
        this.currentLang = defaultLang
    }

    setLang(lang: Locale) {
        this.currentLang = lang
    }

    getLang(): Locale {
        return this.currentLang
    }

    translate(key: TranslationKey): string {
        return this.currentLang[key]
    }
}

export default new Translator(en)
export const lang = {
    english: en as Locale,
    german: de as Locale,
    spanish: es as Locale
}