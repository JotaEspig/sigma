$(document).ready(function () {

    var serializedData = JSON.stringify({
        token: getCookie("auth")
    });

    $.ajax({
        type: "post",
        url: "/test",
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