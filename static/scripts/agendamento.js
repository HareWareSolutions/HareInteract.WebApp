const initialEvents = [
            { 
                id: '1',
                title: 'Consulta com Cliente A',
                start: '2025-11-14', // Dia inteiro
                color: '#4CAF50', // Verde
                extendedProps: { client: 'Cliente A', details: 'Revisão anual de contrato.' }
            },
            {
                id: '2',
                title: 'Reunião de Equipe',
                start: '2025-11-14T10:00:00', // Com hora
                end: '2025-11-14T11:30:00',
                color: '#FF9800', // Laranja
                extendedProps: { client: 'Interno', details: 'Discussão de OKRs do próximo trimestre.' }
            },
            {
                id: '3',
                title: 'Workshop',
                start: '2025-11-20',
                color: '#2196F3', // Azul
                extendedProps: { client: 'Startup X', details: 'Introdução ao Marketing Digital.' }
            }
        ];

        let allAppointments = initialEvents; // Fonte de dados
        let selectedDayEl = null;
        let calendar;

        const appointmentsListEl = document.getElementById('lista-agendamentos');
        const modalEl = document.getElementById('custom-modal');

        // Funções do Modal (para substituir alert() e confirm())
        function showCustomAlert(messageHtml) {
            const modalContent = document.getElementById('modal-content-area');
            if (modalContent) {
                modalContent.innerHTML = messageHtml;
                modalEl.style.display = 'flex'; // Exibe o modal
            }
        }
        window.closeCustomAlert = function() {
            modalEl.style.display = 'none'; // Esconde o modal
        }
        
        // Permite fechar o modal clicando fora
        modalEl.onclick = function(event) {
            if (event.target === modalEl) {
                closeCustomAlert();
            }
        }

        /**
         * FUNÇÃO PRINCIPAL: Renderiza a lista de agendamentos para a data selecionada
         * @param {string} dateStr - Data no formato 'YYYY-MM-DD'
         */
        function loadAppointmentsForDay(dateStr) {
            // 1. Limpa o destaque do dia anterior
            if (selectedDayEl) {
                // Remove a classe do dia anterior
                selectedDayEl.classList.remove('selected-date-day');
            }
            
            // 2. Filtra Agendamentos
            // Filtra os agendamentos que começam no dia
            const filteredAppointments = allAppointments.filter(event => 
                event.start.startsWith(dateStr)
            );
            
            // 3. Destaca o novo dia clicado
            // O FullCalendar usa o atributo data-date nas células do dia
            const dayCell = document.querySelector(`.fc-day[data-date="${dateStr}"]`);
            if (dayCell) {
                dayCell.classList.add('selected-date-day');
                selectedDayEl = dayCell;
            } else {
                // Se a célula do dia não for encontrada imediatamente, pode ser em uma visualização diferente
                // ou se o calendário estiver renderizando.
                console.warn(`Célula do dia ${dateStr} não encontrada no DOM.`);
            }

            // 4. Atualiza Título
            const date = new Date(dateStr + 'T12:00:00'); 
            const formatter = new Intl.DateTimeFormat('pt-BR', { weekday: 'long', day: 'numeric', month: 'long' });
            
            // 5. Renderiza a Lista
            let htmlContent = '';
            
            if (filteredAppointments.length === 0) {
                htmlContent = `
                    <h2>Agendamentos para ${formatter.format(date)}</h2>
                    <p class="text-gray-600 text-sm italic p-4 bg-gray-50 rounded-lg">Nenhum agendamento neste dia. Dia livre!</p>
                `;
            } else {
                // Monta o cabeçalho e a lista
                const listItems = filteredAppointments.map(app => {
                    // Determina a hora de início. Se incluir 'T', mostra hora. Senão, 'Dia Todo'.
                    const time = app.start.includes('T') ? new Date(app.start).toLocaleTimeString('pt-BR', { hour: '2-digit', minute: '2-digit' }) : 'Dia Todo';
                    const details = app.extendedProps?.details || 'Sem detalhes adicionais.';
                    const client = app.extendedProps?.client || 'N/A';
                    const color = app.color || '#3f51b5'; // Cor padrão

                    return `
                        <div class="appointment-item" style="border-left-color: ${color};">
                            <h3 style="color: ${color}; font-weight: 500;">${app.title}</h3>
                            <p style="font-size: 0.9em; color: #555;">
                                <span style="font-weight: bold;">${time}</span> | Cliente: ${client}
                            </p>
                            <p style="font-size: 0.8em; color: #777; margin-top: 5px;">${details}</p>
                        </div>
                    `;
                }).join('');

                htmlContent = `
                    <h2>Agendamentos para ${formatter.format(date)}</h2>
                    <div class="list-content space-y-3 mt-4">${listItems}</div>
                `;
            }

            appointmentsListEl.innerHTML = htmlContent;
        }


        document.addEventListener('DOMContentLoaded', function() {
            var calendarEl = document.getElementById('calendar');

            calendar = new FullCalendar.Calendar(calendarEl, {
                // Configurações básicas
                initialView: 'dayGridMonth',
                locale: 'pt-br',
                headerToolbar: {
                    left: 'prev,next today',
                    center: 'title',
                    right: 'dayGridMonth,timeGridWeek,timeGridDay'
                },
                // Usa os dados estáticos como fonte
                events: allAppointments,
                
                // Função para interagir com o clique em um dia específico
                dateClick: function(info) {
                    // Pega apenas a data (YYYY-MM-DD), ignorando a hora e fuso
                    const dateOnly = info.dateStr.split('T')[0];
                    loadAppointmentsForDay(dateOnly);
                },

                // Função para interagir com os eventos (agendamentos)
                eventClick: function(info) {
                    const event = info.event;
                    const date = new Date(event.start).toLocaleDateString('pt-BR');
                    const time = event.start.includes('T') ? 
                                 new Date(event.start).toLocaleTimeString('pt-BR', { hour: '2-digit', minute: '2-digit' }) : 
                                 'Dia Todo';
                    const details = event.extendedProps.details || 'Nenhum detalhe disponível.';
                    const client = event.extendedProps.client || 'Não informado';

                    showCustomAlert(`
                        <h3 style="color: ${event.color || '#3f51b5'}; font-weight: bold; margin-top: 0;">${event.title}</h3>
                        <p><strong>Cliente:</strong> ${client}</p>
                        <p><strong>Data:</strong> ${date}</p>
                        <p><strong>Hora:</strong> ${time}</p>
                        <p class="mt-3" style="font-size: 0.9em;"><strong>Detalhes:</strong> ${details}</p>
                    `);
                }
            });

            calendar.render();
            
            // Carrega os agendamentos do dia atual ao iniciar
            const today = new Date().toISOString().substring(0, 10);
            loadAppointmentsForDay(today);
        });