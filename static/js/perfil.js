$(document).ready(function(){
    $.ajax({
            type : "get",
            url: "/usuario/get",
            statusCode: { 
                200: function(data) {
                    $("#name").text(data["user"]["name"])
                    $("#username").text(data["user"]["username"])
                    $("#email").text(data["user"]["email"])
                }
            
            }
        })

});

function editarPerfil(nome, apelido, email){
    var nome = $("#name").val()
    var apelido = $("#username").val()
    var email = $("#email").val()

};
    
   
    