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
                token = response["token"]
                if (token != "") {
                    setCookie("auth", token, 48 * 60) // 48 (hours) * 60 (minutes) = 2 days
                    window.location = "/test"
                }
            },
            statusCode: {
                401: function() {
                    $("#Erro").html("Usuário e/ou senha incorretos");
                    $("#senha_cad").val("");
                },
                502: function() {
                    alert("Ocorreu um erro no servidor. Tente novamente.")
                    $("#senha_cad").val("");
                }
            },
        });
    });
});

