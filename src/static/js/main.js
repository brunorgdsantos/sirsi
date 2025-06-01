document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('createUserForm');
    const responseMessage = document.getElementById('responseMessage');

    if (form) { // Verifica se o formulário existe na página
        form.addEventListener('submit', async (event) => {
            event.preventDefault(); // Impede o envio padrão do formulário e o recarregamento da página

            const name = document.getElementById('userName').value;
            const password = document.getElementById('password').value;

            try {
                // Ajuste a URL para o seu endpoint de criação de usuário no Gin
                const response = await fetch('/users', { // Sua rota POST /users
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ name: name, password: password }),
                });

                const data = await response.json(); // Pega a resposta JSON do backend

                if (response.ok) {
                    responseMessage.style.backgroundColor = '#d4edda';
                    responseMessage.style.borderColor = '#28a745';
                    responseMessage.style.color = '#155724';
                    responseMessage.textContent = `Usuário ${data.name} criado com sucesso! ID: ${data.id}`;
                    form.reset(); // Limpa os campos do formulário
                } else {
                    responseMessage.style.backgroundColor = '#f8d7da';
                    responseMessage.style.borderColor = '#dc3545';
                    responseMessage.style.color = '#721c24';
                    responseMessage.textContent = `Erro: ${data.message || 'Erro desconhecido ao criar usuário.'}`;
                }
            } catch (error) {
                console.error('Erro na requisição:', error);
                responseMessage.style.backgroundColor = '#fff3cd';
                responseMessage.style.borderColor = '#ffc107';
                responseMessage.style.color = '#856404';
                responseMessage.textContent = 'Erro de rede ou servidor não disponível.';
            }
        });
    }
});