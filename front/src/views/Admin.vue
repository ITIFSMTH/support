<template>
  <div class="admin">
    <div class="controllers row">
      <Controllers/>
    </div>
    <div class="row">
      <Operators/>
    </div>
  </div>
</template>

<script>
  import Controllers from '../components/Admin/Controllers';
  import Operators from '../components/Admin/Operators';

  export default {
    name: 'admin',
    components: {
      Operators,
      Controllers
    },
    created() {
      document.title = "Orb11ta | Админ"
    }
  }
</script>

<style>
  .admin {
    padding: 10px 50px 10px;
  }

  @media (max-width: 700px) {
    .admin {
      padding: 10px;
    }
  }
</style>