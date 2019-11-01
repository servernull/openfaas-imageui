<template>
  <div class="analytics">
      Analytics
  </div>
</template>

<script>
import { EventBus } from "../event-bus.js";

export default {
  name: 'Analytics',
  props: {},
  created: function() {
    EventBus.$on("search/complete", this.update)
    EventBus.$on("search/refresh", this.update)
  },
  methods: {
    update : function(data) {
      console.log(data);
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.analytics {
    background-color: gray;
    height: 100%;
}
</style>
