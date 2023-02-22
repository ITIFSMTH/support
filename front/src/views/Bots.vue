<template>
  <div class="bots">
    <div class="controllers row">
      <Controllers/>
    </div>
    <div class="row">
      <Bots/>
    </div>
  </div>
</template>

<script>
  import Controllers from '../components/Bots/Controllers';
  import Bots from '../components/Bots/Bots';

  export default {
    name: 'bots',
    components: {
      Bots,
      Controllers
    },
    created() {
      document.title = "Orb11ta | Боты"
    }
  }
</script>

<style>
  .bots {
    padding: 10px 50px 10px;
  }

  @media (max-width: 700px) {
    .bots {
      padding: 10px;
    }
  }
</style>