document.addEventListener('DOMContentLoaded', () => {

    // --- LÓGICA DE TROCA DE ABAS (MANTIDA) ---
    const navItems = document.querySelectorAll('.settings-sidebar .nav-item');
    const tabContents = document.querySelectorAll('.settings-content .tab-content');

    navItems.forEach(item => {
        item.addEventListener('click', () => {
            const tabId = item.getAttribute('data-tab');
            navItems.forEach(nav => nav.classList.remove('active'));
            item.classList.add('active');
            tabContents.forEach(content => content.classList.remove('active'));
            const targetTab = document.getElementById(tabId);
            if (targetTab) {
                targetTab.classList.add('active');
            }
        });
    });


    // --- LÓGICA DO MDC DIALOG DE EDIÇÃO DE USUÁRIO (CORRIGIDA) ---
    const dialogElement = document.getElementById('edit-user-dialog');

    if (dialogElement) {
        // Inicializa o MDC Dialog
        const dialog = new mdc.dialog.MDCDialog(dialogElement);
        
        // 1. Inicializa e armazena as instâncias dos Text Fields do MDC
        // CORREÇÃO: Usa um seletor que EXCLUI explicitamente a div que será usada pelo MDC Select.
        // O MDC Select deve ser identificado pela classe mdc-select.
        const textFieldElements = dialogElement.querySelectorAll('.mdc-text-field:not(.mdc-select)');
        const mdcTextFieldInstances = [];
        
        textFieldElements.forEach(tfElement => {
            // A inicialização agora só ocorre nos elementos que contém o <input>
            mdcTextFieldInstances.push(new mdc.textField.MDCTextField(tfElement)); 
        });

        // 2. Inicializa e armazena a instância do MDC Select
        // Busca o elemento principal do MDC Select (a div com a classe mdc-select)
        const selectWrapper = dialogElement.querySelector('.mdc-select');
        let mdcSelectInstance = null;
        if (selectWrapper) {
            mdcSelectInstance = new mdc.select.MDCSelect(selectWrapper);
        } else {
            console.warn("Elemento wrapper do MDC Select ('mdc-select') não encontrado. Verifique se o div pai tem a classe 'mdc-select'.");
        }
        
        // Referência ao elemento <select> puro (para fallback)
        const selectNivel = dialogElement.querySelector('#edit-nivel-acesso');

        // Referências aos campos do formulário
        const inputId = document.querySelector('#user-edit-form input[name="id"]');
        const inputNome = document.getElementById('edit-nome');
        
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
                
                // Preenche o valor do MDC Select
                if (mdcSelectInstance) {
                    mdcSelectInstance.value = userNivel;
                } else if (selectNivel) {
                    selectNivel.value = userNivel;
                }
                
                // c) Re-inicializa o layout dos componentes
                mdcTextFieldInstances.forEach(instance => {
                    instance.layout();
                });

                if (mdcSelectInstance) {
                    mdcSelectInstance.layout();
                }

                // d) Abre o Dialog
                dialog.open();
            });
        });
        
    } else {
        console.warn("MDC Dialog de Edição ('#edit-user-dialog') não encontrado.");
    }
});