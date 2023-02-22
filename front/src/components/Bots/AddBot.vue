<template>
  <form 
  @submit.prevent="addBot" 
  class="controller-card card small light-blue lighten-5 col s12 m12 l4 xl4">
    <div class="card-content">
      <span class="card-title">Добавить бота</span>

      <div class="inputs">
        <div class="input-field">
          <input v-model="tgapi" class="input-obj validate" id="login" type="text" />
          <label
            class="active"
            for="login"
            >Токен бота</label
          >
        </div>
      </div>

      <div class="buttons center">
        <button class="button-login btn waves-effect green waves-light" type="submit">
          <span class="button-text">Добавить</span>
          <i class="button-icon material-icons">add</i>
        </button>
      </div>
    </div>
  </form>
</template>

<script>
export default {
  name: 'addBot',
  computed: {
    tgapi: {
      get() {
        return this.$store.state.admin.addBot.tgapi;
      },
      set(tgapi) {
        this.$store.commit("updateAddBotTgapi", tgapi);
      },
    },
  },
  methods: {
    async addBot() {
      await this.$store.dispatch("addBot")
    }
  }
}
</script>