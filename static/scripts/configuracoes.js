document.addEventListener('DOMContentLoaded', () => {
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
            document.getElementById(tabId).classList.add('active');
        });
    });
});