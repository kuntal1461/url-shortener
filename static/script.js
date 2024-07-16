document.getElementById('shorten-form').addEventListener('submit', async function(event) {
    event.preventDefault();
    
    const originalUrl = document.getElementById('original-url').value;
    
    const response = await fetch('/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ original_url: originalUrl })
    });
    
    if (response.ok) {
        const result = await response.json();
        const shortenedUrl = result.shortened;
        
        document.getElementById('shortened-url').href = shortenedUrl;
        document.getElementById('shortened-url').textContent = shortenedUrl;
        document.getElementById('result').classList.remove('hidden');
    } else {
        alert('Failed to shorten URL');
    }
});
