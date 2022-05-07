$(document).ready(function () {

    var host = window.location.protocol + "//" + window.location.host

    $("#loginForm").submit(function (e) { 
        e.preventDefault();
        
        var serializedData = $(this).serialize();

        $.ajax({
            type: "post",
            url: `{host}:8080/login`,
            data: serializedData,
            dataType: "json",
            success: function (response) {
                token = response["token"];
                if (token != "") {
                    setCookie("auth", token, 48 * 60); // 48 (hours) * 60 (minutes) = 2 days
                    window.location = "/test";
                }
            },
            statusCode: {
                401: function() {
                    $("#Erro").html("Usuário e/ou senha incorretos");
                    $("#senha_login").val("");
                },
                502: function() {
                    alert("Ocorreu um erro no servidor. Tente novamente.");
                    $("#senha_login").val("");
                }
            },
        });
    });

    $.ajax({
        type: "post",
        url: `{host}:8080/validate_user`,
        data: JSON.stringify({
            token: getCookie("auth")
        }),
        dataType: "json",
        statusCode: {
            200: function() {
                window.location = "/test";
            }
        }
    });
});

