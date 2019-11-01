<template>
  <div class="analytics">
    <div class="mdl-grid">
      <div class="mdl-cell mdl-cell--4-col mdl-cell--6-col-tablet">
          <div class="nsfw_stats">
            <span>Number of images: {{ imageCount }}</span>
            <span>Average NSFW Score: {{ averageNsfw.toPrecision(3) }}</span>
            <span>Minimum NSFW Score: {{ minNsfw.toPrecision(3) }}</span>
            <span>Maximum NSFW Score: {{ maxNsfw.toPrecision(3) }}</span>
          </div>
      </div>
      <div class="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet"></div>
      <div class="mdl-cell mdl-cell--2-col mdl-cell--4-col-phone"></div>
    </div>
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
  data: function() {
    return {
      imageCount: 0,
      averageNsfw: 0.00,
      minNsfw: 0.00,
      maxNsfw: 0.00,
    };
  },
  methods: {
    update : function(data) {
      this.imageCount = data.length
      var nsfw_nums = data.map(d => d.nsfw.nsfw_score)
      this.averageNsfw = nsfw_nums.reduce((a,b) => a + b, 0)/ data.length
      this.minNsfw = Math.min(...nsfw_nums)
      this.maxNsfw = Math.max(...nsfw_nums)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.analytics {
    height: 100%;
}
.nsfw_stats {
  display: flex;
  flex-direction: column;
}
</style>
