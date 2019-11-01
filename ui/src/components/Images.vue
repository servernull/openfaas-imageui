<template>
  <div class="images">
    <div id="p2" v-bind:class="{ hidden: !searching }" class="mdl-progress mdl-js-progress mdl-progress__indeterminate"></div>
    <div v-if="samples.length == 0 && !searching" class="empty">
      No images.  Please search above
    </div>
    <div class="demo-card-wide mdl-card mdl-shadow--2dp" v-bind:key="sample" v-for="sample in samples">
      <div class="mdl-card__title" v-bind:style="{ background: 'url(' + sample.url + ') center / cover' }">
        <h2 class="mdl-card__title-text">{{ sample.filename }}</h2>
      </div>
      <div class="mdl-card__supporting-text">
        <div class="info">
          <label for="nsfw">NSFW score</label>
          <p id="nsfw">{{ sample.nsfw.nsfw_score.toPrecision(3) }}</p>
          <label for="exif">EXIF</label>
          <textarea id="exif" v-model="sample.exif"></textarea>
        </div>
      </div>
      <div class="mdl-card__actions mdl-card--border">
        <a target="_blank" :href="sample.url" class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">
          View Image
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import { EventBus } from "../event-bus.js";

export default {
  name: 'Links',
  created: function() {
    EventBus.$on("search/complete", this.searchComplete)
    EventBus.$on("search/start", this.searchStart)
    EventBus.$on("search/refresh", this.searchRefresh)
  },
  data: function () {
    return {
      samples: [],
      searching: false,
    };
  },
  methods: {
    searchComplete : function(data){
      data = data.map(d => {
        if (d.exif.length > 0) {
          d.exif = d.exif.map(e => {
              var key = Object.keys(e)[0]
              var value = Object.values(e)[0]
              return key+"="+value;
          })
        }
        d.filename = d.url.substring(d.url.lastIndexOf('/')+1);
        return d;
      });
      this.samples = data;
      this.searching = false;
    },
    searchStart : function() {
      this.searching = true;
      this.samples = [];
    },
    searchRefresh: function() {
      this.searching = false;
    }
  },
}
</script>

<style scoped>
.images {
  display: flex;
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
.demo-card-wide.mdl-card {
  width: 300px;
}
.demo-card-wide > .mdl-card__menu {
  color: #fff;
}
.demo-card-wide > .mdl-card__title {
  color: #fff;
  height: 300px;
}
.info {
  flex-direction: column;
  display: flex;
}
</style>
