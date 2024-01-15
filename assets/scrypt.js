document.getElementById('decryptButton').addEventListener('click', function() {
    const inputText = document.getElementById('inputText').value;
    const resultDiv = document.getElementById('result');

let countdown = 3;
resultDiv.textContent = countdown;
resultDiv.classList.remove('hidden');

const interval = setInterval(() => {
    countdown--;
    resultDiv.textContent = countdown;

    if (countdown <= 0) {
        clearInterval(interval);
        // Ici, vous pouvez ajouter la logique de déchiffrement en fonction de inputText
        resultDiv.textContent = "Résultat du déchiffrement : " + inputText; // Remplacez inputText par le résultat réel
    }
}, 1000); // Met à jour toutes les 1 seconde
});