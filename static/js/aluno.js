$(document).ready(function () {

    $.ajax({
        type: "get",
        url: "/logged/user/get",
        success: function (response) {
            $("#username").html(response["user"]["username"]);
            document.title = response["user"]["username"];
        },
        statusCode: {
            401: function() {
                alert("Você não está logado no sistema!")
                window.location = "/login"
            }
        },
    });
});