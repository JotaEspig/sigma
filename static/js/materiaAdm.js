$(document).ready(function () {

    $.ajax({
        type : "get",
        url: window.location.href + "/get",
        statusCode: {
            404: function(){
                alert("Não encontrado")
            },
            200: function(dados){
                classroom = dados["classroom"]
                alert(classroom.name)

            }
        }
    });


})