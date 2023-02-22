const TelegrafI18n = require('telegraf-i18n')
const path = require('path')

const i18n = new TelegrafI18n({
    defaultLanguage: 'ru',
    allowMissing: false,
    directory: path.resolve('locales')
})

module.exports = i18n