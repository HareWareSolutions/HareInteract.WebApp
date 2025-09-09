document.addEventListener('DOMContentLoaded', () => {
    // Lógica para alternar o menu lateral
    const sidebar = document.getElementById('sidebar');
    const toggleBtn = document.getElementById('sidebar-toggle-btn');

    if (toggleBtn && sidebar) {
        toggleBtn.addEventListener('click', () => {
            sidebar.classList.toggle('sidebar--minimized');
        });
    }

    // Lógica para destacar o item de menu ativo
    const path = window.location.pathname.substring(1);
    const links = document.querySelectorAll('.menu-link');

    links.forEach(link => {
        link.classList.remove('active');
        if (link.getAttribute('data-path') === path) {
            link.classList.add('active');
        }
    });

    // Lógica da barra de pesquisa
    const searchInput = document.getElementById('searchInput');
    const cards = document.querySelectorAll('.card');

    if (searchInput && cards.length > 0) {
        searchInput.addEventListener('input', (event) => {
            const searchTerm = event.target.value.toLowerCase();
            cards.forEach(card => {
                const title = card.getAttribute('data-title').toLowerCase();
                if (title.includes(searchTerm)) {
                    card.style.display = 'flex';
                } else {
                    card.style.display = 'none';
                }
            });
        });
    }
});