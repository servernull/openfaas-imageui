<template>
  <div class="links">
    <div id="p2" v-bind:class="{ hidden: !crawling }" class="mdl-progress mdl-js-progress mdl-progress__indeterminate"></div>
    <div v-if="samples.length == 0 && !crawling" class="empty">
      No image links.  Please search above
    </div>
    <ul v-if="samples.length > 0">
      <li  v-bind:key="sample" v-for="sample in samples">
        <a :href="sample">{{sample}}</a>
      </li>  
    </ul>
  </div>
</template>

<script>
import { EventBus } from "../event-bus.js";

export default {
  name: 'Links',
  created: function() {
    EventBus.$on("crawl/complete", this.samplesReady)
    EventBus.$on("crawl/start", this.crawlStart)
  },
  data: function () {
    return {
      samples: [],
      crawling: false,
    };
  },
  methods: {
    samplesReady : function(data){
      this.samples = data;
      this.crawling = false;
    },
    crawlStart : function() {
      this.crawling = true;
      this.samples = [];
    }
  },
}
</script>

<style scoped>
.links {

}
.empty {
  text-align: center;
  padding-top: 100px;
  font-size: x-large;
}
#p2 {
  width: 100%;
}
.hidden {
  display: none;
}
</style>
