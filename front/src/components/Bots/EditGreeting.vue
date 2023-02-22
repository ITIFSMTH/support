<template>
  <form 
  @submit.prevent="editGreeting" 
  class="controller-card card small light-blue lighten-5 col s12 m12 l4 xl4">
    <div class="card-content">
      <span class="card-title">Изменить приветствие</span>

      <div class="inputs">
        <div class="input-field">
          <textarea v-model="greeting" class="input-obj materialize-textarea" id="greeting"></textarea>
          <label
            class="active"
            for="greeting"
            >Приветствие</label>
        </div>
      </div>

      <div class="buttons center">
        <button class="button-login btn waves-effect green waves-light" type="submit">
          <span class="button-text">Изменить</span>
          <i class="button-icon material-icons">edit</i>
        </button>
      </div>
    </div>
  </form>
</template>

<script>
export default {
  name: 'editGreeting',
  computed: {
    greeting: {
      get() {
        return this.$store.state.admin.greeting;
      },
      set(greeting) {
        this.$store.commit("updateGreeting", greeting);
      },
    },
  },
  methods: {
    async editGreeting() {
      await this.$store.dispatch("updateGreeting")
    }
  },
  mounted() {
      this.$store.dispatch("fetchGreeting")
  }
}
</script>