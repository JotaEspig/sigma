$(document).ready(function () {
    $("#tabela_turma").on("click", ".remover", function(){
        let tr = $(this).closest('tr');
        
        let confirmado = confirm('Deseja deletar?');
        if(confirmado){
            alert('Confirmado!');
        }else{
            alert('Negado!'); 
            return;  
        }

        tr.fadeOut(400, function() {
            tr.remove();  
        });
    });
});
