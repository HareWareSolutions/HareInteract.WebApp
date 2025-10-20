document.addEventListener('DOMContentLoaded', () => {

    // --- LÓGICA DE TROCA DE ABAS ---
    const navItems = document.querySelectorAll('.settings-sidebar .nav-item');
    const tabContents = document.querySelectorAll('.settings-content .tab-content');

    navItems.forEach(item => {
        item.addEventListener('click', () => {
            const tabId = item.getAttribute('data-tab');

            // 1. Remove 'active' de todos os botões e adiciona ao clicado
            navItems.forEach(nav => nav.classList.remove('active'));
            item.classList.add('active');

            // 2. Remove 'active' de todos os conteúdos e mostra o conteúdo correspondente
            tabContents.forEach(content => content.classList.remove('active'));
            const targetTab = document.getElementById(tabId);
            if (targetTab) {
                targetTab.classList.add('active');
            }
        });
    });


    // --- LÓGICA DO MDC DIALOG DE EDIÇÃO DE USUÁRIO ---
    const dialogElement = document.getElementById('edit-user-dialog');

    if (dialogElement) {
        // Inicializa o MDC Dialog
        const dialog = new mdc.dialog.MDCDialog(dialogElement);
        
        // Inicializa e armazena as instâncias dos Text Fields do MDC
        const textFieldElements = dialogElement.querySelectorAll('.mdc-text-field');
        const mdcTextFieldInstances = []; // Array para guardar as instâncias
        
        textFieldElements.forEach(tfElement => {
            // Cria a instância e armazena
            mdcTextFieldInstances.push(new mdc.textField.MDCTextField(tfElement));
        });

        // Referências aos campos do formulário (usadas para preencher os valores)
        // **ATENÇÃO: id="edit-user-id" é usado duas vezes no seu HTML.
        // O input é o elemento correto para preencher o valor.
        const inputId = document.querySelector('#user-edit-form input[name="id"]');
        const inputNome = document.getElementById('edit-nome');
        const inputNivel = document.getElementById('edit-nivel-acesso');
        
        // Adiciona ouvintes de evento aos botões da tabela
        const openEditButtons = document.querySelectorAll('.open-edit-dialog-btn');

        openEditButtons.forEach(button => {
            button.addEventListener('click', () => {
                
                // a) Captura os dados do botão
                const userId = button.getAttribute('data-user-id');
                const userNome = button.getAttribute('data-user-nome');
                const userNivel = button.getAttribute('data-user-nivel');
                
                // b) Preenche os campos do formulário
                if (inputId) inputId.value = userId;
                if (inputNome) inputNome.value = userNome;
                if (inputNivel) inputNivel.value = userNivel;
                
                // c) Opcional: Re-inicializa o layout dos campos de texto
                // Isso garante que as labels flutuem para cima quando o valor é setado via JS.
                mdcTextFieldInstances.forEach(instance => {
                    instance.layout();
                });

                // d) Abre o Dialog
                dialog.open();
            });
        });
    } else {
        console.warn("MDC Dialog de Edição ('#edit-user-dialog') não encontrado.");
    }
});