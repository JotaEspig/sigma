$(document).ready(function () {
    $("#form-cadastro-turma").on("submit", function (e) {
        e.preventDefault();

        let data = $(this).serialize();

        $.ajax({
            type: 'post',
            url: '/admin/tools/classroom/add',
            data: data,
            dataType: 'json',
            statusCode: {
                200: function () {
                    swal({
                        title: "Sucesso!",
                        text: "Turma cadastrada com sucesso!",
                        icon: "success",
                        button: "Ok"
                    });
                },
                400: function () {
                    swal({
                        title: "Erro!",
                        text: "Turma já cadastrada!",
                        icon: "error",
                        button: "Ok"
                    });
                }
            }
        });
    });
});