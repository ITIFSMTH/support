<template>
  <div class="wrapper">
    <ChatBar
      ref="chatBar"
      @openChat="openChat"
    />

    <Chat/>
  </div>
</template>

<script>
  import ChatBar from '../components/Chats/ChatBar'
  import Chat from '../components/Chats/Chat'

  export default {
    name: 'chats',
    components: {
      ChatBar,
      Chat
    },
    methods: {
      openChat(id) {
        this.$store.dispatch("sendReqChat", id);
      },
      openChatBar() {
        this.$refs.chatBar.openChatBar()
      }
    },
    async mounted() {
      await this.$store.dispatch("openWS");
      this.$store.dispatch("sendReqChats");
    },
    beforeRouteLeave (to, from, next) {
      this.$store.dispatch("closeWS");
      next()
    },
    created() {
      document.title = "Orb11ta | Чаты"
    }
  }
</script>

<style scoped>
</style>