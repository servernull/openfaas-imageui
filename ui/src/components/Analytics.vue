<template>
  <div class="analytics">
    <div v-if="items.length == 0 && !hasData" class="empty">
      No images.  Please search above
    </div>
    <div v-if="items.length > 0" class="nsfw_stats">
      <div class="demo-card-event mdl-card mdl-shadow--2dp">
        <div class="mdl-card__title mdl-card--expand">
          <h4>
            Images:<br>
            {{ imageCount }}
          </h4>
        </div>
      </div>
      <div class="demo-card-event mdl-card mdl-shadow--2dp">
        <div class="mdl-card__title mdl-card--expand">
          <h4>
            Images with EXIF:<br>
            {{ exifCount }}
          </h4>
        </div>
      </div>
      <div class="demo-card-event mdl-card mdl-shadow--2dp">
        <div class="mdl-card__title mdl-card--expand">
          <h4>
            Averge NSFW Score:<br>
            {{ averageNsfw.toPrecision(3) }}
          </h4>
        </div>
      </div>
      <div class="demo-card-event mdl-card mdl-shadow--2dp">
        <div class="mdl-card__title mdl-card--expand">
          <h4>
            Maximum NSFW Score:<br>
            {{ maxNsfw.toPrecision(3) }}
          </h4>
        </div>
      </div>
      <div class="demo-card-event mdl-card mdl-shadow--2dp">
        <div class="mdl-card__title mdl-card--expand">
          <h4>
            Minimum NSFW Score:<br>
            {{ minNsfw.toPrecision(3) }}
          </h4>
        </div>
      </div>
    </div>
    <div v-if="items.length > 0" class="charts">
      <div id="chart">
        <apexchart type=pie width=500 :options="exifChartOptions" :series="exifSeries" />
      </div>
      <div id="chart2">
        <apexchart type=pie width=500 :options="nsfwChartOptions" :series="nsfwSeries" />
      </div>
    </div>
  </div>
</template>

<script>
import { EventBus } from "../event-bus.js";
import VueApexCharts from 'vue-apexcharts'

export default {
  name: 'Analytics',
  props: {},
  components: {
      apexchart: VueApexCharts,
  },
  created: function() {
    EventBus.$on("search/complete", this.update)
    EventBus.$on("search/refresh", this.update)
  },
  data: function() {
    return {
      hasData: false,
      items: [],
      imageCount: 0,
      exifCount: 0,
      averageNsfw: 0.00,
      minNsfw: 0.00,
      maxNsfw: 0.00,
      exifChartOptions: {
          labels: ['No EXIF', 'EXIF'],
          responsive: [{
            breakpoint: 480,
            options: {
              chart: {
                width: 200
              },
              legend: {
                position: 'bottom'
              }
            }
          }]
      },
      nsfwChartOptions: {
          labels: ['SFW > 0.5', 'NSFW >= 0.5'],
          responsive: [{
            breakpoint: 480,
            options: {
              chart: {
                width: 200
              },
              legend: {
                position: 'bottom'
              }
            }
          }]
      },
    };
  },
  computed: {
    exifSeries: function() {
      return  [this.imageCount, this.exifCount];
    },
    nsfwSeries: function() {
      return [this.items.filter(i => i.nsfw).filter(i => i.nsfw.nsfw_score < 0.5).length, this.items.filter(i => i.nsfw).filter(i => i.nsfw.nsfw_score >= 0.5).length]
    },
  },
  methods: {
    update : function(data) {
      this.items = data
      this.imageCount = data.length
      this.exifCount = data.filter(d => d.exif.length).length
      var nsfw_nums = data.filter(d => d.nsfw).map(d => d.nsfw.nsfw_score)
      this.averageNsfw = nsfw_nums.reduce((a,b) => a + b, 0)/ data.length
      this.minNsfw = Math.min(...nsfw_nums)
      this.maxNsfw = Math.max(...nsfw_nums)
      this.hasData = true;
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
}
.hidden {
  display: none;
}
.empty {
  text-align: center;
  padding-top: 100px;
  font-size: x-large;
  margin: 0 auto;
}
.demo-card-event.mdl-card {
  margin: 5px;
  background: #3E4EB8;
}
.demo-card-event > .mdl-card__actions {
  border-color: rgba(255, 255, 255, 0.2);
}
.demo-card-event > .mdl-card__title {
  align-items: flex-start;
}
.demo-card-event > .mdl-card__title > h4 {
  margin-top: 0;
}
.demo-card-event > .mdl-card__actions {
  display: flex;
  box-sizing:border-box;
  align-items: center;
}
.demo-card-event > .mdl-card__actions > .material-icons {
  padding-right: 10px;
}
.demo-card-event > .mdl-card__title,
.demo-card-event > .mdl-card__actions,
.demo-card-event > .mdl-card__actions > .mdl-button {
  color: #fff;
}
.charts {
  display: flex;
}
#chart, #chart2 {
  margin: 70px;
}
</style>
