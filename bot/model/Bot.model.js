const axios = require("axios")
const config = require("config")

module.exports = class Bot {
    constructor() {}

    static async getTokens() {
        const { data } = await axios.get(config.get("URLs.bot") + "/bots", {
            headers: {
                'Authorization': `Bearer: ${process.env.BOT_JWT}`
            },
            validateStatus: () => true
        })

        return data.data.bots
    }

    static async getGreeting() {
        const { data } = await axios.get(config.get("URLs.bot") + "/greeting", {
            headers: {
                'Authorization': `Bearer: ${process.env.BOT_JWT}`
            },
            validateStatus: () => true
        })

        return data.data.greeting
    }
}