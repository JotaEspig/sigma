$(document).ready(function () {

    var serializedData = JSON.stringify({
        token: getCookie("auth")
    });

    $.ajax({
        type: "post",
        url: "http://127.0.0.1:8080/test",
        data: serializedData,
        dataType: "json",
        success: function (response) {
            $("#username").html(response["username"]);
        },
        statusCode: {
            401: function() {
                alert("Você não está logado no sistema!")
                window.location = "/login"
            }
        },
    });
});