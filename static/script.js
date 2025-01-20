// Theme toggle functionality
const themeToggle = document.getElementById('theme-toggle');
const html = document.documentElement;
const themeIcon = themeToggle.querySelector('i');

// saved theme preference
const savedTheme = localStorage.getItem('theme');
if (savedTheme) {
    html.setAttribute('data-theme', savedTheme);
    updateThemeIcon(savedTheme);
}

themeToggle.addEventListener('click', () => {
    const currentTheme = html.getAttribute('data-theme');
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
    
    html.setAttribute('data-theme', newTheme);
    localStorage.setItem('theme', newTheme);
    updateThemeIcon(newTheme);
});

function updateThemeIcon(theme) {
    themeIcon.className = theme === 'dark' ? 'fas fa-sun' : 'fas fa-moon';
}

// Like, dislike, and comment functionality
document.querySelectorAll('.action-btn').forEach(button => {
    button.addEventListener('click', () => {
        const countSpan = button.querySelector('.count');
        let count = parseInt(countSpan.textContent);
        countSpan.textContent = count + 1;
        
        button.style.transform = 'scale(1.2)';
        setTimeout(() => {
            button.style.transform = 'scale(1)';
        }, 200);
    });
});

document.addEventListener('DOMContentLoaded', () => {
    document.body.style.transition = 'background-color 0.3s, color 0.3s';
});