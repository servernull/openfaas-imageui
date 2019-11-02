
module.exports = {
    filenameHashing: false,
    devServer: {
      disableHostCheck: true,
      proxy: {
        '/action/crawl': {
          target: "http://192.168.99.104:31112/function/openfaas-imagecrawlerdemux",
        },
        '/action/search': {
          target: "http://192.168.99.104:31112/function/openfaas-imagesearch",
        },
        '/action/clear': {
          target: "http://192.168.99.104:31112/function/openfaas-elastic"
        }
      },
    }
}