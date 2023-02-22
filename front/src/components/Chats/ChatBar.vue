<template>
  <ul 
  ref="chatBar"
  id="chat-bar"
  class="chatbar sidenav sidenav-fixed light-blue lighten-5">
    <li
    v-for="chat in chats" 
    :key="chat.chat_id"
    @click="openChat(chat.chat_id)"
    class="chatbar-chat"
    :class="{'active-chat' : isChatActive(chat.chat_id)}">
      <div class="users">
        <span class="username">{{chat.user.Username}}</span>
        <br/>
        <span class="operators">
          <span v-for="(operator,key) of chat.operators">
            <span class="operator">{{operator.Login}}</span>
            <span v-if="key != chat.operators.length - 1">, </span>
          </span>
        </span>
      </div>
      <div 
      v-if="chat.unreaded_count > 0"
      class="unreaded-messages">
        {{chat.unreaded_count}}
      </div>
    </li>
  </ul>
</template>

<script>
export default {
  name: 'chatbar',
  data() {
    return {
      chatBar: {},
      open: false
    }
  },
  methods: {
    isChatActive(id) {
      if (id === this.$store.getters.chat.chat_id) return true
      return false
    },
    openChat(id) {
      this.$emit("openChat", id)
      if (this.open) this.chatBar.close()
      this.open = false
    },
    openChatBar() {
      this.chatBar.open()
      this.open = true
    }
  },
  mounted() {
    this.chatBar = M.Sidenav.init(this.$refs.chatBar, {});
  },
  computed: {
    chats() {
      return this.$store.getters.chats;
    }
  }
}
</script>

<style scoped>
.active-chat {
  background-color: rgb(0, 0, 0, 0.2) !important;
}

.chatbar {
  top: auto;
  box-shadow: 0 2px 2px 0 rgba(0,0,0,0.14),0 3px 1px -2px rgba(0,0,0,0.12),0 5px 5px 0 rgba(0,0,0,0.2);
}

.chatbar-chat {
  position: relative;
  color: white;
  line-height: normal;
  padding: 15px;
  background-color: rgba(0, 0, 0, 0.1);
  cursor: pointer;
}

.username {
  display: inline-block;
  font-weight: 600;
  color: #b388ff;
}

.operators {
  display: inline-block;
}

.operator {
  color: #d32f2f;
}

.users {
  display: inline-block;
}

.unreaded-messages {
  position: absolute;
  display: flex;
  right: 15px;
  top: 50%; 
  transform: translateY(-50%);
  width: 40px;
  height: 40px;
  padding: 5px;
  border-radius: 50%;
  justify-content: center;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.3);
}
</style>
