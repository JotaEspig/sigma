$(document).ready(function () {

    $.ajax({
        type: "get",
        url: "/validate_user",
        success: function (response) {
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