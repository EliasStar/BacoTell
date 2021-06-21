import { Command, Locale, Interaction, Embed } from "/command.ts"

const ping: Command = {
    command: (loc: Locale) => ({
        name: "ping",
        description: loc.cmds.ping.desc
    }),

    handler: (loc: Locale) => (async (inter: Interaction) => {
        inter.defer()
        const msg = await inter.send("Pong!")

        const gateway = inter.client.gateway.ping
        const latency = msg.timestamp.getTime() - inter.timestamp.getTime()

        const embed = new Embed()
            .setTitle(loc.cmds.ping.reply.results)
            .addField("Gateway Ping:", `${gateway} ms`)
            .addField(`Discord API ${loc.cmds.ping.reply.latency}:`, `${latency} ms`)

        if (gateway <= 60) {
            embed.setColor(0x00ff00)
        } else if (gateway <= 120) {
            embed.setColor(0xffff00)
        } else {
            embed.setColor(0xff0000)
        }

        await inter.reply({
            content: "Pong!",
            embeds: [embed]
        })
    })
}

export default ping