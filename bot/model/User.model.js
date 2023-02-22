const axios = require("axios")
const config = require("config")
const errors = require("../helpers/errors")


module.exports = class User {
    constructor() {}

    static async getUser(tgId) {
        const { data } = await axios.get(config.get("URLs.bot") + "/getuser", {
            headers: {
                'Authorization': `Bearer: ${process.env.BOT_JWT}`
            },
            data: {
                tg_id: tgId
            },
            validateStatus: () => true
        })
        if (data.data.user.ID === 0) return false
        return data.data.user
    }

    static async newUser(userData) {
        const { data } = await axios.post(config.get("URLs.bot") + "/adduser", {
            tg_id: userData.id,
            bot_api: userData.bot_api,
            username: userData.username,
            tg_link: userData.tg_link
        }, {
            headers: {
                'Authorization': `Bearer: ${process.env.BOT_JWT}`
            },
            validateStatus: () => true
        })

        if (data.error) {
            return data.error
        }
    }

    static async newMessage(messageData) {
        const { data } = await axios.post(config.get("URLs.bot") + "/addmessage", messageData, {
            headers: {
                'Authorization': `Bearer: ${process.env.BOT_JWT}`
            },
            validateStatus: () => true
        })

        if (data.error) {
            return data.error
        }
    }
}