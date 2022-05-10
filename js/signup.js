// TODO Eduardo: Adicionar um verificador de cadastro
// mínimo de 4 letras pro nome, username, senha, etc

$(document).ready(function () {

    $("#cadastroForm").submit(function (e) { 
        e.preventDefault();

        var name = $("#nome_cad").val();
        var username = $("#username_cad").val();
        var password = $("#senha_cad").val();
        var confirmPassword = $("#confirmar_senha_cad   ").val();
        
        if(name.length < 4){
            swal({
                title: "Erro",
                text: "nome pequeno, deve ter pelo menos 4 caracteres ",
                icon: "error",
                button: "OK"
            });
            $("#senha_cad").val("");
            $("#confirmar_senha_cad").val("");
            return

        }

        if(username.length < 4){
            swal({
                title: "Erro",
                text: "nome de usuário pequeno, deve ter pelo menos 4 caracteres ",
                icon: "error",
                button: "OK"
            });
            $("#senha_cad").val("");
            $("#confirmar_senha_cad").val("");
            return
        }

        if(password.length < 4){
            swal({
                title: "Erro",
                text: "senha pequena, deve ter pelo menos 4 caracteres ",
                icon: "error",
                button: "OK"
            });
            $("#senha_cad").val("");
            $("#confirmar_senha_cad").val("");
            return
        }

        if(password != confirmPassword){
            swal({
                title: "Erro",
                text: "Senhas incompativeis",
                icon: "error"
            });
            $("#senha_cad").val("");
            $("#confirmar_senha_cad").val("");
            return
        }
        var serializedData = $(this).serialize();
        
        $.ajax({
            type: "post",
            url: "/cadastro",
            data: $(this).serialize(),
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
                },
                default: function() {
                    alert("Ocorreu um erro interno inesperado. Tente novamente!");
                }
            }
        });
    });
});