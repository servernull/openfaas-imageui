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
          <div class="mdl-layout-spacer"></div>
          <button @click="clearSearch" class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect">
            Clear
          </button>
        </div>
        <!-- Tabs -->
        <div class="mdl-layout__tab-bar mdl-js-ripple-effect">
          <a href="#fixed-tab-1" class="mdl-layout__tab is-active">Images</a>
          <a href="#fixed-tab-2" class="mdl-layout__tab">Analytics</a>
          <a href="#fixed-tab-3" class="mdl-layout__tab">Links</a>
        </div>
      </header>
      <main class="mdl-layout__content">
        <section class="mdl-layout__tab-panel is-active" id="fixed-tab-1">
          <div class="page-content"><Images /></div>
        </section>
        <section class="mdl-layout__tab-panel" id="fixed-tab-2">
          <div class="page-content"><Analytics /></div>
        </section>
        <section class="mdl-layout__tab-panel" id="fixed-tab-3">
          <div class="page-content"><Links /></div>
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
      crawlUrl: '?action=crawl',
      searchUrl: '?action=search',
      clearUrl: '?action=clear',
    }
  },
  created: function() {
    console.log(process.env.NODE_ENV)
    if (process.env.NODE_ENV === 'development') {
      this.crawlUrl = '/action/crawl'
      this.searchUrl = '/action/search'
      this.clearUrl = '/action/clear'
    }
  },
  methods: {
    siteUrlKeydown: function() { 
      if (!this.url.startsWith("http")) {
        this.url = "http://" + this.url
      }
      this.clearSearch()
      EventBus.$emit("crawl/start", {})
      axios.post(this.crawlUrl,  this.url).then(response => {
        EventBus.$emit("crawl/complete", response.data)
      })
      this.stopRefresh()
      EventBus.$emit("search/start", {})
      axios.post(this.searchUrl, this.url).then(response => {
        EventBus.$emit("search/complete", response.data)
        this.refreshSearch()
      })
    },
    refreshSearch: function() {
      if (this.interval !== 0)
        return
      var sUrl = this.searchUrl
      var sPattern = this.url
      this.interval = setInterval(function() {
        axios.post(sUrl, sPattern).then(response => {
          EventBus.$emit("search/refresh", response.data)
        })
      }, 5000)
    },
    stopRefresh: function() {
      clearInterval(this.interval)
      this.interval = 0
    },
    clearSearch: function() {
      EventBus.$emit("crawl/complete",[])
      EventBus.$emit("search/complete", [])
      axios.post(this.clearUrl, "clear")
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