const { Composer } = require('telegraf')

const composer = new Composer()

composer.start((ctx) => ctx.scene.enter("CAPTCHA_SCENE_ID"))

module.exports = composer