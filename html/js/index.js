$(document).ready(function () {
    $.ajax({
        type: "GET",
        async: false,
        url: "http://127.0.0.1:8080/meun/all",
        dataType: "jsonp",
        jsonpCallback:"?",
        success: function (data) {
            console.log("xxxxx")
        },
        error: function(err){
            console.log(err)
        }
    });
});