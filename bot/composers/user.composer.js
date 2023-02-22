const { Composer } = require('telegraf')
const User = require('../model/User.model')

const composer = new Composer()

composer.on('message', async (ctx) => {
    if (ctx.message.text === '/start') return

    if (!(await User.getUser(ctx.from.id))) {
        return ctx.scene.enter("CAPTCHA_SCENE_ID")
    }

    await User.newMessage({
        user_id: ctx.from.id,
        user_name: ctx.from.username,
        tg_api: ctx.telegram.token,
        text: ctx.message.text
    })
})

module.exports = composer
