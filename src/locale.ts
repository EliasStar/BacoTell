import defaultLang from "/lang/en.ts"

export type Locale = typeof defaultLang
export type LocaleIdentifier = "en" | "de" | "es"

const localeDB = new Map<string, LocaleIdentifier>() // TODO: Add DB

export async function loadLocale(identifier: LocaleIdentifier): Promise<Locale> {
    try {
        return (await import(`/lang/${identifier}.ts`)).default as Locale
    } catch (error) {
        console.error(`loading locale "${identifier}" error: ${error}`)
        return defaultLang
    }
}

export async function getLocaleFromGuild(guild: string): Promise<Locale> {
    const locale = localeDB.get(guild)
    return locale != null ? await loadLocale(locale) : defaultLang
}

export function setLocaleForGuild(guild: string, locale: LocaleIdentifier) {
    localeDB.set(guild, locale)
}