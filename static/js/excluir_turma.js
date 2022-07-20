var tabela = document.querySelector("#tabela_turma");

tabela.addEventListener("click", function() {
    var elementoClicado = event.target;
    if (elementoClicado.classList.contains("btn-excluir")) {
        var celula = elementoClicado.parentNode;
        var linha = celula.parentNode;
        linha.remove();
    }
})