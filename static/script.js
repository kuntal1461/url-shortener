document.getElementById('shortenButton').addEventListener('click', function() {
    const originalUrl = document.getElementById('urlInput').value;
    fetch('/api/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ originalUrl: originalUrl })
    })
    .then(response => response.json())
    .then(data => {
        if (data.status === 'success') {
            document.getElementById('result').textContent = `Short URL: ${data.data.shortURL}`;
        } else {
            alert(data.message);
        }
    })
    .catch(error => console.error('Error:', error));
});
