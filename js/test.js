$(document).ready(function () {

    $.ajax({
        type: "get",
        url: "/validate_user",
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