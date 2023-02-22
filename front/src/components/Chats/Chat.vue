<template>
  <div class="chat-wrapper">
    <div class="chat" v-if="chat.chat_id !== 0">
      <div class="messages">
        <div 
        v-for="(message, idx) in chat.messages"
        :key="message.message.ID"
        :class="message.message.UserID > 10000 ? 'message-user' : 'message-operator'"
        class="message-wrap">
          <div style="margin-top: 0; margin-bottom: 25px;" class="newDay" v-if="idx === 0">
            <span class="newDay-text">{{ getDate(message.message.CreatedAt) }}</span>
          </div>
          <div class="message">
            <div class="message-header">
              <span class="message-title">{{ message.sender_name }}</span>
              <span class="message-time  blue-grey-text text-lighten-3">{{ getTime(message.message.CreatedAt)}}</span>
            </div>
            <p class="message-text" v-html="message.message.Text.split('\n').join('<br/>')"></p>
          </div>
          <div class="newDay" v-if="isNextDay(message.message.CreatedAt, chat.messages[idx+1] ? chat.messages[idx+1].message.CreatedAt : message.message.CreatedAt)">
            <span class="newDay-text">{{ getDate(chat.messages[idx+1].message.CreatedAt) }}</span>
          </div>
        </div>
      </div>
      <form
      @submit.prevent="sendMessage" 
      class="send-message-form">
        <button class="usernameButton button btn waves-effect green waves-light" type="button"
        @click="getUsername()">
          <i class="material-icons">account_circle</i>
        </button>
        <div class="input-field">
          <textarea
            v-model="message"
            v-on:keyup.enter="sendMessage"
            class="input-obj materialize-textarea"
            id="mail"
            ref="messageTextarea"
            placeholder="Введите сообщение"
          ></textarea>
        </div>
        <button 
        class="button btn waves-effect green waves-light" 
        type="submit">
          <span class="button-text">Отправить</span>
          <i class="button-icon material-icons">login</i>
        </button>
      </form>
    </div>

    <div v-if="chat.chat_id === 0" class="no-chat">
      <div class="no-chat-text">
        Выберите чат для продолжения
      </div>
    </div>
  </div>
</template>

<script>
//chat.user.TGLink !== "" ? `(@${chat.user.TGLink})` : ""
import moment from 'moment';
moment.locale("ru")

export default {
  name: "chat",
  methods: {
    getTime(value) {
      return moment(value).format("HH:mm");
    },
    getDate(value) {
      return moment(value).format("D MMMM")
    },
    getUsername() {
      const username = this.$store.state.chats.chat.user.TGLink
      if (username.length > 0) M.toast({html: '@' + username})
      else M.toast({html: "У этого пользователя нет Username."})
    },
    isNextDay(from, to) {
      return moment(from).format("D") !== moment(to).format("D")
    },
    sendMessage() {
      this.$store.dispatch("sendMessage", this.$refs.messageTextarea);
    }
  },
  computed: {
    chat() {
      return this.$store.getters.chat;
    },
    message: {
      get() {
        return this.$store.state.chats.chat.message;
      },
      set(message) {
        this.$store.commit("updateMessage", message);
      },
    },
  }, 
  updated() {
    window.scrollTo(0, document.body.scrollHeight || document.documentElement.scrollHeight);
    this.$store.commit("setTextarea", this.$refs.messageTextarea);
},
  mounted() {
    window.scrollTo(0, document.body.scrollHeight || document.documentElement.scrollHeight);
  }
};
</script>

<style scoped>
.newDay {
  width: 100%;
  position: relative;
  text-align: center;
  padding: 5px 0;
  margin-top: 5px;
}

.newDay-text {
  position: relative;
  padding: 0 10px;
  z-index: 20;
  color: #263238;
  background-color: #f5f5f5;
}

.newDay::before {
  content: '';
  position: absolute;
  left: 0;
  right: 0;
  top: 50%;
  transform: translateY(50%);
  height: 1px;
  background-color: #263238;
}

.message-time {
  font-size: 12px;
  color: grey;
}

.no-chat {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}

.no-chat-text { 
  padding: 5px 15px;
  background-color: rgba(0, 0, 0, 0.3); 
  box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.12),
    0 1px 5px 0 rgba(0, 0, 0, 0.2);
  border-radius: 10px;
}

.messages {
  width: 100%;
  display: flex;
  flex-wrap: wrap;
  padding-bottom: 50px;
}

.message-wrap {
  display: flex;
  flex-wrap: wrap;
  width: 100%;
  margin: 5px 0;
}

.message-operator {
  justify-content: end;
}

.message-operator .message-title {
  color: #d32f2f;
}

.message-user .message-title {
  color: #b388ff;
}

.message-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.message-time {
  padding-left: 20px;
}

.message {
  padding: 10px;
  background-color: #e0e0e0;
  border-radius: 3px;
}

.message-text {
  padding-top: 5px;
  margin: 0;
}

.chat-wrapper {
  display: flex;
  
  min-height: calc(100vh - 64px);
}

.chat {
  display: flex;
  flex-wrap: wrap;
  align-content: space-between;

  max-width: 992px;
  width: 100%;
  padding: 20px;
  margin: 0 auto;
  background-color: #f5f5f5;
  box-sizing: content-box;
}

.send-message-form {
  position: fixed;
  bottom: 0;
  width: 100%;
  max-width: inherit;
  display: flex;
  align-items: center;
  padding: 5px 10px;
  background-color: #e1f5fe ;
  border-radius: 3px;
  box-sizing: border-box;
}

.input-obj {
  color: #263238;
  background: none;
}

.input-obj::placeholder {
  color: #263238;
}

.input-field {
  width: 100%;
  margin: 0 20px 0 20px;
}

.button-text {
  font-family: "Roboto", sans-serif;
  font-weight: 500;
  text-transform: none;
  vertical-align: top;
}

.button-icon {
  margin-left: 5px;
}

.usernameButton {
  width: 40px;
  padding: 0 5px;
}

@media only screen and (min-width: 993px) {
  .chat-wrapper {
    margin-left: 300px;
  }
}


@media only screen and (max-width: 992px) {
  .send-message-form {
    width: auto;
    right: 20px;
    left: 20px;
  }
}
</style>
