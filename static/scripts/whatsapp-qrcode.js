document.addEventListener('DOMContentLoaded', () => {
    const qrCodeImage = document.getElementById('qrcode-image');

    function fetchAndDisplayQrCode() {
        fetch('/whatsapp/qrcode')
            .then(response => {
                if (!response.ok) {
                    throw new Error(`Erro HTTP: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                const base64Data = data.qrCode;
                console.log(base64Data);

                if (base64Data) {
                    qrCodeImage.src = base64Data;
                    qrCodeImage.alt = 'QR Code para login do WhatsApp';
                } else {
                    qrCodeImage.alt = 'Dados do QR Code não encontrados.';
                }
            })
            .catch(() => {
                qrCodeImage.alt = 'Não foi possível carregar o QR Code.';
            });
    }

    fetchAndDisplayQrCode();
    setInterval(fetchAndDisplayQrCode, 20000);
});
