import en from "/lang/en.ts"

export type Locale = typeof en
export type LocaleIdentifier = "en" | "de" | "es"

const localeDB = new Map<string, LocaleIdentifier>() // TODO: Add DB

export async function loadLocale(identifier: LocaleIdentifier): Promise<Locale> {
    try {
        return (await import(`/lang/${identifier}.ts`)).default as Locale
    } catch (error) {
        console.error(`error while loading locale "${identifier}": ${error}`)
        return en
    }
}

export function getLocaleFromGuild(guildId: string): LocaleIdentifier {
    const locale = localeDB.get(guildId)
    return locale != null ? locale : "en"
}

export function setLocaleForGuild(guildId: string, locale: LocaleIdentifier) {
    localeDB.set(guildId, locale)
}