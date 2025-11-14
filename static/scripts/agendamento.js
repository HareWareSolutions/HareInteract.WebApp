
// Inicializa o calendário usando FullCalendar
document.addEventListener('DOMContentLoaded', function() {
        var calendarEl = document.getElementById('calendar');
        var calendar = new FullCalendar.Calendar(calendarEl, {
          initialView: 'dayGridMonth', // Exibe o calendário no modo mês
          height: '70vh', // Define a altura do calendário
          headerToolbar: { // Configura a barra de ferramentas do cabeçalho
            left: 'prev,next today',
            center: 'title',
            right: 'dayGridMonth,timeGridWeek,timeGridDay'
          },
          eventClick: function(info) { //Função eventClick é chamada quando um evento é clicado
            info.jsEvent.preventDefault(); // Impede o comportamento padrão do link
            let event = info.event;
            let id = event.id;
            let title = event.title;
            let start = event.start;

            let extendedProps = event.extendedProps;
            let description = extendedProps.contato || '';
            let responsavel = extendedProps.responsavel || '';
            let link = extendedProps.link || '';

            //Carrega os dados do evento no modal de edição
            document.getElementById('modal-edit-event').querySelector('.modal-content').innerHTML = `
                <p><strong>Título:</strong> ${title}</p>
                <p><strong>Data e Hora:</strong> ${start.toLocaleString()}</p>
                <p><strong>Contato:</strong> ${description}</p>
                <p><strong>Responsável:</strong> ${responsavel}</p>
                <p><strong>Link:</strong> <a href="${link}" target="_blank">${link}</a></p>
            `;
            // Abre o modal de edição de evento

            modalEditEvent.style.display = 'flex'; // Usa 'flex' para centralizar o conteúdo
          },
          dateClick: function(info){
            alert('Data clicada: ' + info.dateStr);
            alert('Eu vou implementar a funcionalidade de agendamento aqui!');
          },
          events: [ // Eventos de exemplo
            {
                id: '1',
                title: 'Exame do Dedo',
                start: '2025-11-14T10:00:00',
                extendedProps: {  
                contato:'João Silva',
                responsavel:'Dra. Maria Oliveira',
                link:'https://www.pudim.com.br/'
                }
            },
            {
                id: '2',
                title: 'Consulta de Rotina',
                start: '2025-11-16T14:00:00',
                extendedProps: {  
                contato:'Maria Souza',
                responsavel:'Dr. Carlos Pereira',
                link:'https://www.pudim.com.br/'
                }
            }
          ]
        });
        calendar.render();

        setTimeout(function(){
          calendar.updateSize();
        }, 0);
});

// Função para abrir o modal de edição de evento
const modalEditEvent = document.getElementById('modal-edit-event');
const modalCancelButton = document.getElementById('modal-cancel-button');

// Adiciona evento para fechar o modal
if (modalCancelButton && modalEditEvent) {
    modalCancelButton.addEventListener('click', () => {
        modalEditEvent.style.display = 'none';
    });
}