(function(t){function e(e){for(var s,n,l=e[0],c=e[1],o=e[2],u=0,f=[];u<l.length;u++)n=l[u],Object.prototype.hasOwnProperty.call(i,n)&&i[n]&&f.push(i[n][0]),i[n]=0;for(s in c)Object.prototype.hasOwnProperty.call(c,s)&&(t[s]=c[s]);d&&d(e);while(f.length)f.shift()();return r.push.apply(r,o||[]),a()}function a(){for(var t,e=0;e<r.length;e++){for(var a=r[e],s=!0,l=1;l<a.length;l++){var c=a[l];0!==i[c]&&(s=!1)}s&&(r.splice(e--,1),t=n(n.s=a[0]))}return t}var s={},i={app:0},r=[];function n(e){if(s[e])return s[e].exports;var a=s[e]={i:e,l:!1,exports:{}};return t[e].call(a.exports,a,a.exports,n),a.l=!0,a.exports}n.m=t,n.c=s,n.d=function(t,e,a){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:a})},n.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var a=Object.create(null);if(n.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var s in t)n.d(a,s,function(e){return t[e]}.bind(null,s));return a},n.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/";var l=window["webpackJsonp"]=window["webpackJsonp"]||[],c=l.push.bind(l);l.push=e,l=l.slice();for(var o=0;o<l.length;o++)e(l[o]);var d=c;r.push([0,"chunk-vendors"]),a()})({0:function(t,e,a){t.exports=a("56d7")},"027f":function(t,e,a){},"034f":function(t,e,a){"use strict";var s=a("85ec"),i=a.n(s);i.a},"10b3":function(t,e,a){"use strict";var s=a("37eb"),i=a.n(s);i.a},"37eb":function(t,e,a){},"56d7":function(t,e,a){"use strict";a.r(e);a("e260"),a("e6cf"),a("cca6"),a("a79d");var s=a("2b0e"),i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"app"}},[a("div",{staticClass:"mdl-layout mdl-js-layout mdl-layout--fixed-header\n              mdl-layout--fixed-tabs"},[a("header",{staticClass:"mdl-layout__header"},[a("div",{staticClass:"mdl-layout__header-row"},[a("span",{staticClass:"mdl-layout-title"},[t._v("Crawl and search for images")]),a("div",{staticClass:"mdl-textfield mdl-js-textfield mdl-textfield--expandable\n                    mdl-textfield--floating-label mdl-textfield--align-right"},[t._m(0),a("div",{staticClass:"mdl-textfield__expandable-holder"},[a("input",{directives:[{name:"model",rawName:"v-model",value:t.url,expression:"url"}],staticClass:"mdl-textfield__input",attrs:{type:"text",name:"sample",id:"site"},domProps:{value:t.url},on:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.siteUrlKeydown(e)},input:function(e){e.target.composing||(t.url=e.target.value)}}})])]),a("div",{staticClass:"mdl-layout-spacer"}),a("button",{staticClass:"mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect",on:{click:t.clearSearch}},[t._v(" Clear ")])]),t._m(1)]),a("main",{staticClass:"mdl-layout__content"},[a("section",{staticClass:"mdl-layout__tab-panel is-active",attrs:{id:"fixed-tab-1"}},[a("div",{staticClass:"page-content"},[a("Images")],1)]),a("section",{staticClass:"mdl-layout__tab-panel",attrs:{id:"fixed-tab-2"}},[a("div",{staticClass:"page-content"},[a("Analytics")],1)]),a("section",{staticClass:"mdl-layout__tab-panel",attrs:{id:"fixed-tab-3"}},[a("div",{staticClass:"page-content"},[a("Links")],1)])])])])},r=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("label",{staticClass:"mdl-button mdl-js-button mdl-button--icon",attrs:{for:"site"}},[a("i",{staticClass:"material-icons"},[t._v("search")])])},function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"mdl-layout__tab-bar mdl-js-ripple-effect"},[a("a",{staticClass:"mdl-layout__tab is-active",attrs:{href:"#fixed-tab-1"}},[t._v("Images")]),a("a",{staticClass:"mdl-layout__tab",attrs:{href:"#fixed-tab-2"}},[t._v("Analytics")]),a("a",{staticClass:"mdl-layout__tab",attrs:{href:"#fixed-tab-3"}},[t._v("Links")])])}],n=(a("2ca0"),function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"images"},[a("div",{staticClass:"mdl-progress mdl-js-progress mdl-progress__indeterminate",class:{hidden:!t.searching},attrs:{id:"p2"}}),0!=t.samples.length||t.searching?t._e():a("div",{staticClass:"empty"},[t._v(" No images. Please search above ")]),a("div",{staticClass:"cards"},t._l(t.samples,(function(e){return a("div",{key:e.id,staticClass:"demo-card-wide mdl-card mdl-shadow--2dp"},[a("div",{staticClass:"mdl-card__title",style:{background:"url("+e.url+") center / cover"}},[a("h2",{staticClass:"mdl-card__title-text"},[t._v(t._s(e.filename))])]),a("div",{staticClass:"mdl-card__supporting-text"},[a("div",{staticClass:"info"},[a("label",{attrs:{for:"nsfw"}},[t._v("NSFW score")]),a("p",{attrs:{id:"nsfw"}},[t._v(t._s(e.nsfw?e.nsfw.nsfw_score.toPrecision(3):"error"))]),a("label",{attrs:{for:"exif"}},[t._v("EXIF")]),a("textarea",{directives:[{name:"model",rawName:"v-model",value:e.exif,expression:"sample.exif"}],attrs:{id:"exif"},domProps:{value:e.exif},on:{input:function(a){a.target.composing||t.$set(e,"exif",a.target.value)}}})])]),a("div",{staticClass:"mdl-card__actions mdl-card--border"},[a("a",{staticClass:"mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect",attrs:{target:"_blank",href:e.url}},[t._v(" View Image ")])])])})),0)])}),l=[],c=(a("4de4"),a("baa5"),a("d81d"),a("b64b"),a("07ac"),new s["a"]),o={name:"Links",created:function(){c.$on("search/complete",this.searchComplete),c.$on("search/start",this.searchStart),c.$on("search/refresh",this.searchRefresh)},data:function(){return{samples:[],searching:!1}},methods:{searchComplete:function(t){t=t.map((function(t){return t.exif.length>0&&(t.exif=t.exif.map((function(t){var e=Object.keys(t)[0],a=Object.values(t)[0];return e+"="+a})),t.exif=t.exif.filter((function(t){return"Error=no EXIF found in image"!==t}))),t.filename=t.url.substring(t.url.lastIndexOf("/")+1),t})),this.samples=t,this.searching=!1},searchStart:function(){this.searching=!0,this.samples=[]},searchRefresh:function(){this.searching=!1}}},d=o,u=(a("10b3"),a("2877")),f=Object(u["a"])(d,n,l,!1,null,"670d89f0",null),m=f.exports,h=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"links"},[a("div",{staticClass:"mdl-progress mdl-js-progress mdl-progress__indeterminate",class:{hidden:!t.crawling},attrs:{id:"p2"}}),0!=t.samples.length||t.crawling?t._e():a("div",{staticClass:"empty"},[t._v(" No image links. Please search above ")]),t.samples.length>0?a("ul",t._l(t.samples,(function(e){return a("li",{key:e},[a("a",{attrs:{href:e}},[t._v(t._s(e))])])})),0):t._e()])},p=[],v={name:"Links",created:function(){c.$on("crawl/complete",this.samplesReady),c.$on("crawl/start",this.crawlStart)},data:function(){return{samples:[],crawling:!1}},methods:{samplesReady:function(t){this.samples=t,this.crawling=!1},crawlStart:function(){this.crawling=!0,this.samples=[]}}},_=v,b=(a("fde3"),Object(u["a"])(_,h,p,!1,null,"b1b4d998",null)),g=b.exports,w=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"analytics"},[0!=t.items.length||t.hasData?t._e():a("div",{staticClass:"empty"},[t._v(" No images. Please search above ")]),t.items.length>0?a("div",{staticClass:"nsfw_stats"},[a("div",{staticClass:"demo-card-event mdl-card mdl-shadow--2dp"},[a("div",{staticClass:"mdl-card__title mdl-card--expand"},[a("h4",[t._v(" Images:"),a("br"),t._v(" "+t._s(t.imageCount)+" ")])])]),a("div",{staticClass:"demo-card-event mdl-card mdl-shadow--2dp"},[a("div",{staticClass:"mdl-card__title mdl-card--expand"},[a("h4",[t._v(" Images with EXIF:"),a("br"),t._v(" "+t._s(t.exifCount)+" ")])])]),a("div",{staticClass:"demo-card-event mdl-card mdl-shadow--2dp"},[a("div",{staticClass:"mdl-card__title mdl-card--expand"},[a("h4",[t._v(" Averge NSFW Score:"),a("br"),t._v(" "+t._s(t.averageNsfw.toPrecision(3))+" ")])])]),a("div",{staticClass:"demo-card-event mdl-card mdl-shadow--2dp"},[a("div",{staticClass:"mdl-card__title mdl-card--expand"},[a("h4",[t._v(" Maximum NSFW Score:"),a("br"),t._v(" "+t._s(t.maxNsfw.toPrecision(3))+" ")])])]),a("div",{staticClass:"demo-card-event mdl-card mdl-shadow--2dp"},[a("div",{staticClass:"mdl-card__title mdl-card--expand"},[a("h4",[t._v(" Minimum NSFW Score:"),a("br"),t._v(" "+t._s(t.minNsfw.toPrecision(3))+" ")])])])]):t._e(),t.items.length>0?a("div",{staticClass:"charts"},[a("div",{attrs:{id:"chart"}},[a("apexchart",{attrs:{type:"pie",width:"500",options:t.exifChartOptions,series:t.exifSeries}})],1),a("div",{attrs:{id:"chart2"}},[a("apexchart",{attrs:{type:"pie",width:"500",options:t.nsfwChartOptions,series:t.nsfwSeries}})],1)]):t._e()])},x=[],C=(a("13d5"),a("284c")),y=a("1321"),j=a.n(y),O={name:"Analytics",props:{},components:{apexchart:j.a},created:function(){c.$on("search/complete",this.update),c.$on("search/refresh",this.update)},data:function(){return{hasData:!1,items:[],imageCount:0,exifCount:0,averageNsfw:0,minNsfw:0,maxNsfw:0,exifChartOptions:{labels:["No EXIF","EXIF"],responsive:[{breakpoint:480,options:{chart:{width:200},legend:{position:"bottom"}}}]},nsfwChartOptions:{labels:["SFW > 0.5","NSFW >= 0.5"],responsive:[{breakpoint:480,options:{chart:{width:200},legend:{position:"bottom"}}}]}}},computed:{exifSeries:function(){return[this.imageCount,this.exifCount]},nsfwSeries:function(){return[this.items.filter((function(t){return t.nsfw})).filter((function(t){return t.nsfw.nsfw_score<.5})).length,this.items.filter((function(t){return t.nsfw})).filter((function(t){return t.nsfw.nsfw_score>=.5})).length]}},methods:{update:function(t){this.items=t,this.imageCount=t.length;var e=t.filter((function(t){return t.nsfw})).map((function(t){return t.nsfw.nsfw_score}));this.averageNsfw=e.reduce((function(t,e){return t+e}),0)/t.length,this.minNsfw=Math.min.apply(Math,Object(C["a"])(e)),this.maxNsfw=Math.max.apply(Math,Object(C["a"])(e)),this.hasData=!0}}},S=O,k=(a("ce5d"),Object(u["a"])(S,w,x,!1,null,"34f5e047",null)),$=k.exports,N=a("bc3a"),P=a.n(N),I={name:"app",components:{Images:m,Links:g,Analytics:$},data:function(){return{url:"",interval:0,crawlUrl:"?action=crawl",searchUrl:"?action=search",clearUrl:"?action=clear"}},created:function(){console.log("production")},methods:{siteUrlKeydown:function(){this.url.startsWith("http")||(this.url="http://"+this.url),c.$emit("crawl/start",{}),P.a.post(this.crawlUrl,this.url).then((function(t){c.$emit("crawl/complete",t.data)})),this.stopRefresh(),c.$emit("search/start",{}),P.a.post(this.searchUrl,this.url).then((function(t){c.$emit("search/complete",t.data)}))},refreshSearch:function(){this.interval=setInterval((function(){P.a.post("/action/search?action=search",this.url).then((function(t){c.$emit("search/refresh",t.data)}))}),5e3)},stopRefresh:function(){clearInterval(this.interval)},clearSearch:function(){c.$emit("crawl/complete",[]),c.$emit("search/complete",[]),P.a.post(this.clearUrl,"clear")}}},E=I,F=(a("034f"),Object(u["a"])(E,i,r,!1,null,null,null)),M=F.exports;s["a"].config.productionTip=!1,new s["a"]({render:function(t){return t(M)}}).$mount("#app")},"85ec":function(t,e,a){},c9a4:function(t,e,a){},ce5d:function(t,e,a){"use strict";var s=a("c9a4"),i=a.n(s);i.a},fde3:function(t,e,a){"use strict";var s=a("027f"),i=a.n(s);i.a}});
//# sourceMappingURL=app.js.map