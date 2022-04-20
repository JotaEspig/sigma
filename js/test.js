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
                    alert("Cookie with the token has been set")
                }
            },
            error: function() {
                $("#Erro").html("Usuário e/ou senha incorretos");
            }
        });
    });
});

