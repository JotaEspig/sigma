/**
 * get url from response.
 * Obs.: response must have the keys: username and type.
 * @param {object} response 
 * @returns url to be redirect
 */
function getPageURLFromResponse(response) {
    switch (response["type"]) {
        case "student":
            return "/aluno";

        case "teacher":
            return "/professor"

        case "admin":
            return "/admin";

        default:
            return "/usuario";
    }
}

$(document).ready(function () {

    $("#loginForm").submit(function (e) { 
        e.preventDefault();
        
        var serializedData = $(this).serialize();

        $.ajax({
            type: "post",
            url: "/login",
            data: serializedData,
            dataType: "json",
            statusCode: {
                200: function(response) {
                    var token = response["token"];
                    setCookie("auth", token, 48 * 60); // 48 (hours) * 60 (minutes) = 2 days
                    window.location = getPageURLFromResponse(response);
                },
                401: function() {
                    $("#Erro").html("Usuário e/ou senha incorretos");
                    $("#username_login").val("");
                    $("#senha_login").val("");
                },
                502: function() {
                    alert("Ocorreu um erro no servidor. Tente novamente.");
                    $("$username_login").val("");
                    $("#senha_login").val("");
                }
            },
        });
    });

    if (getCookie("auth") != null) {
        // Does a request to check if the cookie is legit and hasn't expired
        $.ajax({
            type: "get",
            url: "/login/validate",
            statusCode: {
                200: function(response) {
                    window.location = getPageURLFromResponse(response);
                }
            }
        });
    }
});

