function onDelete(id) {
    let resposta = confirm("Tem certeza que deseja deletar esse treino?");
    if (resposta) {
        fetch(`/delete/id/${id}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (response.ok) {
                // Redirecionar ou fazer algo após a exclusão
                window.location.href = '/index';
            } else {
                console.error('Erro ao excluir:', response.status);
            }
        })
        .catch(error => {
            console.error('Erro na solicitação DELETE:', error);
        });
    }
}
