$(document).ready(function () {

    $.ajax({
        type : "get",
        url: "/admin/tools/classroom/get",
        statusCode: {
            200: function(data){
                for (let item of data){
                    let id = item.id
                    let name = item.name
                    let year = item.year
                    addClassroomOnTable(id, name, year)
                    

                }
            }
        }
    })

    $("#form-user-search").on("submit", function (e) {
        e.preventDefault();

        let username = $("#username-search").val();
        let url = "/search/users/"+username;

        $.ajax({
            type: "get",
            url: url,
            statusCode: {
                200: function (data) {
                    alert(data)
                }
            }
        });
    });

    $("#form-cadastro-turma").on("submit", function (e) {
        e.preventDefault();

        let data = $(this).serialize();

        $.ajax({
            type: 'post',
            url: '/admin/tools/classroom/add',
            data: data,
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

function addClassroomOnTable(id, name, year) {
    $("#corpo_tabela").append(`
        <tr id='${id}'>
            <td scope="row" id='name'></td>
            <td id='year'></td>
            <td class="td-acoes">
            <span class="remover btn btn-danger">
                Excluir
            </span>
            </td>
        </tr>
    `);

    let classroomTable = $("#corpo_tabela").find("#"+id);
    classroomTable.find("#name").text(name);
    classroomTable.find("#year").text(year);
}