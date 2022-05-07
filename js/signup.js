// TODO Eduardo: Adicionar um verificador de cadastro
// mínimo de 4 letras pro nome, username, senha, etc

$(document).ready(function () {

    $("#cadastroForm").submit(function (e) { 
        e.preventDefault();
        
        var serializedData = $(this).serialize();

        $.ajax({
            type: "post",
            url: "/cadastro",
            data: serializedData,
            dataType: "json",
            statusCode: {
                200: function() {
                    swal({
                        title: "Sucesso!",
                        text: "Usuário criado",
                        icon: "success",
                        button: "OK",
                    })
                    .then(() => {
                        eraseCookie("auth");
                        window.location = "/login";
                    });
                },
                409: function() {
                    swal({
                        title: "Algo deu errado!",
                        text: "Nome de usuário já existe!",
                        icon: "error",
                        button: "OK"
                    })
                    .then(() => {
                        $("#senha_cad").val("");
                        $("#username_cad").val("");
                    });
                }
            }
        });
    });
});