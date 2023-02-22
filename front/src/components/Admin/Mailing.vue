<template>
  <form 
  @submit.prevent="sendMail" 
  class="controller-card card small light-blue lighten-5 col s12 m12 l4 xl4">
    <div class="card-content">
      <span class="card-title">Рассылка</span>

      <div class="inputs">
        <div class="input-field">
          <textarea v-model="mail" class="input-obj materialize-textarea" id="mail"></textarea>
          <label
            class="active"
            for="mail"
            >Сообщение</label>
        </div>
      </div>

      <div class="buttons center">
        <button class="button-login btn waves-effect green waves-light" type="submit">
          <span class="button-text">Отправить</span>
          <i class="button-icon material-icons">send</i>
        </button>
      </div>
    </div>
  </form>
</template>

<script>
export default {
  name: 'mailing',
  computed: {
    mail: {
      get() {
        return this.$store.state.admin.mailing.mail;
      },
      set(mail) {
        this.$store.commit("updateMail", mail);
      },
    }
  },
  methods: {
    async sendMail() {
      await this.$store.dispatch("sendMail")
    }
  }
}
</script>