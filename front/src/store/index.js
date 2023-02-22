import { createStore } from "vuex";
import axios from 'axios';
import VueCookies from 'vue-cookies';
import jwtDecode from "jwt-decode";

let BASE_URL;
let WS_URL;

if (process.env.VUE_APP_PROD == "true") {
  BASE_URL = "http://" + location.host + "/api";
  WS_URL = "ws://" + location.host + "/api/operator/ws"
} else {
  BASE_URL = process.env.VUE_APP_BASEURL_DEV;
  WS_URL = process.env.VUE_APP_WSURL_DEV;
}

export default createStore({
  state: {
    login: {
      loginError: undefined,
      passwordError: undefined,
      login: '',
      password: ''
    },
    admin: {
      operators: [],
      bots: [],
      greeting: '',
      statistic: {},
      operatorEdit: {
        loginError: undefined,
        passwordError: undefined,
        login: '',
        password: '',
      },
      addOperator: {
        loginError: undefined,
        passwordError: undefined,
        login: '',
        password: ''
      },
      mailing: {
        mail: ''
      },
      addBot: {
        tgapi: ''
      }
    },
    chats: {
      ws: {},
      chats: [],
      chat: {
        message: '',
        textarea: {},
        chat_id: 0,
        messages: [],
        user: {}
      }
    }
  },
  getters: {
    operators(state) {
      return state.admin.operators
    },
    operatorStatistic(state) {
      return state.admin.statistic
    },
    bots(state) {
      return state.admin.bots
    },
    chats(state) {
      return state.chats.chats
    },
    chat(state) {
      return state.chats.chat
    },
    greeting(state) {
      return state.admin.greeting
    },
    isUserAdmin(state) {
      const data = jwtDecode(VueCookies.get('jwt'))
      return data.role === 'admin'
    }
  },
  mutations: {
    updateLogin(state, login) {
      if (login.length === 0) state.login.loginError = undefined
      else if (login.length < 3 || login.length > 20) state.login.loginError = true
      else state.login.loginError = false
      state.login.login = login;
    },
    updatePassword(state, password) {
      if (password.length === 0) state.admin.operatorEdit.passwordError = undefined
      else if (password.length < 8 || password.length > 80) state.login.passwordError = true
      else state.login.passwordError = false
      state.login.password = password;
    },
    setJWT(state, jwt) {
      const d = new Date();
      d.setTime(d.getTime() + 3 * 24 * 60 * 60 * 1000);
      VueCookies.set("jwt", jwt, d.toUTCString());
    },
    logout(state) {
      state.auth = {}
      VueCookies.remove("jwt")
    },
    setBots(state, bots) {
      state.admin.bots = bots
    },
    deleteBot(state, botID) {
      state.admin.bots = state.admin.bots.filter(bot => bot.ID !== botID)
    },
    addBot(state, bot) {
      state.admin.bots.push(bot)
    },
    setOperators(state, operators) {
      state.admin.operators = operators
    },
    setOperatorStatistic(state, statistic) {
      state.admin.statistic = statistic
    },
    setTextarea(state, textarea) {
      state.chats.chat.textarea = textarea;
    },
    deleteOperator(state, operatorID) {
      state.admin.operators = state.admin.operators.filter(operator => operator.ID !== operatorID)
    },
    addOperator(state, operator) {
      state.admin.operators.push(operator)
    },
    updateOperatorLoginState(state, login) {
      if (login.length === 0) state.admin.operatorEdit.loginError = undefined
      else if (login.length < 3 || login.length > 20) state.admin.operatorEdit.loginError = true
      else state.admin.operatorEdit.loginError = false
      state.admin.operatorEdit.login = login;
    },
    cancelUpdateOperatorLogin(state) {
      state.admin.operatorEdit.loginError = undefined
      state.admin.operatorEdit.login = ''
    },
    updateOperatorLogin(state, id, login) {
      state.admin.operators = state.admin.operators.map(operator => {
        if (operator.ID === id) operator.Login = login
        return operator
      })
    },
    updateOperatorPasswordState(state, password) {
      if (password.length === 0) state.admin.operatorEdit.passwordError = undefined
      else if (password.length < 8 || password.length > 80) state.admin.operatorEdit.passwordError = true
      else state.admin.operatorEdit.passwordError = false
      state.admin.operatorEdit.password = password;
    },
    cancelUpdateOperatorPassword(state) {
      state.admin.operatorEdit.passwordError = undefined
      state.admin.operatorEdit.password = ''
    },
    updateAddOperatorLogin(state, login) {
      if (login.length === 0) state.admin.addOperator.loginError = undefined
      else if (login.length < 3 || login.length > 20) state.admin.addOperator.loginError = true
      else state.admin.addOperator.loginError = false
      state.admin.addOperator.login = login;
    },
    updateAddOperatorPassword(state, password) {
      if (password.length === 0) state.admin.addOperator.passwordError = undefined
      else if (password.length < 8 || password.length > 80) state.admin.addOperator.passwordError = true
      else state.admin.addOperator.passwordError = false
      state.admin.addOperator.password = password;
    },
    cancelAddOperator(state) {
      state.admin.addOperator.loginError = undefined
      state.admin.addOperator.passwordError = undefined
      state.admin.addOperator.login = ''
      state.admin.addOperator.password = ''
    },
    cancelAddBot(state) {
      state.admin.addBot.tgapi = ""
    },
    setWS(state, ws) {
      state.chats.ws = ws
    },
    removeWS(state) {
      state.chats.ws = {}
    },
    setChats(state, chats) {
      if (chats) chats.sort((a,b) => Date.parse(new Date(b.last_message)) - Date.parse(new Date(a.last_message)));
      state.chats.chats = chats || []
    },
    setChat(state, data) {
      state.chats.chat.messages = data.messages
      state.chats.chat.user = data.user
      state.chats.chat.chat_id = data.chat_id
    },
    addChat(state, chat) {
      state.chats.chats.unshift(chat)
    },
    readChat(state, id) {
      state.chats.chats = state.chats.chats.map(chat => {
        if (chat.chat_id === id) {
          chat.unreaded_count = 0
        }
        return chat
      })
    },
    addMessage(state, message) {
      if (state.chats.chat.chat_id === message.message.ChatID) {
        state.chats.chat.messages.push(message)
      } else {
        state.chats.chats = state.chats.chats.map((chat) => {
          if (chat.chat_id === message.message.ChatID) chat.unreaded_count++
          return chat
        })
      }

      const thisChat = state.chats.chats.find((chat) => chat.chat_id === message.message.ChatID)
      const newChats = state.chats.chats.filter((chat) => chat.chat_id !== message.message.ChatID)
      
      newChats.unshift(thisChat)
      state.chats.chats = newChats
    },
    updateOperators(state, data) {
      state.chats.chats = state.chats.chats.map((chat) => {
        if (chat.chat_id !== data.chat_id) return chat

        const op = chat.operators.find(o => o.Login === data.sender)
        if (!op) chat.operators.push({Login: data.sender})
        
        return chat
      })
    },
    updateMessage(state, message) {
      state.chats.chat.message = message
    },
    updateMail(state, mail) {
      state.admin.mailing.mail = mail
    },
    clearMail(state) {
      state.admin.mailing.mail = ''
    },
    clearMessage(state) {
      state.chats.chat.message = ''
      state.chats.chat.textarea.style.height = '';
    },
    updateAddBotTgapi(state, tgapi) {
      state.admin.addBot.tgapi = tgapi;
    },
    updateGreeting(state, greeting) {
      state.admin.greeting = greeting
    }
  },
  actions: {
    async login(ctx) {
      if (this.state.login.loginError || this.state.login.passwordError || this.state.login.loginError === undefined || this.state.login.passwordError === undefined) return 'Поля заполнены неверно'

      const { data } = await axios.post(BASE_URL + '/auth/login', {
        login: this.state.login.login,
        password: this.state.login.password
      }, { validateStatus: () => true })
      if (data.error) M.toast({html: data.error})
      else M.toast({html: 'Успешно'})

      if (data.error) return false

      this.commit('setJWT', data.data.token)
      return true
    },
    async fetchOperators(ctx) {
      const { data } = await axios.get(BASE_URL + '/admin/operators', {
        headers: {
          'Authorization': `Bearer: ${VueCookies.get('jwt')}`
        }, validateStatus: () => true
      })
      if (data.error) return data.error

      this.commit('setOperators', data.data.operators)
    },
    async fetchOperatorStatistic(ctx, id) {
      const { data } = await axios.post(BASE_URL + '/admin/operator/statistic', {
        id: id
      },{
        headers: {
          'Authorization': `Bearer: ${VueCookies.get('jwt')}`
        }, 
        validateStatus: () => true
      })
      if (data.error) return data.error

      this.commit('setOperatorStatistic', data.data.statistic)
    },
    async fetchBots(ctx) {
      const { data } = await axios.get(BASE_URL + '/admin/bots', {
        headers: {
          'Authorization': `Bearer: ${VueCookies.get('jwt')}`
        }, validateStatus: () => true
      })
      if (data.error) return data.error

      this.commit('setBots', data.data.bots)
    },
    async fetchGreeting(ctx) {
      const { data } = await axios.get(BASE_URL + '/admin/greeting', {
        headers: {
          'Authorization': `Bearer: ${VueCookies.get('jwt')}`
        }, validateStatus: () => true
      })
      if (data.error) return data.error

      this.commit('updateGreeting', data.data.greeting)
    },
    async deleteOperator(ctx, id) {
      const { data } = await axios.post(BASE_URL + '/admin/operator/remove', {
        id: id,
      }, {
        headers: {
          'Authorization': `Bearer: ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })
      if (data.error) return M.toast({html: data.error})
      else M.toast({html: 'Успешно'})

      this.commit('deleteOperator', id)
    },
    async deleteBot(ctx, id) {
      const { data } = await axios.post(BASE_URL + '/admin/bot/remove', {
        id: id,
      }, {
        headers: {
          'Authorization': `Bearer: ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })
      if (data.error) return M.toast({html: data.error})
      else M.toast({html: 'Успешно'})

      this.commit('deleteBot', id)
    },
    async updateLoginOperator(ctx, id) {
      if (this.state.admin.operatorEdit.loginError || this.state.admin.operatorEdit.loginError === undefined) return 'Поля заполнены неверно'

      const { data } = await axios.post(BASE_URL + '/admin/operator/edit/login', {
        id: id,
        login: this.state.admin.operatorEdit.login
      }, {
        headers: {
          'Authorization': `Bearer: ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })
      if (data.error) return M.toast({html: data.error})
      else M.toast({html: 'Успешно'})

      this.commit('updateOperatorLogin', id, this.state.admin.operatorEdit.login)
      this.commit("cancelUpdateOperatorLogin")

      return false
    },
    async updatePasswordOperator(ctx, id) {
      if (this.state.admin.operatorEdit.passwordError || this.state.admin.operatorEdit.passwordError === undefined) return 'Поля заполнены неверно'

      const { data } = await axios.post(BASE_URL + '/admin/operator/edit/password', {
        id: id,
        password: this.state.admin.operatorEdit.password
      }, {
        headers: {
          'Authorization': `Bearer: ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })
      if (data.error) return M.toast({html: data.error})
      else M.toast({html: 'Успешно'})

      this.commit("cancelUpdateOperatorPassword")
    },
    async updateGreeting(ctx) {
      const { data } = await axios.post(BASE_URL + '/admin/edit/greeting', {
        greeting: this.state.admin.greeting
      }, {
        headers: {
          'Authorization': `Bearer: ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })
      if (data.error) return M.toast({html: data.error})
      else M.toast({html: 'Успешно'})
    },
    async addOperator(ctx) {
      if (this.state.admin.addOperator.loginError || this.state.admin.addOperator.passwordError || this.state.admin.addOperator.loginError === undefined || this.state.admin.addOperator.passwordError === undefined) return 'Поля заполнены неверно'

      const { data } = await axios.post(BASE_URL + '/admin/operator/add', {
        login: this.state.admin.addOperator.login,
        password: this.state.admin.addOperator.password
      }, {
        headers: {
          'Authorization': `Bearer: ${VueCookies.get('jwt')}`
        }, validateStatus: () => true
      })
      if (data.error) return M.toast({html: data.error})
      else M.toast({html: 'Успешно'})

      this.commit("addOperator", data.data.operator)
      this.commit("cancelAddOperator")
    },
    async addBot(ctx) {
      const { data } = await axios.post(BASE_URL + '/admin/bot/add', {
        tg_api: this.state.admin.addBot.tgapi,
      }, {
        headers: {
          'Authorization': `Bearer: ${VueCookies.get('jwt')}`
        }, validateStatus: () => true
      })
      if (data.error) return M.toast({html: data.error})
      else M.toast({html: 'Успешно'})

      this.commit("addBot", data.data.bot)
      this.commit("cancelAddBot")
    },
    async sendMail(ctx) {
      const { data } = await axios.post(BASE_URL + '/admin/mailing', {
        mail: this.state.admin.mailing.mail,
      }, {
        headers: {
          'Authorization': `Bearer: ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })
      if (data.error) return M.toast({html: data.error})
      else M.toast({html: 'Успешно'})

      this.commit("clearMail")
    },
    async openWS(ctx) {
      const ws = await connect()
      ws.onmessage = msg => this.dispatch("onSocketMessage", msg)
      this.commit('setWS', ws)
    },
    closeWS(ctx) {
      if (this.state.chats.ws.close) this.state.chats.ws.close()
      this.commit('removeWS')
    },
    sendReqChats(ctx) {
      const data = JSON.stringify({
        "action": "get chats",
        "id": (new Date()).getTime(),
      })

      this.state.chats.ws.send(data)
    },
    sendReqChat(ctx, id) {
      const data = JSON.stringify({
        "action": "get chat",
        "id": (new Date()).getTime(),
        "chat_id": id,
      })

      this.state.chats.ws.send(data)
      this.dispatch("sendRead", id)
    },
    sendRead(ctx, id) {
      const data = JSON.stringify({
        "action": "read chat",
        "id": (new Date()).getTime(),
        "chat_id": id
      })

      this.state.chats.ws.send(data)
    },
    sendMessage(ctx) {
      const data = JSON.stringify({
        "action": "send message",
        "id": (new Date()).getTime(),
        "chat_id": this.state.chats.chat.chat_id,
        "text": this.state.chats.chat.message
      })

      this.commit("clearMessage")
      this.state.chats.ws.send(data)
    },
    onSocketMessage({state}, msg) {
      const data = JSON.parse(msg.data)
      console.log(data)
      if (data.action === "get chats") {
        this.commit("setChats", data.chats)
      } else if (data.action === "get chat") {
        this.commit("setChat", data)
      } else if (data.action === "new message") {
        if (state.chats.chat.chat_id === data.chat_id) {
          this.dispatch("sendRead", data.chat_id)
        }

        if (data.user_id < 10000) {
          this.commit("updateOperators", {
            chat_id: data.chat_id,
            sender: data.sender,
          })
        }

        this.commit("addMessage", {
          message: {
            ChatID: data.chat_id,
            UserID: data.user_id,
            Text: data.text,
          },
          sender_name: data.sender,
        })
      } else if (data.action === "new chat") {
        this.commit("addChat", {
          chat_id: data.chat_id,
          operators: [],
          unreaded_count: 0,
          user: data.user
        })
      } else if (data.action === "readed chat") {
        this.commit("readChat", data.chat_id)
      }
    }
  },
});

const connect = () => {
  return new Promise(function (resolve, reject) {
    const ws = new WebSocket(`${WS_URL}?token=Bearer: ${VueCookies.get('jwt')}`)
    ws.onopen = function () {
      resolve(ws);
    };
    ws.onerror = function (err) {
      reject(err);
    };
  });
}
