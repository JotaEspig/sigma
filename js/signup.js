$(document).ready(function () {

    var host = window.location.protocol + "//" + window.location.host

    $("#cadastroForm").submit(function (e) { 
        e.preventDefault();
        
        var serializedData = $(this).serialize();

        $.ajax({
            type: "post",
            url: `{host}:8080/cadastro`,
            data: serializedData,
            dataType: "json",
            statusCode: {
                200: function() {
                    alert("Usuário criado!");
                    window.location = "/login";
                },
                409: function() {
                    alert("Esse nome de usuário já existe");
                    $("#senha_cad").val("");
                    $("#username_cad").val("");
                }
            }
        });
    });
});