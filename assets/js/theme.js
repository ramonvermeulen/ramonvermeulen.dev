let currentTheme;

function updateThemeIcon(isDarkMode) {
    const icons = document.querySelectorAll('.bx');
    icons.forEach(icon => {
        icon.classList.toggle('bxs-moon', !isDarkMode);
        icon.classList.toggle('bxs-sun', isDarkMode);
    });
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