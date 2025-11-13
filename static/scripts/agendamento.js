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
          eventClick: function(info){ // Ação ao clicar em um evento
            alert('Evento: ' + info.event.title + '\nInício: ' + info.event.start.toLocaleString());
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
                end: '2025-11-14T11:00:00'
            },
            {
                id: '2',
                title: 'Consulta de Rotina',
                start: '2025-11-16T14:00:00',
                end: '2025-11-16T15:00:00'  
            }
          ]
        });
        calendar.render();
});