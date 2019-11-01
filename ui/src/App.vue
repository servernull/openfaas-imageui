<template>
  <div id="app">
    <div class="mdl-layout mdl-js-layout mdl-layout--fixed-header
                mdl-layout--fixed-tabs">
      <header class="mdl-layout__header">
        <div class="mdl-layout__header-row">
          <!-- Title -->
          <span class="mdl-layout-title">Crawl and search for images</span>
          <div class="mdl-textfield mdl-js-textfield mdl-textfield--expandable
                      mdl-textfield--floating-label mdl-textfield--align-right">
            <label class="mdl-button mdl-js-button mdl-button--icon"
                  for="site">
              <i class="material-icons">search</i>
            </label>
            <div class="mdl-textfield__expandable-holder">
              <input @keyup.enter="siteUrlKeydown"  v-model="url" class="mdl-textfield__input" type="text" name="sample"
                    id="site">
            </div>
          </div>
        </div>
        <!-- Tabs -->
        <div class="mdl-layout__tab-bar mdl-js-ripple-effect">
          <a href="#fixed-tab-1" class="mdl-layout__tab is-active">Links</a>
          <a href="#fixed-tab-2" class="mdl-layout__tab">Images</a>
          <a href="#fixed-tab-3" class="mdl-layout__tab">Analytics</a>
        </div>
      </header>
      <main class="mdl-layout__content">
        <section class="mdl-layout__tab-panel is-active" id="fixed-tab-1">
          <div class="page-content"><Links /></div>
        </section>
        <section class="mdl-layout__tab-panel" id="fixed-tab-2">
          <div class="page-content"><Images /></div>
        </section>
        <section class="mdl-layout__tab-panel" id="fixed-tab-3">
          <div class="page-content"><Analytics /></div>
        </section>
      </main>
    </div>
  </div>
</template>

<script>
import Images from './components/Images.vue'
import Links from './components/Links.vue' 
import Analytics from './components/Analytics.vue'
import axios from 'axios';
import { EventBus } from "./event-bus.js";

export default {
  name: 'app',
  components: {
    Images,
    Links,
    Analytics
  },
  data() {
    return {
      url: '',
      interval: 0,
    }
  },
  methods: {
    siteUrlKeydown: function() { 
      EventBus.$emit("crawl/start", {})
      axios.post('/action/crawl?action=crawl',  this.url).then(response => {
        EventBus.$emit("crawl/complete", response.data)
      })
      this.stopRefresh()
      EventBus.$emit("search/start", {})
      axios.post('/action/search?action=search', this.url).then(response => {
        EventBus.$emit("search/complete", response.data)
        // this.refreshSearch()
      })
    },
    refreshSearch: function() {
      this.interval = setInterval(function() {
        axios.post('/action/search?action=search', this.url).then(response => {
          EventBus.$emit("search/refresh", response.data)
        })
      }, 5000)
    },
    stopRefresh: function() {
      clearInterval(this.interval)
    }
  }
}
</script>

<style>
#app {
}
.page-content {
  height: 100%;
}
.mdl-layout__tab-panel {
  height: 100%;
}
</style>
