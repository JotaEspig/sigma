$(document).ready(function () {
    $("#loginForm").submit(function (e) { 
        e.preventDefault();
        
        var serializedData = $(this).serialize();

        request = $.ajax({
            type: "post",
            url: "http://127.0.0.1:8080/login",
            data: serializedData,
            dataType: "json",
            success: function (response) {
                if (response != "") {
                    document.cookie = "auth=" + response["token"]
                }
            },
            error: function() {
                alert("FUDEU")
            }
        });
    });
});

