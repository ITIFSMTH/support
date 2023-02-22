<template>
  <div>
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Бот</th>
          <th>Создан</th>
          <th class="right">Действия</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="bot in bots" :key="bot.ID">
          <td>{{ bot.ID }}</td>
          <td>@{{ bot.TGName }}</td>
          <td>{{ dateTime(bot.CreatedAt) }}</td>
          <td class="right">
            <span
              @click="deleteModalOpen(bot.ID, bot.TGName)"
              class="action red-text"
              >Удалить</span
            >
          </td>
        </tr>
      </tbody>
    </table>

    <div ref="deleteModal" id="delete-modal" class="modal">
      <div class="modal-content">
        <h6>Вы уверены, что хотите удалить бота {{ botToDelete.name }}?</h6>
      </div>
      <div class="modal-footer">
        <a class="modal-close white-text green waves-effect btn-flat">Отмена</a>
        <a
          @click="deleteBot()"
          class="modal-close white-text red waves-effect btn-flat"
        >
          Удалить
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import moment from "moment";

export default {
  name: "bots",
  data() {
    return {
      deleteModal: {},
      botToDelete: {
        id: 0,
        name: "",
      },
    };
  },
  methods: {
    dateTime(value) {
      return moment(value).format("YYYY-MM-DD hh:mm:ss");
    },
    deleteModalOpen(id, name) {
      this.botToDelete = {
        id: id,
        name: name,
      };
      this.deleteModal.open();
    },
    async deleteBot() {
      await this.$store.dispatch("deleteBot", this.botToDelete.id);
    },
  },
  computed: {
    bots() {
      return this.$store.getters.bots;
    },
  },
  mounted() {
    this.$store.dispatch("fetchBots");

    this.deleteModal = M.Modal.init(this.$refs.deleteModal, {
      dismissible: false
    });
  },
};
</script>

<style scoped>
th:first-child,
td:first-child,
th:last-child,
td:last-child {
  padding: 15px;
}

tbody tr:nth-child(2n + 1) {
  background-color: #e1f5fe;
}

.action {
  cursor: pointer;
}

.modal,
.modal-footer {
  background-color: white;
}

.modal-footer {
  padding-left: 24px;
  padding-right: 24px;
}

.modal-footer a {
  margin: 0 0 0 20px !important;
}

.input-obj {
  color: white;
  background: none;
}

.inputs {
  padding-top: 10px;
}
</style>