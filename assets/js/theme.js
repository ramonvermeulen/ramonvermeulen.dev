let currentTheme;

function updateThemeIcon(isDarkMode) {
    const iconElement = document.querySelector('.bx');
    if (iconElement) {
        if (isDarkMode) {
            iconElement.classList.replace('bxs-moon', 'bxs-sun');
        } else {
            iconElement.classList.replace('bxs-sun', 'bxs-moon');
        }
    }
}

function applyTheme() {
    const isDarkMode = currentTheme === "dark" ||
        (!currentTheme && window.matchMedia("(prefers-color-scheme: dark)").matches);
    document.documentElement.classList.toggle("dark", isDarkMode);
    updateThemeIcon(isDarkMode);
}

function toggleTheme() {
    const isDarkMode = document.documentElement.classList.contains("dark");
    currentTheme = isDarkMode ? "light" : "dark";
    localStorage.currentTheme = currentTheme;
    applyTheme();
}

function toggleMenu() {
    const menu = document.getElementById('menu');
    menu.classList.toggle('hidden');
}

document.addEventListener("DOMContentLoaded", () => {
    currentTheme = localStorage.currentTheme;
    applyTheme();

    const menuToggle = document.getElementById('menu-toggle');
    if (menuToggle) {
        menuToggle.addEventListener('click', toggleMenu);
    }
});