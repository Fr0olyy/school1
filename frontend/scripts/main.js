// Смена языка
const languageSelector = document.querySelector('.language-selector');
const languageLinks = languageSelector.querySelectorAll('a');

languageLinks.forEach(link => {
    link.addEventListener('click', () => {
        languageLinks.forEach(link => link.classList.remove('active'));
        link.classList.add('active');
    });
});

