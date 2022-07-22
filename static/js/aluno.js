$(document).ready(function () {

    var url = window.location.pathname;
    var username = url.split('/')[2];
    var newURL = "/aluno/"+username+"/get"

    $.ajax({
        type: "get",
        url: newURL,
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