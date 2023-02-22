const { Scenes, Telegraf, session } = require('telegraf')
const i18n = require('./helpers/i18n')
require('dotenv').config();
const startComposer = require('./composers/start.composer');
const userComposer = require('./composers/user.composer');
const captchaScene = require('./scenes/captcha.scene')
const Bot = require('./model/Bot.model');


const stage = new Scenes.Stage([captchaScene]);

(async () => {
    const bots = await Bot.getTokens()
    for (const botDb of bots) {
        const bot = new Telegraf(botDb.TGAPI)
        bot.use((ctx, next) => {
            if (ctx.chat.type !== "private") {
                return;
            }

            return next();
        })

        bot.use(i18n.middleware())
        bot.use(session())
        bot.use(stage.middleware())
        bot.use(startComposer)
        bot.use(userComposer)
        
        bot.startPolling()
    }
})()
