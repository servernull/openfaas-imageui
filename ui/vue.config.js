
module.exports = {
    filenameHashing: false,
    devServer: {
      disableHostCheck: true,
      proxy: {
        '/action/crawl': {
          target: "http://127.0.0.1:8080/function/openfaas-imagecrawlerdemux",
        },
        '/action/search': {
          target: "http://127.0.0.1:8080/function/openfaas-imagesearch",
        },
        '/action/clear': {
          target: "http://127.0.0.1:8080/function/openfaas-elastic"
        }
      },
    }
}