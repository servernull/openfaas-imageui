document.getElementById("crawl").addEventListener("click", function(){
    fetch('?action=crawl', {
        method: 'post',
        body: document.getElementById("site").value
      }).then(function(response) {
        return response.json();
        }).then(function(data) {
            $("#content").empty();
            $("#content").append("<ul>")
            data.forEach(d => { 
                $("#content").append("<li><a href=\""+d+"\">"+d+"</a></li>")
            });
            $("#content").append("</ul>")
    });
});
document.getElementById("search").addEventListener("click", function(){
    fetch('?action=search', {
        method: 'post',
        body: document.getElementById("site").value
      }).then(function(response) {
        return response.json();
        }).then(function(data) {
            $("#content").empty();
            data.forEach(d => {
                $("#content").append("<img src=\""+d.url+"\"/>")
            })
    });
});